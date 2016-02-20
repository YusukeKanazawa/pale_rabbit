package mdse

import (
	"pale_rabbit/common"
)

// 商品構造体
type Mdse struct {
	MdseCode  string
	Name  string
	Price common.Price
	Enable bool
}

