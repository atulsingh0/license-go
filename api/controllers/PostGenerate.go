package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/atulsingh0/license-go/pkg/generate"
)

type GeneratorAPI struct {
	Destinations []string
}

func (api GeneratorAPI) PostGenerate(c *gin.Context) {
	// reading and binding the post data
	inpBody := generate.Rlic{}

	if err := c.ShouldBindJSON(&inpBody); err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusNoContent,
		})
		return
	}

	// Validating the input data
	if err := inpBody.InputValidation(); err != nil {
		c.AbortWithStatusJSON(http.StatusPartialContent, gin.H{
			"error": err.Error(),
			"code":  http.StatusPartialContent,
		})
		return
	}

	// generating the license
	sl, lic, err := inpBody.Generate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
			"code":  http.StatusUnprocessableEntity,
		})
		return
	}

	for _, dst := range api.Destinations {
		if err := SelectStorage(dst).StoreLicense(sl, lic); err != nil {
			// Handle it and maybe abort? But you decide what you want to do!
		}
	}

	// // Writing to Plugins
	// if err := storage.Plugins(sl, lic); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": err.Error(),
	// 		"code":  http.StatusUnprocessableEntity,
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, lic)
}

type Storer interface {
	StoreLicense(slic generate.Slic, license string) error
}

type MongoDB struct {
	// ...
}

func (db MongoDB) StoreLicense(slic generate.Slic, license string) error {
	// TODO: NOT IMPLEMENTED
	return nil
}

func SelectStorage(dst string) Storer {
	switch dst {
	case "mongdb":
		return MongoDB{}
	}
}
