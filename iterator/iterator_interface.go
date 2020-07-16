package iterator

import "github.com/polipopoliko/todo/todo2/model"

type Iterator interface {
	HasNext() bool
	GetNext() model.Todo
}

type TodoIterator struct {
	index    int
	todoList []model.Todo
}

func (o *TodoIterator) HasNext() bool {
	if o.index < len(o.todoList) {
		return true
	}
	return false
}

func (o *TodoIterator) GetNext() model.Todo {
	if o.HasNext() {
		todo := o.todoList[o.index]
		o.index++
		return todo
	}
	return model.Todo{}
}

type Collection interface {
	CreateTodoCollection() Iterator
}

type TodoCollection struct {
	Todos []model.Todo
}

func (o TodoCollection) CreateTodoCollection() Iterator {
	return &TodoIterator{
		index:    0,
		todoList: o.Todos,
	}
}
