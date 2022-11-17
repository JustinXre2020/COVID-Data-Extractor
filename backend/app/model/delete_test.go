package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Delete(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	//Based on FYP.connection tables
	dataSource := "users"
	conditions := make(map[string]string)
	conditions["id"] = "1"
	DeleteSQL(dataSource, conditions)
}
