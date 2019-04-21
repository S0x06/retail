package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"retail/pkg/errno"
	"retail/pkg/version"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//成功返回
func SendResponse(c *gin.Context) {
	// code, message := errno.DecodeErr(err)

	var queries map[string]interface{} = map[string]interface{}{
		"rel":    "search",
		"href":   "dfg",
		"prompt": "dfg",
		"data":   "ghj",
	}

	var template map[string]interface{} = map[string]interface{}{
		"rel":    "search",
		"href":   "dg",
		"prompt": "Search",
		"data":   "sdf",
	}

	var data map[string]interface{} = map[string]interface{}{
		"version":  version.TagInfo.String(),
		"href":     "abc",
		"links":    "DF",
		"items":    "sss",
		"queries":  queries,
		"template": template,
	}

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "OK",
		Data:    data,
	})
}

//失败返回
func SendErrorResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
