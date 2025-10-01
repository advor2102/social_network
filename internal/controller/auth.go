package controller

import (
	"errors"
	"net/http"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/advor2102/socialnetwork/pkg"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	FullName     string `json:"full_name" db:"full_name"`
	EmployeeName string `json:"employee_name" db:"employee_name"`
	Password     string `json:"password" db:"password"`
}

// SignUp
// @Summary Create employee
// @Description Create new employee and add to database
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignUpRequest true "new employee data"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-up [post]
func (ctrl *Controller) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if err := ctrl.service.CreateEmployee(c, models.Employee{
		FullName:     input.FullName,
		EmployeeName: input.EmployeeName,
		Password:     input.Password,
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Employee created successfully"})
}

type SignIpRequest struct {
	EmployeeName string `json:"employee_name" db:"employee_name"`
	Password     string `json:"password" db:"password"`
}

type TokenPairResponse struct {
	AccessToken  string `json:"access_token_ttl_minutes"`
	RefreshToken string `json:"refresh_token_ttl_days"`
}

// SignIn
// @Summary Enter
// @Description Enter as an employee
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignIpRequest true "login and password"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-in [post]
func (ctrl *Controller) SignIn(c *gin.Context) {
	var input SignIpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	employeeID, employeeRole, err := ctrl.service.Authenticate(c, models.Employee{
		EmployeeName: input.EmployeeName,
		Password:     input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	accessToken, refreshToken, err := ctrl.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "X-Refresh_token"
)

// RefreshTokenPairs
// @Summary Refresh token pairs
// @Description Refresh token pairs
// @Tags Auth
// @Produce json
// @Param X-Refresh_token header string true "input refresh token"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/refresh [get]
func (ctrl *Controller) RefreshTokenPairs(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, employeeRole, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	accessToken, refreshToken, err := ctrl.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
