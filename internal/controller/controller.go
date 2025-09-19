package controller

import (
	"errors"
	"net/http"

	"github.com/advor2102/socialnetwork/internal/errs"
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

func (controller *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
