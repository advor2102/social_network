package controller

import (
	"errors"
	"net/http"

	"github.com/advor2102/socialnetwork/internal/configs"
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

type SignIpResponse struct {
	AccessToken  string `json:"access_token_ttl_minutes"`
	RefreshToken string `json:"refresh_token_ttl_days"`
}

func (ctrl *Controller) SignIn(c *gin.Context) {
	var input SignIpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	accessToken, refreshToken, err := ctrl.service.Authenticate(c, models.Employee{
		EmployeeName: input.EmployeeName,
		Password:     input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, SignIpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "X-Refresh_token"
)

func (ctrl *Controller) RefreshTokenPairs(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	accessToken, err := pkg.GenerateToken(employeeID, configs.AppSettings.AuthParams.AccessTokenTtlMinutes, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: errs.ErrSomethingWentWrong.Error()})
		return
	}

	refreshToken, err := pkg.GenerateToken(employeeID, configs.AppSettings.AuthParams.RefreshTokenTtlDays, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{Error: errs.ErrSomethingWentWrong.Error()})
		return
	}

	c.JSON(http.StatusOK, SignIpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
