package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func AdvanceTwo(c *gin.Context) {

	result, err := model.AdvanceQueryTwoSQL()
	if err != nil {
		response.Error(c, "Fail to execute query 2 with error:"+err.Error())
		return
	}
	jsonResult, err := json.Marshal(result)
	responseData := &models.AdvanceQueryOneOutPut{
		ResponseCode:   200,
		ResponseResult: string(jsonResult),
	}
	response.Success(c, "Success", responseData)
}
