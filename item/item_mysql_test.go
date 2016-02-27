package item

import (
	"testing"
)
//func TestItemTypeMySql(t *testing.T) {
//	repo := GetRepository("mysql")
//	item := repo.FindByCode("TW")
//	t.Log(item)
//}

func TestMysqlRepository(t *testing.T) {
	repo := GetRepository("mysql")
	item := NewItemType("TEST", "test")
	repo.Add(item)
	items := repo.GetAll()
	item = repo.FindByCode(items[0].Code)
	t.Log(item)
	repo.Drop(item)
	items = repo.GetAll()
	t.Log(items)

}
