package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetAllUsers(c *gin.Context) {
	users, err := controller.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (controller *Controller) GetUserByID(c *gin.Context) {

}

func (controller *Controller) CreateUser(c *gin.Context) {

}

func (controller *Controller) UpdateUserByID(c *gin.Context) {

}

func (controller *Controller) DeleteUserbyID(c *gin.Context) {

}
