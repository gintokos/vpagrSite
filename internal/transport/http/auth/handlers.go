package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/serverdb/pkg/logger"
	"github.com/gintokos/vpagrSite/internal/services/auth"
	"github.com/golang-jwt/jwt"
)

var ttlJWTtoken = 3600 * 72

type Auth struct {
	logger logger.CustomLogger
}

func InitRoutes(eng *gin.Engine, logger logger.CustomLogger) Auth {
	a := Auth{
		logger: logger,
	}

	eng.GET("/login/*any", a.HandleLogin)

	logger.Info("Auth handlers was inited")

	return a
}

func (a *Auth) HandleLogin(c *gin.Context) {
	usertkn := c.Param("any")[1:]
	id, err := auth.Auth.Login(c.Request.Context(), usertkn)
	if err != nil {
		a.logger.Error("got error on login", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}

	jwtToken, err := a.generateToken(id)
	if err != nil {
		a.logger.Error("error on generatring JWTtoken, error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()
		return
	}

	c.SetCookie(tokenCookie, jwtToken, ttlJWTtoken, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"path":    usertkn,
	})
	a.logger.Info("User succesfully authorithed")
}

func (a *Auth) generateToken(id int64) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(ttlJWTtoken)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SecretKey)
}
