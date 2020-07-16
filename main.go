package main

import (
	"fmt"
	"log"

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
		s = storage.ConnectDb("postgres", "127.0.0.1", "5432", "postgres", "password", "storage_database", "disable")
	default:
		log.Fatal("Memory or Storage not found")
	}
	return s
}

func main() {
	// var memStore = storage.NewMemory()
	var store = memoryFactory(3)

	// create Todo
	// if store.Create(model.Todo{
	// 	Id:          1,
	// 	Title:       "First Todo",
	// 	Description: "My First Todo",
	// 	CreatedAt:   time.Now(),
	// }) != nil {
	// 	log.Fatal("Fail to create todo")
	// }

	// if store.Create(model.Todo{
	// 	Id:          2,
	// 	Title:       "Second Todo",
	// 	Description: "My Second Todo",
	// 	CreatedAt:   time.Now(),
	// }) != nil {
	// 	log.Fatal("Fail to create todo")
	// }

	// if store.Create(model.Todo{
	// 	Id:          3,
	// 	Title:       "Third Todo",
	// 	Description: "My Third Todo",
	// 	CreatedAt:   time.Now(),
	// }) != nil {
	// 	log.Fatal("Fail to create todo")
	// }

	// obj, err := store.Detail(1)
	// if err == nil {
	// 	fmt.Println("Detail Todo with id", obj.Id, "title", obj.Title, "Description", obj.Description)
	// } else {
	// 	fmt.Println(err)
	// }

	arrTodo, err := store.List()
	if err != nil {
		log.Fatal(err)
	}

	iterator := arrTodo.CreateIterator()
	for iterator.HasNext() {
		v := iterator.GetNext()
		fmt.Println("Todo with id", v.Id, "title", v.Title, "Description", v.Description)
	}

	// for _, v := range arrTodo {
	// 	fmt.Println("Todo with id", v.Id, "title", v.Title, "Description", v.Description)
	// }

}