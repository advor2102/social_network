package controller

import "github.com/advor2102/socialnetwork/internal/service"

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}
