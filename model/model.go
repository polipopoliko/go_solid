package model

import (
	"time"
)

type Todo struct {
	Id          int32
	Title       string
	Description string
	CreatedAt   time.Time
}

// func (o []Todo) CreateIterator() storage.Iterator {
// 	return TodoIterator{
// 		index:    0,
// 		todoList: o,
// 	}
// }
