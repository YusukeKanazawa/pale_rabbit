package common

import (
	"log"
	"testing"
)

func TestPrice(t *testing.T) {
	var price Price = 128
	var intax = price.GetInTax()
	if intax == 138 {
		log.Printf("%d => %d", price, intax)
	} else {
		t.Fatalf("Unexpected price %d. Expected 138.", intax)
	}

	price = 258
	intax = price.GetInTax()
	if intax == 279 {
		log.Printf("%d => %d", price, intax)
	} else {
		t.Fatalf("Unexpected price %d. Expected 279.", intax)
	}

}
