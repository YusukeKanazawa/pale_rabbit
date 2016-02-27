package api

import (
	"encoding/json"
	//	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"pale_rabbit/item"
)

func ItemTypeHandler(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")

	repo := item.GetRepository("mysql")
	if code != "" {
		item := repo.FindByCode(code)
		log.Println(item)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if item == nil {
			w.Write([]byte("{}"))
		} else {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
		}
	} else {
		items := repo.GetAll()
		log.Println(items[0])
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		res, err :=  json.Marshal(items); 
		if err != nil{
			panic(err)
		}
		
		_, err = w.Write(res);
		if err != nil{
			panic(err)
		}
	}
}

func ItemTypePutHandler(w http.ResponseWriter, req *http.Request) {
	repo := item.GetRepository("mysql")
	var item_type item.ItemType
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 4096))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &item_type); err != nil {
		repo.Add(&item_type)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

}
