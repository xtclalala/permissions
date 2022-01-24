package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"permissions/utils"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(status int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Data:    data,
		Message: msg,
	})
}

func Ok(c *gin.Context) {
	Result(utils.SUCCESS, map[string]interface{}{}, utils.GetErrorMessage(utils.SUCCESS), c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(utils.SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(utils.SUCCESS, data, "", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(utils.ERROR, map[string]interface{}{}, msg, c)
}

func FailWhitStatusAndMessage(status int, msg string, c *gin.Context) {
	Result(status, map[string]interface{}{}, msg, c)
}

func FailWhitStatus(status int, c *gin.Context) {
	Result(status, map[string]interface{}{}, utils.GetErrorMessage(status), c)
}
