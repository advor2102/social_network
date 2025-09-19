package controller

import (
	"net/http"

	_ "github.com/advor2102/socialnetwork/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (controller *Controller) RegisterEndpoints() {
	controller.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controller.router.GET("/ping", controller.Ping)

	controller.router.GET("/users", controller.GetAllUsers)
	controller.router.GET("/users/:id", controller.GetUserByID)
	controller.router.POST("/users", controller.CreateUser)
	controller.router.PUT("/users/:id", controller.UpdateUserByID)
	controller.router.DELETE("/users/:id", controller.DeleteUserByID)
}

// Ping
// @Summary Health-check
// @Description Check of the service
// @Tags Ping
// @Produce json
// @Success 200 {object} models.User
// @Router /ping [get]
func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, CommonResponse{Message: "Service's up and running"})
}

func (controller *Controller) RunServer(address string) error {
	controller.RegisterEndpoints()

	if err := controller.router.Run(address); err != nil {
		return err
	}
	return nil
}
