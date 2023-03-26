package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == http.MethodOptions {
		c.JSON(204, nil)
		return
	}

	c.Next()
}
