package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/gintokos/serverdb/pkg/logger"
	"github.com/gintokos/vpagrSite/internal/data/db"
	"github.com/gintokos/vpagrSite/pkg/telegramauth"
)

var Auth auth

type auth struct {
	logger *logger.CustomLogger
}

func InitAuth(logger *logger.CustomLogger) {
	Auth = auth{
		logger: logger,
	}
}

func (a *auth) Login(ctx context.Context, usertkn string) (int64, error) {
	bot := telegramauth.Tbot

	ok, id := bot.IsUsertokenExists(usertkn)
	if !ok {
		return 0, fmt.Errorf("not found this login link")
	}

	err := db.Database.CreateUser(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrUserAlreadyExists) {
			a.logger.Info("User with this id already exists")
			return id, nil
		}
		return 0, err
	}

	return id, nil
}
