package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kerukukku1/procon30_battle_server/cmd/"
)

func main() {
	db, err := sql.Open("mysql", "root:"+os.Getenv("GOMYSQLPASSWD")+"/my_database")
	if err != nil {
		panic(err.Error())
	}
	c, err := pkg.MySQLConfigRead("config.json")
	defer db.Close()
	fmt.Printf("type is %T \n", db)
	fmt.Printf("c is %T \n", c)
}
