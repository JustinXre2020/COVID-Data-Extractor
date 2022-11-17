package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Delete(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteSQLInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "wrong parameter")
	}
	dataSource := c.PostForm("dataSource")
	conditionJson := c.PostForm("condition")

	condition := model.ConvertStringToMap(conditionJson)
	//execute sql query
	err := model.DeleteSQL(dataSource, condition)
	if err != nil {
		response.Error(c, "Fail to delete query with error:"+err.Error())
		return
	}
	responseData := &models.DeleteSQLOutput{
		ResponseCode: 200,
	}
	response.Success(c, "Success", responseData)
}
