package storage

import "github.com/polipopoliko/todo/todo2/model"

type Storage interface {
	// identify activity dari storage
	Create(obj model.Todo) error
	Detail(id int32) (model.Todo, error)
	Delete(id int32) error
	List() ([]model.Todo, error)
}
