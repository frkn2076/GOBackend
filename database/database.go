package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgreDB *sql.DB
var GormDB *gorm.DB

func init() {
	PostgreDB = initPostgreDB()
	GormDB = initGorm()
	
	GormDB.AutoMigrate(&Todo{})
}

func initPostgreDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://uSeR1:12345@localhost:5432/SampleDB?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func initGorm() *gorm.DB {
	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: PostgreDB,
		}),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	return gormDB
}
