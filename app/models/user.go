package models

import (
	"errors"

	g "github.com/Gr1N/revel-gorm/app"
)

var (
	ErrUserNotExist      = errors.New("User does not exist")
	ErrUserTokenNotExist = errors.New("User token does not exist")
)

type User struct {
	Model

	Services []Service
}

func CreateUserByService(serviceName string, userServiceId int64,
	userServiceName, userServiceEmail string) (*User, error) {

	user := User{
		Services: []Service{{
			Name:             serviceName,
			UserServiceId:    userServiceId,
			UserServiceName:  userServiceName,
			UserServiceEmail: userServiceEmail,
		}},
	}
	if err := g.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserToken(id int64, audience, value string) (*Token, error) {
	token := Token{
		UserId:   id,
		Audience: audience,
		Value:    value,
	}
	if err := g.DB.Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	if g.DB.First(&user, id).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserByService(serviceName string, userServiceId int64) (*User, error) {
	var service Service
	if g.DB.Where(&Service{
		Name:          serviceName,
		UserServiceId: userServiceId,
	}).First(&service).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	var user User
	if g.DB.Model(&service).Related(&user).RecordNotFound() {
		return nil, ErrUserTokenNotExist
	}

	return &user, nil
}

func GetUserTokens(id int64) ([]*Token, error) {
	var tokens []*Token
	if err := g.DB.Find(&tokens, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	return tokens, nil
}

func GetUserToken(id, tokenId int64) (*Token, error) {
	var token Token
	if g.DB.First(&token, "id = ? AND user_id = ?", tokenId, id).RecordNotFound() {
		return nil, ErrUserTokenNotExist
	}

	return &token, nil
}

func DeleteUserToken(id, tokenId int64) error {
	return g.DB.Where("id = ? AND user_id = ?", tokenId, id).
		Delete(Token{}).Error
}
