package mdse

import (
	"fmt"
	"testing"
	"encoding/csv"
	"strings"
	"io"
	"log"
)

var repo Repositor

func Init() {
	repo = GetRepository("memory")
	repo.Add(NewMdse("A0001", "mdse1", 100))
	repo.Add(NewMdse("A0002", "mdse2", 100))
	repo.Add(NewMdse("A0003", "mdse3", 100))
	repo.Add(NewMdse("A0004", "mdse4", 100))
}

func TestMdse(t *testing.T) {
	Init()
	mdses := repo.GetAll()
	for _, mdse := range mdses {
		fmt.Println(mdse)
	}

}

func TestSpecItem(t *testing.T) {
	spec_item := NewSpecItem("color", "")
	fmt.Println(spec_item)
}

func TestCvs(t *testing.T) {
	in := `A0001,BLUE,S
A0001,BLUE,M
A0001,BLUE,L
A0001,RED,S
A0001,RED,M
A0001,RED,L
`
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0])
	}
}
