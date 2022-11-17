package server

import (
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
)

func Add(c *gin.Context) {
	//if err := c.ShouldBind(&models.UpdateSQLInput{}); err != nil {
	//	fmt.Println(err.Error())
	//	response.Error(c, "wrong parameter")
	//}
	dataSource := c.PostForm("dataSource")
	addJson := c.PostForm("updates")

	adds := make(map[string]string)

	if len(addJson) != 0 {
		adds = model.ConvertStringToMap(addJson)
	}

	//execute sql query
	err := model.AddSQL(dataSource, adds)
	if err != nil {
		response.Error(c, "Fail to add query with error:"+err.Error())
		return
	}
	responseData := &models.AddSQLOutput{
		ResponseCode: 200,
	}
	response.Success(c, "Success", responseData)
}
