package common

import (
	"math"
)

const (
	CONSUMPTION_TAX_RATE = 0.08
	ROUND                = iota
	TRUNC
	CEIL
	TAX_MODE = ROUND
)

type Price int

// Value Object.
func (price *Price) GetInTax() int {
	rate_over := CONSUMPTION_TAX_RATE + 1
	intax := rate_over * float64(*price)
	switch {
	case TAX_MODE == TRUNC:
		return int(math.Trunc(intax))
	case TAX_MODE == CEIL:
		return int(math.Ceil(intax))
	default:
		return int(math.Floor(intax + .5))
	}
}
