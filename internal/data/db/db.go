package db

import (
	"context"
	"github.com/gintokos/vpagrSite/internal/domain/models"
)

var Database DB

type DB interface {
	GetUser(ctx context.Context, id int64) (user models.User, err error)
	CreateUser(ctx context.Context, id int64) error
}

func SetDataBase(db DB) {
	Database = db
}
