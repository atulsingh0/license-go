package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atulsingh0/license-go/pkg/generate"
	"github.com/atulsingh0/license-go/pkg/storage"
)

func PostGenerate(c *gin.Context) {

	// reading and binding the post data
	inpBody := generate.Rlic{}

	// Validating the input data
	if err := c.ShouldBindJSON(&inpBody); err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusNoContent,
		})
		return
	}

	if err := inpBody.InputValidation(); err != nil {
		c.AbortWithStatusJSON(http.StatusPartialContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusPartialContent,
		})
		return
	}

	//generating the license
	sl, lic, err := inpBody.Generate()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
			"code":  http.StatusUnprocessableEntity,
		})
		return
	}

	// Writing to Plugins
	if err = storage.Plugins(sl, lic); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
			"code":  http.StatusUnprocessableEntity,
		})
		return
	}
	c.JSON(http.StatusOK, lic)

}
