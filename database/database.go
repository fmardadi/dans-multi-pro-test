package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB
var err error

func Connect() {
	DB, err = sql.Open("mysql", "root@tcp(localhost:3306)/dans_multi_pro")
	if err != nil {
		fmt.Printf("Fail connected to Database!, error: %s", err.Error())
	}

	fmt.Println("Connected to Database!")
}
