package models

import (
	"github.com/Gr1N/pacman/modules/errors"
)

var (
	errUserNotExist = errors.New(
		"user_not_exist", "User does not exist")
)

// User represents the object of individual user.
type User struct {
	Model

	Services []Service
}

// CreateUserByService creates record of a new user using given service.
func CreateUserByService(serviceName, serviceAccessToken string, userServiceID int64,
	userServiceName, userServiceEmail string) (*User, error) {

	user := User{
		Services: []Service{{
			Name:             serviceName,
			AccessToken:      serviceAccessToken,
			UserServiceID:    userServiceID,
			UserServiceName:  userServiceName,
			UserServiceEmail: userServiceEmail,
		}},
	}
	if err := g.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID returns the user object by given id if exists.
func GetUserByID(id int64) (*User, error) {
	var user User
	if g.First(&user, id).RecordNotFound() {
		return nil, errUserNotExist
	}

	return &user, nil
}

// GetUserByService returns the user object by given service if exists.
func GetUserByService(serviceName string, userServiceID int64) (*User, error) {
	var service Service
	if g.Where(&Service{
		Name:          serviceName,
		UserServiceID: userServiceID,
	}).First(&service).RecordNotFound() {
		return nil, errServiceNotExist
	}

	var user User
	if g.Model(&service).Related(&user).RecordNotFound() {
		return nil, errUserNotExist
	}

	return &user, nil
}
