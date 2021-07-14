package database

import (
	"gorm.io/gorm"
)

type Todo struct {
	Id   uint `gorm:"uniqueIndex;autoIncrement:true"`
	Name string
	gorm.Model
}

func (Todo) TableName() string {
	return "Todo"
}
