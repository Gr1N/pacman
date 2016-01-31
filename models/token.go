package models

import (
	"github.com/Gr1N/pacman/modules/errors"
)

var (
	errUserTokenNotExist = errors.New(
		"service_not_exist", "User token does not exist")
)

type Token struct {
	Model

	UserID int64 `sql:"not null;unique_index:idx_userid_value"`

	Audience string `sql:"not null"`
	Value    string `sql:"not null;unique_index:idx_userid_value"`
}
