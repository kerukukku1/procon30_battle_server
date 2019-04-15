package pkg

import (
	"encoding/json"
	"io/ioutil"
)

type MySQLConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func MySQLConfigRead(filename string) (*MySQLConfig, error) {
	c := new(MySQLConfig)

	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(jsonStr, c)
	if err != nil {
		return c, err
	}

	return c, nil
}
