package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("faihsdo87tfasdb2q33fwsfdvxcvxbbsdfg")
var tokenCookie = "tkn"

func (a *Auth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie(tokenCookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization required",
			})
			c.Abort()
			return
		}

		tkn, err := jwt.Parse(jwtToken, a.keyFunc)
		if err != nil || !tkn.Valid {
			a.logger.Warn(fmt.Sprintf("Got bad jwt isvalid: %t; error on parsing; %s", tkn.Valid, err.Error()))
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
				"error":   "Authorization required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (a *Auth) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok :=  claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if time.Now().After(expirationTime) {
				return nil, fmt.Errorf("token has expired")
			}
		} else {
			return nil, fmt.Errorf("exp claim is missing in token")
		}
	}

	return SecretKey, nil
}
