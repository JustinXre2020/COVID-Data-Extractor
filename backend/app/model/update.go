package model

import (
	"fmt"
	"strconv"
)

func UpdateSQL(dataSource string, updates map[string]string, conditions map[string]string) error {
	query := "update "
	query += dataSource
	query += " set "
	for key, value := range updates {
		_, err := strconv.Atoi(value)
		if err != nil {
			//value is string
			query += key + "=" + "'" + value + "'"
			query += " ,"
		} else {
			query += key + "=" + value
			query += " ,"
		}
	}
	query = query[:len(query)-2]

	if len(conditions) != 0 {
		query += " where "
		for key, value := range conditions {
			query += key + "=" + "'" + value + "'"
			query += " and "
		}
		query = query[:len(query)-5]
	}
	fmt.Printf(query)
	if _, err := Db.Exec(query); err != nil {
		return err
	}
	return nil
}
