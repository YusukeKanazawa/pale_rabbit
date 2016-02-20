package order

import (
  "time"
)


func Create() *Order{
  return &Order{
    Id:        generateId(),
    Details:   make([]*Detail, 0, 10),
    UpdatedAt: time.Now(),
  }
}
var seq_no = 0
func generateId() int {
	id := seq_no
	seq_no++
	return id
}