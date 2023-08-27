package main

import (
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/datagenx/license-generator/api/controllers"
	"github.com/datagenx/license-generator/internal/initializers"
)

func main() {

	// initilization
	runtime.GOMAXPROCS(runtime.NumCPU())
	initializers.LoadEnvVar()

	if strings.ToUpper(os.Getenv("ENV")) == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	//

	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/", welcome)
	router.POST("/generate", controllers.PostGenerate)
	router.POST("/validate", controllers.PostValidate)

	router.Run()
}

// welcome msg
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to License-GO",
	})
}
