package main

import (
	"io"
	"os"

	"github.com/CledsonSilvaAraujo/video-manager/controller"
	"github.com/CledsonSilvaAraujo/video-manager/middlewares"
	"github.com/CledsonSilvaAraujo/video-manager/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":8080")
}
