package item

import (
	"time"
)

// 品目種別
type ItemType struct {
	Code      string          `json:"code"`
	Name      string          `json:"name"`
	Attrs     []*itemTypeAttr `json:"attrs"`
	UpdatedBy string          `json:"updated_by"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// 新しい品目種別
func NewItemType(itemTypeCode string, name string) *ItemType {
	return &ItemType{
		Code: itemTypeCode,
		Name: name,
	}
}
func (itemType *ItemType) AppendAttr(code string, name string) {

}
func (itemType *ItemType) AddAttr(attr *itemTypeAttr) {

}

// 品目種別属性
type itemTypeAttr struct {
	AttrCode  string    `json:"attr_code"`
	AttrName  string    `json:"attr_name"`
	Rank      int       `json:"rank"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ドメイン内部で生成
func newItemTypeAttr(code string, name string) *itemTypeAttr {
	return &itemTypeAttr{
		AttrCode: code,
		AttrName: name,
	}
}
