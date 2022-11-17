package models

type UpdateSQLInput struct {
	DataSource string
	Updates    map[string]string
	Condition  map[string]string
}

type UpdateSQLOutput struct {
	ResponseCode int
}
