package controller

import (
	"net/http"

	"github.com/advor2102/socialnetwork/pkg"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	employeeIDCtx       = "employeeID"
)

func (ctrl *Controller) checkUserAuthentication(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	c.Set(employeeIDCtx, employeeID)
}
