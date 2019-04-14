package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/my_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Printf("type is %T \n", db)
}
