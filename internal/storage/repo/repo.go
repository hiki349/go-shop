package repo

import (
	"go-shop/internal/storage/db"
	"time"
)

type Repo struct {
	db *db.DB
}

const maxTimeToDoDbOperation = time.Millisecond * 200

func New(db *db.DB) *Repo {
	return &Repo{
		db: db,
	}
}
