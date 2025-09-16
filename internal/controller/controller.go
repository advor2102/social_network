package controller

import (
	"github.com/advor2102/socialnetwork/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	router  *gin.Engine
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}
