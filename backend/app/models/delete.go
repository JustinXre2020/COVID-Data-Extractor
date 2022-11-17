package models

// DeleteSQLInput SQL通用字段  delete from [1] where [2]
// 1:data source=> [tableName:value]
// 2:condition=> [conditionName:value]

type DeleteSQLInput struct {
	DataSource string
	Condition  map[string]string
}

type DeleteSQLOutput struct {
	ResponseCode int
}
