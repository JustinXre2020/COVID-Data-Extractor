package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Update(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	//Based on covid_data.Health tables
	dataSource := "Health"
	conditions := make(map[string]string)
	conditions["Region"] = "AD"
	updates := make(map[string]string)
	updates["Smoking_prevalence"] = "15"
	updates["Diabetes_prevalenve"] = "20"
	UpdateSQL(dataSource, updates, conditions)
}
