package database

import (
	"log"
)

func CreateTodo(todo Todo) {
	if err := GormDB.Create(&todo).Error; err != nil {
		log.Fatal(err)
	}
}

func GetAllTodos() []string {
	var todos []Todo
	var res []string
	if err := GormDB.Find(&todos).Error; err != nil {
		log.Fatal(err)
		return nil
	}
	for _, todo := range todos {
		res = append(res, todo.Name)
	}
	return res
}

// Added for integration tests inserts
func RemoveLastTodo() {
	var todo Todo
	if err := GormDB.Last(&todo).Error; err != nil {
		log.Fatal(err)
	} else {
		if err := GormDB.Delete(&todo).Error; err != nil {
			log.Fatal(err)
		}
	}
}