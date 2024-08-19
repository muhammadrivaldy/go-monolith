package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goutil "github.com/muhammadrivaldy/go-util"
)

func (e endpoint) HealthCheck(c *gin.Context) { goutil.ResponseOK(c, http.StatusOK, nil) }
