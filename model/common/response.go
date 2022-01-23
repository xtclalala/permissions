package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"permissions/utils"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: utils.GetErrorMessage(code),
	})
}

func Ok(c *gin.Context) {
	Result(utils.SUCCESS, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(utils.SUCCESS, data, c)
}

func Fail(c *gin.Context) {
	Result(utils.ERROR, map[string]interface{}{}, c)
}

func FailWhitMessage(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, c)
}
