package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordFound = errors.New("record not found")
)

type Models struct {
	Movie MovieModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Movie: MovieModel{DB: db},
	}
}
