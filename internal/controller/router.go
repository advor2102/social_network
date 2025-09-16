package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) RegisterEndpoints() {
	controller.router.GET("/ping", controller.Ping)

	controller.router.GET("/users", controller.GetAllUsers)
	controller.router.GET("/users/:id", controller.GetUserByID)
	controller.router.POST("/users/", controller.CreateUser)
	controller.router.PUT("/users/:id", controller.UpdateUserByID)
	controller.router.DELETE("/users/:id", controller.DeleteUserbyID)
}

func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (controller *Controller) RunServer(address string) error {
	controller.RegisterEndpoints()

	if err := controller.router.Run(address); err != nil {
		return err
	}
	return nil
}
