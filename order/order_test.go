package order

import (
	"log"
	//	"reflect"
	"testing"
)

//func TestOrder(t *testing.T) {
//	order := NewOrder()
//	//fmt.Println(order)
//	if order == nil {
//		t.Fatal("new 'order' is null")
//	}
//	order.NewDetail()
//	order.NewDetail()
//	//fmt.Println(order)
//
//	repo := GetRepository()
//	repo.Add(order)
//}
var order *Order
var repo Repositor = GetRepository("memory")

func TestCreate(t *testing.T) {
	repo = GetRepository("memory")
	order = Create()
	orders, _ := repo.GetAll()
	log.Println(orders)
}
