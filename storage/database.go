package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/polipopoliko/todo/todo2/model"
)

type pqDatabase struct {
	db *sql.DB
}

func ConnectDb(dbType, host, port, user, password, dbName, sslmode string) pqDatabase {
	dbCon, err := sql.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode))
	if err != nil {
		log.Fatal(err)
	}
	err = dbCon.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db := pqDatabase{db: dbCon}
	// createTable(db)
	return db
}

func checkIfTableExist(o pqDatabase) bool {
	_, err := o.db.Query("SELECT * FROM TODOS;")
	if err != nil {
		return false
	}
	return true
}

func createTable(o pqDatabase) {
	if checkIfTableExist(o) {
		fmt.Println("Table todos exist")
	} else {
		_, err := o.db.Exec("CREATE TABLE TODOS (id int, title varchar(50), description varchar(150), created_at timestamp);")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Create table TODOS success")
	}
}

func (o pqDatabase) Create(obj model.Todo) error {
	_, err := o.db.Exec("INSERT INTO TODOS (id, title, description, created_at) VALUES ($1, $2, $3, $4)", obj.Id, obj.Title, obj.Description, obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (o pqDatabase) Detail(id int32) (todo model.Todo, err error) {
	err = o.db.QueryRow("SELECT ID, TITLE, DESCRIPTION, CREATED_AT FROM TODOS WHERE ID = $1", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
	return
}

func (o pqDatabase) Delete(id int32) error {
	_, err := o.db.Exec("DELETE FROM TODOS WHERE ID = $1", string(id))
	if err != nil {
		return err
	}
	return nil
}

func (o pqDatabase) List() ([]model.Todo, error) {
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
