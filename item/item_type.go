package item

import (
	"time"
)

// 品目種別
type ItemType struct {
	Code      string           `json:"code"`
	Name      string           `json:"name"`
	Attrs     itemTypeAttrList `json:"attrs"`
	UpdatedBy string           `json:"updated_by"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// 新しい品目種別
func NewItemType(itemTypeCode string, name string) *ItemType {
	return &ItemType{
		Code: itemTypeCode,
		Name: name,
	}
}

// 新規品目種別属性を追加
func (itemType *ItemType) AppendAttr(code string, name string) {
	item := newItemTypeAttr(code, name)
	itemType.Attrs[item.AttrCode] = item
}

// 品目種別属性追加
func (itemType *ItemType) AddAttr(attr *itemTypeAttr) {
	itemType.Attrs[attr.AttrCode] = attr
}

// 品目種別属性削除
func (itemType *ItemType) RemoveAttr(attrCode string) {
	delete(itemType.Attrs, attrCode)
}

// 品目種別属性を更新
func (itemType *ItemType) UpdateAttr(attr *itemTypeAttr) {
	itemType.Attrs[attr.AttrCode] = attr
}

// 品目種別属性
type itemTypeAttr struct {
	AttrCode  string    `json:"attr_code"`
	AttrName  string    `json:"attr_name"`
	Rank      int       `json:"rank"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type itemTypeAttrList map[string]*itemTypeAttr

// 新しい品目種別属性
// ドメイン内部で生成
func newItemTypeAttr(code string, name string) *itemTypeAttr {
	return &itemTypeAttr{
		AttrCode: code,
		AttrName: name,
	}
}
