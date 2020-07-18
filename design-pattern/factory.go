package pattern

import (
	"log"

	"github.com/polipopoliko/todo/todo2/storage"
)

func MemoryFactory(memoryType int) storage.Storage {
	var s storage.Storage
	switch memoryType {
	case 1:
		s = storage.NewMemory()
	case 2:
		s = new(storage.Mock)
	case 3:
		s = storage.ConnectDb("postgres", "postgres", "5432", "postgres", "password", "db", "disable")
	default:
		log.Fatal("Memory or Storage not found")
	}
	return s
}
