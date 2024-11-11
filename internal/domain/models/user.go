package models

import "time"

type User struct {
	Tid       int64
	CreatedAt time.Time
}
