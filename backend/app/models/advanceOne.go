package models

type AdvanceQueryOneInput struct {
	Count  int
	Region string
}

type AdvanceQueryOneOutPut struct {
	ResponseCode   int
	ResponseResult string
}
