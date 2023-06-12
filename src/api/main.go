package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/atulsingh0/license-go/src/controllers"
	"github.com/atulsingh0/license-go/src/initializers"

	"github.com/gin-gonic/gin"
)

// initializing the app
func init() {
	initializers.LoadEnvVar()

	if strings.ToUpper(os.Getenv("ENV")) == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/", welcome)
	router.POST("/generate", controllers.PostGenerate)
	router.POST("/validate", controllers.PostValidate)

	router.Run()
}

// welcome msg
func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to License-GO",
	})
}
