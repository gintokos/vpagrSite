package services

import (
	"github.com/gintokos/serverdb/pkg/logger"
	"github.com/gintokos/vpagrSite/internal/services/auth"
)

func InitServices(logger *logger.CustomLogger) error {
	auth.InitAuth(logger)
	
	return nil
}