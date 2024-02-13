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
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong, Gin!",
		})
	})
	r.GET("/alunos", ExibeTodosAlunos)
	// r.GET("/videos", ExibeVideos)
	r.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	r.POST("/videos", SalvaVideos)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
