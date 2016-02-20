package order

import (
	"time"
)

type Order struct {
	Id        int `json:"id"`
	Details   []*Detail
	UpdatedAt time.Time `json:"updated_at"`
}

func (this *Order) NewDetail(item string, amount int) {
	this.Details = append(
		this.Details,
		&Detail{
			DetailNo: len(this.Details),
			ItemCode: item,
			Amount:   amount,
		})
}
