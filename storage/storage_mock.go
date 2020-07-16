package storage

import (
	"errors"
	"time"

	"github.com/polipopoliko/todo/todo2/model"
)

type Mock struct {
}

func (o Mock) Create(obj model.Todo) error {
	return nil
}

func (o Mock) Detail(id int32) (model.Todo, error) {
	if id == 1 {
		return model.Todo{Id: 1, Title: "Mock Title", Description: "Me Mock", CreatedAt: time.Now()}, nil
	}
	return model.Todo{}, errors.New("Todo not found")
}

func (o Mock) Delete(id int32) error {
	return nil
}

func (o Mock) List() ([]model.Todo, error) {
	return []model.Todo{
		model.Todo{Id: 2, Title: "TodoMock2", Description: "TodoMock2", CreatedAt: time.Now()},
		model.Todo{Id: 3, Title: "TodoMock3", Description: "TodoMock3", CreatedAt: time.Now()},
	}, nil
}
