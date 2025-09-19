package repository

import (
	"database/sql"
	"errors"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db    *sqlx.DB
	cache *Cache
}

func NewRepository(db *sqlx.DB, cache *Cache) *Repository {
	return &Repository{
		db:    db,
		cache: cache,
	}
}

func (r *Repository) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotFound
	default:
		return err
	}
}
