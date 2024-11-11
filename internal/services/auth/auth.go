package auth

import "github.com/gintokos/serverdb/pkg/logger"

var Auth auth

type auth struct {
	logger *logger.CustomLogger
}

func InitAuth(logger *logger.CustomLogger) {
	Auth = auth{
		logger: logger,
	}
}
