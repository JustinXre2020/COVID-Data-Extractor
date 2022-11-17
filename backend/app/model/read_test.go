package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Read(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	//Based on covid_data.Health tables
	dataSource := "Health"
	conditions := make(map[string]string)
	conditions["Region"] = "AQ"
	updates := make(map[string]string)
	updates["Region"] = "2"
	updates["Life_expectancy"] = "20"
	ReadSQL(dataSource, updates, conditions)
}
