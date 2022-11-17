package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Add(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	//Based on FYP.connection tables
	dataSource := "Health"
	adds := make(map[string]string)
	adds["Region"] = "QQ"
	adds["Life_expectancy"] = "27"

	AddSQL(dataSource, adds)
}
