package main

import (
	"fmt"
	"log"
	"time"

	"github.com/polipopoliko/todo/todo2/iterator"

	pattern "github.com/polipopoliko/todo/todo2/design-pattern"
	"github.com/polipopoliko/todo/todo2/model"
)

func main() {
	var store = pattern.MemoryFactory(3)

	// create Todo
	if store.Create(model.Todo{
		Id:          100,
		Title:       "Hundreds of Todo",
		Description: "My Tenth Todo",
		CreatedAt:   time.Now(),
	}) != nil {
		log.Fatal("Fail to create todo")
	}

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
}
