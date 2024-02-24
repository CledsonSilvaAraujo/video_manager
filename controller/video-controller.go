package controller

import (
	"github.com/CledsonSilvaAraujo/video-manager/entity"
	"github.com/CledsonSilvaAraujo/video-manager/service"
	"github.com/CledsonSilvaAraujo/video-manager/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()

	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{ // Change the receiver type to pointer
		service: service,
	}
}

func (c controller) FindAll() []entity.Video { // Change the receiver type to non-pointer
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) error {
	var video entity.Video
	err := context.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
