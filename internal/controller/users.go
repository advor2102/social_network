package controller

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetAllUsers").Logger()
	employeeID := c.GetInt(employeeIDCtx)
	if employeeID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid employeeID in context"})
		return
	}

	logger.Debug().Int("employee_id", employeeID).Msg("GetUser")

	users, err := controller.service.GetAllUsers()
	if err != nil {
		controller.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID
// @Summary Get user by ID
// @Description Get user's data by ID
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.User
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (controller *Controller) GetUserByID(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetUserByID").Logger()
	employeeID := c.GetInt(employeeIDCtx)
	if employeeID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid employeeID in context"})
		return
	}

	logger.Debug().Int("employee_id", employeeID).Msg("GetUser")

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id < 1 {
		controller.handleError(c, errs.ErrInvalidUserID)
		return
	}

	user, err := controller.service.GetUserByID(id)
	if err != nil {
		controller.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

type CreateUserRequest struct {
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Age      int    `json:"age" db:"age"`
}

// CreateUser
// @Summary Create user
// @Description Create new user and add to database
// @Tags Users
// @Consume json
// @Produce json
// @Param request_body body CreateUserRequest true "new user data"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users [post]
func (controller *Controller) CreateUser(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.CreateUser").Logger()
	employeeID := c.GetInt(employeeIDCtx)
	if employeeID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid employeeID in context"})
		return
	}

	logger.Debug().Int("employee_id", employeeID).Msg("GetUser")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		controller.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if user.UserName == "" || user.Age < 0 || user.Email == "" {
		controller.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := controller.service.CreateUser(user); err != nil {
		controller.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully"})
}

// UpdateUserByID
// @Summary Update user by ID
// @Description Update user's data by ID
// @Tags Users
// @Consume json
// @Produce json
// @Param id path int true "user id"
// @Param request_body body CreateUserRequest true "updated user data"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (controller *Controller) UpdateUserByID(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.UpdateUserByID").Logger()
	employeeID := c.GetInt(employeeIDCtx)
	if employeeID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid employeeID in context"})
		return
	}

	logger.Debug().Int("employee_id", employeeID).Msg("GetUser")

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id < 1 {
		controller.handleError(c, errs.ErrInvalidUserID)
		return
	}

	var user models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		controller.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if user.UserName == "" || user.Age < 0 || user.Email == "" {
		controller.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	user.ID = id

	if err = controller.service.UpdateUserByID(user); err != nil {
		controller.handleError(c, err)
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "User updated successfully"})
}

// DeleteUserByID
// @Summary Delete user by ID
// @Description Delete user's data by ID
// @Tags Users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [delete]
func (controller *Controller) DeleteUserByID(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.UpdateUserByID").Logger()
	employeeID := c.GetInt(employeeIDCtx)
	if employeeID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid employeeID in context"})
		return
	}

	logger.Debug().Int("employee_id", employeeID).Msg("GetUser")

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id < 1 {
		controller.handleError(c, errs.ErrInvalidUserID)
		return
	}

	if err = controller.service.DeleteUserByID(id); err != nil {
		controller.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "User deleted successfully"})
}
