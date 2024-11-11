package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/serverdb/pkg/logger"
)

type Auth struct {
	logger logger.CustomLogger
}

func InitHanlers(eng *gin.Engine, logger logger.CustomLogger) {
	a := Auth{
		logger: logger,
	}

	eng.GET("/login/*any", a.HandleLogin)

	logger.Info("Auth handlers was inited")
}


func (a *Auth) HandleLogin(c *gin.Context) {
	usertkn := c.Param("any")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"path":    usertkn,
	})
}