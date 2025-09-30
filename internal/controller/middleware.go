package controller

import (
	"net/http"
	"strings"

	"github.com/advor2102/socialnetwork/pkg"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	employeeIDCtx       = "employeeID"
)

func (ctrl *Controller) checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "empty authorization header"})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "invalid authorization header"})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "empty token"})
	}

	token := headerParts[1]
	employeeID, isRefresh, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token",})
		return
	}

	c.Set(employeeIDCtx, employeeID)
}
