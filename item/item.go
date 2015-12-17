package item

import (
	"pale_rabbit/common"
)

// 商品構造体
type Item struct {
	ItemCode  string
	Name  string
	Price common.Price
	Enable bool
}
