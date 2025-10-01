package controller

import (
	"net/http"

	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/advor2102/socialnetwork/pkg"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	employeeIDCtx       = "employeeID"
	employeeRoleCtx     = "employeeRole"
)

func (ctrl *Controller) checkUserAuthentication(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, employeeRole, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	c.Set(employeeIDCtx, employeeID)
	c.Set(employeeRoleCtx, string(employeeRole))
}

func (ctrl *Controller) checkIsAdmin(c *gin.Context) {
	role := c.GetString(employeeRoleCtx)
	if role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "role is not in context"})
		return
	}

	if role != models.RoleAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, CommonError{Error: "permission denied"})
		return
	}

	c.Next()
}
