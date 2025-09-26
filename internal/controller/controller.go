package controller

import (
	"errors"
	"net/http"

	"github.com/advor2102/socialnetwork/internal/contracts"
	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	router  *gin.Engine
	service contracts.ServiceI
}

func NewController(service contracts.ServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}
}

func (controller *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotFound) || errors.Is(err, errs.ErrEmployeeNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) || errors.Is(err, errs.ErrEmployeeNameAlreadyExist):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectEmployeeNameOrPassword):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
