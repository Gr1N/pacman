package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrUserNotExist = errors.New("User does not exist")
)

type User struct {
	Model

	Services []Service
}

func CreateUserByService(db *gorm.DB, serviceName string, userServiceId int64,
	userServiceName, userServiceEmail string) *User {

	user := User{
		Services: []Service{{
			Name:             serviceName,
			UserServiceId:    userServiceId,
			UserServiceName:  userServiceName,
			UserServiceEmail: userServiceEmail,
		}},
	}
	db.Create(&user)

	return &user
}

func GetUserById(db *gorm.DB, id int64) (*User, error) {
	var user User
	db.First(&user, id)

	if user.Id == 0 {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserByService(db *gorm.DB, serviceName string, userServiceId int64) (*User, error) {
	var service Service
	db.Where(&Service{
		Name:          serviceName,
		UserServiceId: userServiceId,
	}).First(&service)

	if service.Id == 0 {
		return nil, ErrUserNotExist
	}

	var user User
	db.Model(&service).Related(&user)

	return &user, nil
}
