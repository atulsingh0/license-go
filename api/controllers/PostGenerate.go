package controllers

import (
	"net/http"

	"github.com/atulsingh0/license-go/src/generate"

	"github.com/gin-gonic/gin"
)

func PostGenerate(c *gin.Context) {

	// reading and binding the post data
	inpBody := generate.Rlic{}
	c.ShouldBindJSON(&inpBody)

	// Validating the input data
	err := inpBody.InputValidation()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPartialContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusPartialContent,
		})
	} else {

		//generating the license
		lic, err := inpBody.Generate()

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
				"code":  http.StatusUnprocessableEntity,
			})
		} else {
			c.JSON(http.StatusOK, lic)
		}

	}

}
