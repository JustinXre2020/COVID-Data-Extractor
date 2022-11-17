package server

import (
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Update(c *gin.Context) {
	//if err := c.ShouldBind(&models.UpdateSQLInput{}); err != nil {
	//	fmt.Println(err.Error())
	//	response.Error(c, "wrong parameter")
	//}
	dataSource := c.PostForm("dataSource")
	updatesJson := c.PostForm("updates")
	conditionJson := c.PostForm("condition")

	updates := make(map[string]string)
	condition := make(map[string]string)
	if len(updatesJson) != 0 {
		updates = model.ConvertStringToMap(updatesJson)
	}
	if len(conditionJson) != 0 {
		condition = model.ConvertStringToMap(conditionJson)
	}

	//execute sql query
	err := model.UpdateSQL(dataSource, updates, condition)
	if err != nil {
		response.Error(c, "Fail to update query with error:"+err.Error())
		return
	}
	responseData := &models.UpdateSQLOutput{
		ResponseCode: 200,
	}
	response.Success(c, "Success", responseData)
}
