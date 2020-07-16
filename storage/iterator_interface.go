package storage

import "github.com/polipopoliko/todo/todo2/model"

type Iterator interface {
	HasNext() bool
	GetNext() model.Todo
}

type Collection interface {
	CreateTodoCollection() Iterator
}
