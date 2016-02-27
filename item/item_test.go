package item

import (
	"testing"
)

func TestItemType(t *testing.T) {
	var itemType = NewItemType("TEST01", "test item")
	t.Log(itemType);
}
