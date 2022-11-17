//package model
//
//import (
//	"fmt"
//	"pro/app/models"
//)
//
//func ReadSQL(dataSource string, updates map[string]string, conditions map[string]string) ([]models.Patients, error) {
//	query := "select "
//
//	for key, _ := range updates {
//
//		//value is string
//		query += key
//		query += " ,"
//
//	}
//	query = query[:len(query)-2]
//	query += " from "
//	query += dataSource
//
//	if len(conditions) != 0 {
//		query += " where "
//		for key, value := range conditions {
//			query += key + "=" + "'" + value + "'"
//			query += " and "
//		}
//		query = query[:len(query)-5]
//	}
//	var result = make([]models.Patients, 0)
//	readQuery := models.Patients{}
//	fmt.Println(query)
//	rows, err := Db.Query(query)
//
//	for rows.Next() {
//		err = rows.Scan(&readQuery.PatientName, &readQuery.PatientId, &readQuery.Region)
//		checkErr(err)
//		result = append(result, readQuery)
//	}
//	fmt.Printf(query)
//	return result, nil
//}
package model

import "fmt"

func ReadSQL(dataSource string, updates map[string]string, conditions map[string]string) ([]map[string]string, error) {
	query := "select "

	for key, _ := range updates {

		//value is string
		query += key
		query += " ,"

	}
	query = query[:len(query)-2]
	query += " from "
	query += dataSource

	if len(conditions) != 0 {
		query += " where "
		for key, value := range conditions {
			query += key + "=" + "'" + value + "'"
			query += " and "
		}
		query = query[:len(query)-5]
	}
	query += " limit 15"
	fmt.Printf(query)
	rows, err := Db.Query(query)

	if _, err := Db.Query(query); err != nil {
		return nil, err
	}

	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns() //获取列的信息
	count := len(columns)        //列的数量
	fmt.Println(count)
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]string, 0) //创建返回值：不定长的map类型切片
	for rows.Next() {
		err := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := make(map[string]string) //用于存放1列的 [键/值] 对
		if err != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			b, _ := raw_value.([]byte)
			v := string(b) //将raw数据转换成字符串
			m[colName] = v //colName是键，v是值
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	defer rows.Close()

	if len(ret) != 0 {
		return ret, nil
	}
	return nil, nil
}
