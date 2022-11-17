package model

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ConvertStringToMap(secretString string) map[string]string {
	secretMap := map[string]string{}
	if err := json.Unmarshal([]byte(secretString), &secretMap); err != nil {
		panic(err)
	}
	return secretMap
}
