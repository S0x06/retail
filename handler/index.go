package handler

import (
	"github.com/gin-gonic/gin"
)

// Name will print hello name
// @Summary Print
// @Accept json
// @Tags Name
// @Security Bearer
// @Produce  json
// @Param name path string true "name"
// @Resource Name
// @Router /hello/{name} [get]
// @Success 200 {object} main.Message

func Index(c *gin.Context) {

	// 返回JSON数据
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
