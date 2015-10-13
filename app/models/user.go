package models

import (
	"errors"

	g "github.com/Gr1N/revel-gorm/app"
)

var (
	ErrUserNotExist = errors.New("User does not exist")
)

type User struct {
	Model

	Services []Service
}

func CreateUserByService(serviceName string, userServiceId int64,
	userServiceName, userServiceEmail string) *User {

	user := User{
		Services: []Service{{
			Name:             serviceName,
			UserServiceId:    userServiceId,
			UserServiceName:  userServiceName,
			UserServiceEmail: userServiceEmail,
		}},
	}
	g.DB.Create(&user)

	return &user
}

func GetUserById(id int64) (*User, error) {
	var user User
	g.DB.First(&user, id)

	if user.Id == 0 {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserByService(serviceName string, userServiceId int64) (*User, error) {
	var service Service
	g.DB.Where(&Service{
		Name:          serviceName,
		UserServiceId: userServiceId,
	}).First(&service)

	if service.Id == 0 {
		return nil, ErrUserNotExist
	}

	var user User
	g.DB.Model(&service).Related(&user)

	return &user, nil
}
