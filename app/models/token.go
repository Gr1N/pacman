package models

import (
	"github.com/Gr1N/pacman/app/routes"
)

type Token struct {
	Model

	UserId int64 `sql:"not null;unique_index:idx_userid_value"`

	Audience string `sql:"not null"`
	Value    string `sql:"not null;unique_index:idx_userid_value"`
}

func (t Token) Url() string {
	return routes.Token.Read(t.Id)
}
