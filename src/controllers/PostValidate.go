package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostValidate(c *gin.Context) {

	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"msg": "work to do",
	})
}
