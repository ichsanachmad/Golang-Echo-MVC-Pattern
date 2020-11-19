package utils

import (
	"Golang-Echo-MVC-Pattern/constant"
	"Golang-Echo-MVC-Pattern/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func GetDataToken(c echo.Context) *entity.MyCustomClaims {
	tokenID := getAuthorizationToken(c)
	tokenKey := os.Getenv(constant.JWTKey)
	token, _ := getParseWithClaimsJWT(tokenKey, tokenID)
	claims, _ := token.Claims.(*entity.MyCustomClaims)
	return claims
}

func IsTokenValid(c echo.Context) bool {
	tokenID := getAuthorizationToken(c)
	tokenKey := os.Getenv(constant.JWTKey)
	token, err := getParseWithClaimsJWT(tokenKey, tokenID)

	if err != nil {
		return false
	}

	_, ok := token.Claims.(*entity.MyCustomClaims)
	if !ok {
		return false
	}

	return true
}

func getAuthorizationToken(c echo.Context) string {
	tokenID := c.Request().Header.Get(constant.Authorization)
	tokenID = TrimString(tokenID, constant.Bearer)
	tokenID = TrimString(tokenID, " ")
	return tokenID
}

func getParseWithClaimsJWT(tokenKey string, tokenID string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenID, &entity.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !IsTokenValid(c) {
			return c.JSON(http.StatusUnauthorized, GetResNoData(constant.StatusError, constant.MessageFailedJWT))
		}
		return next(c)
	}
}
