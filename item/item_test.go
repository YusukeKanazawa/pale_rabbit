package item

import (
	"fmt"
	"testing"
)

var repo Repositor

func Init() {
	repo = GetRepository("memory")
	repo.Add(NewItem("A0001", "item1", 100))
	repo.Add(NewItem("A0002", "item2", 100))
	repo.Add(NewItem("A0003", "item3", 100))
	repo.Add(NewItem("A0004", "item4", 100))
}

func TestItem(t *testing.T) {
	Init()
	items := repo.GetAll()
	for _, item := range items {
		fmt.Println(item)
	}
}
