package models

import (
	"github.com/Gr1N/pacman/modules/errors"
)

var (
	errServiceNotExist = errors.New(
		"service_not_exist", "User service does not exist")
)

// Service represents the object of external service.
type Service struct {
	Model

	UserID int64 `sql:"not null;unique_index:idx_userid_userserviceid"`

	Name string `sql:"not null;index"`

	AccessToken string `sql:"not null"`

	UserServiceID    int64 `sql:"not null;unique_index:idx_userid_userserviceid"`
	UserServiceName  string
	UserServiceEmail string

	Repos []Repo
}
