package boiler

import (
	"database/sql"
	"fmt"
)

type Dao struct {
	DB *sql.DB
}

func New() *Dao {
	db, err := sql.Open("mysql", "root:@/gobyexample?parseTime=true")
	if err != nil {
		panic(fmt.Sprintln("mysql connect failed:", err.Error()))
	}
	return &Dao{DB: db}
}
