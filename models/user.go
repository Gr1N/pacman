package models

import (
	"github.com/pborman/uuid"

	"github.com/Gr1N/pacman/modules/errors"
	"github.com/Gr1N/pacman/modules/helpers"
)

var (
	errUserNotExist = errors.New(
		"user_not_exist", "User does not exist")
)

// User represents the object of individual user.
type User struct {
	Model

	Tokens   []Token
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

// CreateUserToken creates record of a new authentication token for user
// with specified `audience`.
func CreateUserToken(id int64, audience string) (*Token, error) {
	value := uuid.NewRandom().String()
	value = helpers.EncodeSha1(value)

	token := Token{
		UserID:   id,
		Audience: audience,
		Value:    value,
	}
	if err := g.Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
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

// GetUserByToken returns the user object by given token value if exists.
func GetUserByToken(value string) (*User, error) {
	var token Token
	if g.Where(&Token{
		Value: value,
	}).First(&token).RecordNotFound() {
		return nil, errUserTokenNotExist
	}

	var user User
	if g.Model(&token).Related(&user).RecordNotFound() {
		return nil, errUserNotExist
	}

	return &user, nil
}

// GetUserService returns the user service object by given user id and service
// if exists.
func GetUserService(id int64, serviceName string) (*Service, error) {
	var service Service
	if g.Where(&Service{
		UserID: id,
		Name:   serviceName,
	}).First(&service).RecordNotFound() {
		return nil, errServiceNotExist
	}

	return &service, nil
}

// GetUserReposByService returns list of user repos by given user id and service
// if exists.
func GetUserReposByService(id int64, serviceName string) ([]*Repo, error) {
	service, err := GetUserService(id, serviceName)
	if err != nil {
		return nil, err
	}

	var repos []*Repo
	if err := g.Where(&Repo{
		ServiceID: service.ID,
	}).Find(&repos).Error; err != nil {
		return nil, err
	}

	return repos, nil
}
