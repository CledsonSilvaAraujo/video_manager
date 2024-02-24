package main

import (
	"github.com/CledsonSilvaAraujo/task-manager/controller"
	"github.com/CledsonSilvaAraujo/task-manager/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":     "1",
		"alunos": "Cledson Silva",
	})
}

// func ExibeVideos(c *gin.Context) {
// 	c.JSON(200, videoController.FindAll())
// }

func SalvaVideos(c *gin.Context) {
	c.JSON(200, videoController.Save(c))
}

func main() {
	// r := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	// server.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Pong, Gin!",
	// 	})
	// })

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
