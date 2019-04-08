package util

import (
	"github.com/gin-gonic/gin"
)

func Add() (int, error) {
	var i = 1000
	var j = 10
	var total = i + j
	return total, nil
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
