package mdse

import (
	"pale_rabbit/common"
)

// Create new marchandise.
func NewMdse(code string, name string, price common.Price) *Mdse {
	mdse := &Mdse{
		MdseCode: code,
		Name:     name,
		Price:    price,
		Enable:   true,
	}
	return mdse
}

// Create new specification item.
func NewSpecItem(name string, unit string) *SpecItem {
	id := 0
	spec_item := &SpecItem{
		Id: id,
		Name: name,
		Unit: unit,
	}
	return spec_item
}