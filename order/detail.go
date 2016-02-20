package order

import (
	"time"
)

type Detail struct {
	DetailNo  int       `josn:"detail_no" sql:"pk"`
	ItemCode  string	`json:"item_code"`
	Amount int          `json:"item_code"`
	UpdatedAt time.Time `json:"updated_at"`
}
