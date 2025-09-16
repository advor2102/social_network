package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) RegisterEndpoints() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)

	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetUserByID)
	r.POST("/users/", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUserByID)
	r.DELETE("/users/:id", controller.DeleteUserbyID)
}

func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
