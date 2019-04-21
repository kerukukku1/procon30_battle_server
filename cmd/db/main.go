package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kerukukku1/procon30_battle_server/internal/pkg"
)

func main() {
	c, err := pkg.MySQLConfigRead("../../config.json")
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open("mysql", c.Username+":"+c.Password+"@tcp(localhost:3306)/"+c.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var token string
		if err := rows.Scan(&id, &token, &name); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, token, name)
	}
}
