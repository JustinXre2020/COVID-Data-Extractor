package server

import (
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/models"
)

//Visitor Section
func Test(c *gin.Context) {
	responseData := &models.TestOutput{
		ResponseCode:  200,
		ResponseValue: "Success",
	}
	response.Success(c, "Success", responseData)
}
