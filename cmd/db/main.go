package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func main() {
	db, err := sql.Open("mysql", "root:"+os.Getenv("GOMYSQLPASSWD")+"/my_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Printf("type is %T \n", db)
}
