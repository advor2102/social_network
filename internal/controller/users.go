package controller

import (
	"net/http"
	"strconv"

	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllUsers
// @Summary Get all users
// @Description Get list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} CommonError
// @Router /users [get]
func (controller *Controller) GetAllUsers(c *gin.Context) {
	users, err := controller.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID
// @Summary Get user by ID
// @Description Get user's data by IF
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (controller *Controller) GetUserByID(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
		return
	}

	user, err := controller.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser
// @Summary Create user
// @Description Create new user and add to database
// @Tags Users
// @Produce json
// @Consume json
// @Param request_body body models.User true "new user data"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
func (controller *Controller) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
		return
	}

	if err := controller.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully"})
}

func (controller *Controller) UpdateUserByID(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ID = id

	if err = controller.service.UpdateUserByID(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (controller *Controller) DeleteUserByID(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = controller.service.DeleteUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
