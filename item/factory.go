package item

import (
	"pale_rabbit/common"
)

// Create new item
func NewItem(code string, name string, price common.Price) *Item {
	item := &Item{
		ItemCode: code,
		Name:     name,
		Price:    price,
		Enable:   true,
	}
	return item
}
