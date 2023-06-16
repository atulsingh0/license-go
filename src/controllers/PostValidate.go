package controllers

import (
	"net/http"

	"github.com/atulsingh0/license-go/src/validate"
	"github.com/gin-gonic/gin"
)

func PostValidate(c *gin.Context) {

	// reading and binding the post data
	inpBody := validate.Slic{}
	c.ShouldBindJSON(&inpBody)

	// Validating the input License
	err := inpBody.Validate()

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
			"code":  http.StatusForbidden,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "License is valid.",
		})
	}
}
