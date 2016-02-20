package api

import (
	"pale_rabbit/item"
	"testing"
)

func TestOrder(t *testing.T) {
	item_repo := item.GetRepository("meomry")
	item_repo.Add(&item.Item{
		ItemCode:   "A0001",
		Name:   "item1",
		Price:  100,
		Enable: true,
	})
	t.Log(item_repo)
	items := item_repo.GetAll()
	t.Log(items[0])
	odr:= order.Create()
	t.Log(odr)
	odr.NewDetail(items[0].ItemCode,10)
	t.Log(odr)
}
