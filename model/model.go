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
