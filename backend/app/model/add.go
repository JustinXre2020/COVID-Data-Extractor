package model

import (
	"fmt"
	"strconv"
)

func AddSQL(dataSource string, adds map[string]string) error {
	query := "insert into "
	query += dataSource
	query += " ("
	for key, _ := range adds {
		_, err := strconv.Atoi(key)
		if err != nil {
			//value is string
			query += key
			query += " ,"

		}
	}

	query = query[:len(query)-2]
	query += ")"
	query += " values "
	query += "("
	for _, value := range adds {
		_, err := strconv.Atoi(value)
		if err != nil {
			//value is string
			query += "'" + value + "'"
			query += " ,"
		} else {
			query += value
			query += " ,"
		}
	}
	query = query[:len(query)-2]
	query += ")"
	fmt.Printf(query)

	if _, err := Db.Exec(query); err != nil {
		return err
	}
	return nil
}
