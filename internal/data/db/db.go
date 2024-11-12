package db

import (
	"context"
	"errors"

	"github.com/gintokos/vpagrSite/internal/domain/models"
)

var Database DB

var ErrUserAlreadyExists = errors.New("record with this id was created before")

type DB interface {
	GetUser(ctx context.Context, id int64) (user models.User, err error)
	CreateUser(ctx context.Context, id int64) error
}

func SetDataBase(db DB) {
	Database = db
}
