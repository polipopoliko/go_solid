package storage

import (
	"errors"

	"github.com/polipopoliko/todo/todo2/model"
)

type memory struct {
	store map[int32]model.Todo
}

func NewMemory() memory {
	return memory{
		store: make(map[int32]model.Todo),
	}
}

func (o memory) Create(obj model.Todo) error {
	o.store[obj.Id] = obj
	return nil
}

func (o memory) Detail(id int32) (model.Todo, error) {
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Todo{}, errors.New("Error 404 Todo not found")
}

func (o memory) Delete(id int32) error {
	delete(o.store, id)
	return nil
}

func (o memory) List() ([]model.Todo, error) {
	arrTodo := []model.Todo{}
	for _, v := range o.store {
		arrTodo = append(arrTodo, v)
	}
	if len(arrTodo) > 0 {
		return arrTodo, nil
	}
	return arrTodo, errors.New("Todo Doesn't exist")
}
