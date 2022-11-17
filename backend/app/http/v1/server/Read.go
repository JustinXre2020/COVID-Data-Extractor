package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Read(c *gin.Context) {
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
	result, err := model.ReadSQL(dataSource, updates, condition)
	if err != nil {
		response.Error(c, "Fail to search with error:"+err.Error())
		return
	} else {
		jsonResult, _ := json.Marshal(result)
		fmt.Println(string(jsonResult))
		//jsonResult, err := json.Marshal(result)
		responseData := &models.AdvanceQueryOneOutPut{
			ResponseCode:   200,
			ResponseResult: string(jsonResult),
		}
		response.Success(c, "Success", responseData)
		return
	}
}
