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

	authG := controller.router.Group("/auth")
	{
		authG.POST("/sign-up", controller.SignUp)
		authG.POST("/sign-in", controller.SignIn)
		authG.GET("/refresh", controller.RefreshTokenPairs)
	}

	apiG := controller.router.Group("/api", controller.checkUserAuthentication)
	{
		apiG.GET("/users", controller.GetAllUsers)
		apiG.GET("/users/:id", controller.GetUserByID)
		apiG.POST("/users", controller.checkIsAdmin, controller.CreateUser)
		apiG.PUT("/users/:id", controller.checkIsAdmin, controller.UpdateUserByID)
		apiG.DELETE("/users/:id", controller.checkIsAdmin, controller.DeleteUserByID)
	}
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
