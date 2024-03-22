package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movie       MovieModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionsModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Movie:       MovieModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
		Permissions: PermissionsModel{DB: db},
	}
}
