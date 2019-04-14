package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kerukukku1/procon30_battle_server/internal/pkg"
)

func main() {
	c, err := pkg.MySQLConfigRead("config.json")
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open("mysql", c.Username+":"+c.Password+"@/"+c.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
