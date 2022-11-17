package models

type AdvanceQueryTwoInput struct {
	Region             string
	Smoking_prevalence float64
}

type AdvanceQueryTwoOutPut struct {
	ResponseCode   int
	ResponseResult string
}
