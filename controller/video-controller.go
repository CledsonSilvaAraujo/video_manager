package controller

import (
	"github.com/CledsonSilvaAraujo/video-manager/entity"
	"github.com/CledsonSilvaAraujo/video-manager/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{ // Change the receiver type to pointer
		service: service,
	}
}

func (c controller) FindAll() []entity.Video { // Change the receiver type to non-pointer
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
