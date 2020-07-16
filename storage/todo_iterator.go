package storage

import "github.com/polipopoliko/todo/todo2/model"

type TodoIterator struct {
	index    int
	todoList []model.Todo
}

func (o TodoIterator) HasNext() bool {
	if o.index < len(o.todoList) {
		return true
	}
	return false
}

func (o TodoIterator) GetNext() model.Todo {
	if o.HasNext() {
		todo := o.todoList[o.index]
		o.index++
		return todo
	}
	return model.Todo{}
}

type TodoCollection struct {
	Todos []model.Todo
}

func (o TodoCollection) CreateTodoCollection() Iterator {
	return TodoIterator{
		todoList: o.Todos,
	}
}
