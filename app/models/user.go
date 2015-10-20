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

func CreateUser() (*User, error) {
	user := User{}
	if err := g.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

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
	if err := g.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserToken(id int64, audience, value string) (*Token, error) {
	token := Token{
		UserID:   id,
		Audience: audience,
		Value:    value,
	}
	if err := g.DB.Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func GetUserByID(id int64) (*User, error) {
	var user User
	if g.DB.First(&user, id).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserByService(serviceName string, userServiceID int64) (*User, error) {
	var service Service
	if g.DB.Where(&Service{
		Name:          serviceName,
		UserServiceID: userServiceID,
	}).First(&service).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	var user User
	if g.DB.Model(&service).Related(&user).RecordNotFound() {
		return nil, ErrUserNotExist
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

func GetUserToken(id, tokenID int64) (*Token, error) {
	var token Token
	if g.DB.First(&token, "id = ? AND user_id = ?", tokenID, id).RecordNotFound() {
		return nil, ErrUserTokenNotExist
	}

	return &token, nil
}

// ATTENTION: Use it only in tests
func DeleteUser(id int64) {
	g.DB.Where("user_id = ?", id).Delete(Token{})
	g.DB.Where("user_id = ?", id).Delete(Service{})
	g.DB.Where("id = ?", id).Delete(User{})
}

func DeleteUserToken(id, tokenID int64) error {
	return g.DB.Where("id = ? AND user_id = ?", tokenID, id).
		Delete(Token{}).Error
}
