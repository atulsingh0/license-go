package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atulsingh0/license-go/pkg/generate"
)

func PostGenerate(c *gin.Context) {

	// reading and binding the post data
	inpBody := generate.Rlic{}
	err := c.ShouldBindJSON(&inpBody)

	// Validating the input data
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusNoContent,
		})
	} else if err := inpBody.InputValidation(); err != nil {
		c.AbortWithStatusJSON(http.StatusPartialContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusPartialContent,
		})
	} else {

		//generating the license
		lic, err := inpBody.Generate()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
				"code":  http.StatusUnprocessableEntity,
			})
		} else {
			c.JSON(http.StatusOK, lic)
		}

	}

}
