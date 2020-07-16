package storage

import (
	"errors"

	"github.com/polipopoliko/todo/todo2/model"
)

type Memory struct {
	store map[int32]model.Todo
}

func NewMemory() Memory {
	return Memory{
		store: make(map[int32]model.Todo),
	}
}

func (o Memory) Create(obj model.Todo) error {
	o.store[obj.Id] = obj
	return nil
}

func (o Memory) Detail(id int32) (model.Todo, error) {
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Todo{}, errors.New("Error 404 Todo not found")
}

func (o Memory) Delete(id int32) error {
	delete(o.store, id)
	return nil
}

func (o Memory) List() ([]model.Todo, error) {
	arrTodo := []model.Todo{}
	for _, v := range o.store {
		arrTodo = append(arrTodo, v)
	}
	if len(arrTodo) > 0 {
		return arrTodo, nil
	}
	return arrTodo, errors.New("Todo Doesn't exist")
}
