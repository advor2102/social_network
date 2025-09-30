package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
	EmployeeID int `json:"employee_ID"`
	IsRefresh bool `json:"is_refresh"`
}

func GenerateToken(employeeID int, ttl int, isRefresh bool) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{},
		EmployeeID: employeeID,
		IsRefresh: isRefresh,
	}

	if isRefresh{
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * 24 * time.Hour)
	} else {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * time.Minute)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (employeeID int, isRefresh bool, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, false, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.EmployeeID, claims.IsRefresh, nil
	}

	return 0, false, fmt.Errorf("invalid token")
}
