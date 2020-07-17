package main

import (
	"fmt"
	"log"
	"time"

	"github.com/polipopoliko/todo/todo2/iterator"

	"github.com/polipopoliko/todo/todo2/model"
	"github.com/polipopoliko/todo/todo2/storage"
)

func memoryFactory(id int) storage.Storage {
	var s storage.Storage
	switch id {
	case 1:
		s = storage.NewMemory()
	case 2:
		s = storage.Mock{}
	case 3:
		// s = storage.ConnectDb("postgres", "127.0.0.1", "5432", "postgres", "password", "storage_database", "disable")
		s = storage.ConnectDb("postgres", "postgres", "5432", "postgres", "password", "db", "disable")
	default:
		log.Fatal("Memory or Storage not found")
	}
	return s
}

func main() {
	// var memStore = storage.NewMemory()
	var store = memoryFactory(3)

	// create Todo
	if store.Create(model.Todo{
		Id:          100,
		Title:       "Hundreds of Todo",
		Description: "My Tenth Todo",
		CreatedAt:   time.Now(),
	}) != nil {
		log.Fatal("Fail to create todo")
	}

	// if store.Create(model.Todo{
	// 	Id:          101,
	// 	Title:       "Eleven Todo",
	// 	Description: "My Eleventh Todo",
	// 	CreatedAt:   time.Now(),
	// }) != nil {
	// 	log.Fatal("Fail to create todo")
	// }

	// if store.Create(model.Todo{
	// 	Id:          12,
	// 	Title:       "Twelve Todo",
	// 	Description: "My Twelfth Todo",
	// 	CreatedAt:   time.Now(),
	// }) != nil {
	// 	log.Fatal("Fail to create todo")
	// }

	obj, err := store.Detail(1)
	if err == nil {
		fmt.Println("Detail Todo with id", obj.Id, "title", obj.Title, "Description", obj.Description)
	} else {
		fmt.Println(err)
	}

	arrTodo, err := store.List()
	if err != nil {
		log.Fatal(err)
	}

	idp := iterator.TodoCollection{Todos: arrTodo}.CreateTodoCollection()
	for idp.HasNext() {
		v := idp.GetNext()
		fmt.Println("Todo with id", v.Id, "title", v.Title, "Description", v.Description)
	}

	// for _, v := range arrTodo {
	// 	fmt.Println("Todo with id", v.Id, "title", v.Title, "Description", v.Description)
	// }
}
