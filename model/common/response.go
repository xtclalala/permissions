package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"permissions/utils"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func Result(status int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Result:  data,
		Message: msg,
	})
}

func Ok(c *gin.Context) {
	Result(utils.SUCCESS, map[string]any{}, utils.GetErrorMessage(utils.SUCCESS), c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(utils.SUCCESS, map[string]any{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(utils.SUCCESS, data, utils.GetErrorMessage(utils.SUCCESS), c)
}

func Fail(c *gin.Context) {
	Result(utils.ERROR, map[string]any{}, utils.GetErrorMessage(utils.ERROR), c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(utils.ERROR, map[string]any{}, msg, c)
}

func FailWhitStatusAndMessage(status int, msg string, c *gin.Context) {
	Result(status, map[string]any{}, msg, c)
}

func FailWhitStatus(status int, c *gin.Context) {
	Result(status, map[string]any{}, utils.GetErrorMessage(status), c)
}

type PageVO struct {
	Items any   `json:"items"`
	Total int64 `json:"total"`
}
