package models

type AddSQLInput struct {
	DataSource string
	Adds       map[string]string
}

type AddSQLOutput struct {
	ResponseCode int
}
