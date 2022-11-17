package models

type ReadSQLInput struct {
	DataSource string
	Updates    map[string]string
	Condition  map[string]string
}

type Patients struct {
	PatientName string
	PatientId   int
	Region      string
}

type Hospitalizations struct {
	Region                          string
	Date                            string
	Current_ventilator_patients     float64
	Current_hospitalized_patients   float64
	Current_intensive_care_patients float64
}

type ReadSQLOutput struct {
	ResponseCode   int
	ResponseResult string
}
