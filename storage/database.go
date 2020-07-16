package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/polipopoliko/todo/todo2/model"
)

type PqDatabase struct {
	db *sql.DB
}

func ConnectDb(dbType, host, port, user, password, dbName, sslmode string) PqDatabase {
	dbCon, err := sql.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode))
	if err != nil {
		log.Fatal(err)
	}
	err = dbCon.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return PqDatabase{db: dbCon}
}

func (o PqDatabase) Create(obj model.Todo) error {
	_, err := o.db.Exec("INSERT INTO TODOS (id, title, description, created_at) VALUES ($1, $2, $3, $4)", obj.Id, obj.Title, obj.Description, obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (o PqDatabase) Detail(id int32) (model.Todo, error) {
	rows, err := o.db.Query("SELECT ID, TITLE, DESCRIPTION, CREATED_AT FROM TODOS WHERE ID = $1", id)
	if err != nil {
		return model.Todo{}, nil
	}
	defer rows.Close()
	var todo model.Todo
	for rows.Next() {
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
	}
	return todo, nil
}

func (o PqDatabase) Delete(id int32) error {
	_, err := o.db.Exec("DELETE FROM TODOS WHERE ID = $1", string(id))
	if err != nil {
		return err
	}
	return nil
}

func (o PqDatabase) List() ([]model.Todo, error) {
	var todoList = []model.Todo{}
	rows, err := o.db.Query("SELECT ID, TItle, Description, Created_at FROM TODOS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var temp model.Todo = model.Todo{}
		rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.CreatedAt)
		todoList = append(todoList, temp)
	}
	return todoList, nil
}
