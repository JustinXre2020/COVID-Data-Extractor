package model

import (
	"database/sql"
	"fmt"
	"log"
	"pro/config"
)

var Db *sql.DB

type health struct {
	region              string
	Life_expectancy     float64
	Smoking_prevalence  float64
	Diabetes_prevalenve float64
	Health_expenditure  float64
}

//gorm链接数据库-本地数据库
func Run() error {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			config.Mysql.UserName,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.Database,
		),
	)
	if err != nil {
		return err
	}
	/*连接池信息*/
	//db.DB().SetMaxIdleConns(config.Mysql.MaxIdleConns) //设置最大空闲数
	//db.DB().SetMaxOpenConns(config.Mysql.MaxOpenConns) //设置最大连接数
	//db.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "ww_" + defaultTableName
	//}
	if err != nil {
		fmt.Printf("DSN[%s]格式错误: %v \n", config.Mysql.Database, err) // 如果出现错误，则DSN格式错误
		return err
	}

	err = db.Ping() // 尝试连接mysql

	if err != nil {
		fmt.Printf("连接数据库失败: %v \n", err)
		return err
	}

	fmt.Println("连接数库成功")

	Db = db

	var u health
	sqlStr := "select * from Health where Region = ?"                                                                                         // 构建查询语句
	err = db.QueryRow(sqlStr, "AD").Scan(&u.region, &u.Life_expectancy, &u.Smoking_prevalence, &u.Diabetes_prevalenve, &u.Health_expenditure) // 查询单条数据并扫描结果

	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("未查询到结果!")
			return nil
		}
		log.Printf("查询失败: %s", err)
		return nil
	}

	log.Printf("查询结果：%#v", u)

	return nil
}

//TCP Connection
func RunCloud() error {
	dbUser := "root"           // e.g. 'my-db-user'
	dbPwd := "schoolsucks1234" // e.g. 'my-db-password'
	dbName := "covid_data"     // e.g. 'my-database'
	dbPort := "3306"
	dbTCPHost := "35.188.29.29"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}
	err = dbPool.Ping()
	if err != nil {
		return err
	}
	Db = dbPool
	return nil
}

//Unix Connection
//func RunCloud() error {
//	dbUser := "root"                                                                 // e.g. 'my-db-user'
//	dbPwd := "schoolsucks1234"                                                       // e.g. 'my-db-password'
//	dbName := "covid_data"                                                           // e.g. 'my-database'
//	unixSocketPath := "/cloudsql/covid-19-data-extractor:us-central1:is411teamrocks" // e.g. '/cloudsql/project:region:instance'
//
//	dbURI := fmt.Sprintf("%s:%s@unix(/%s)/%s?parseTime=true",
//		dbUser, dbPwd, unixSocketPath, dbName)
//
//	// dbPool is the pool of database connections.
//	dbPool, err := sql.Open("mysql", dbURI)
//	if err != nil {
//		return fmt.Errorf("sql.Open: %v", err)
//	}
//	err = dbPool.Ping()
//	if err != nil {
//		return err
//	}
//	Db = dbPool
//	return nil
//}
