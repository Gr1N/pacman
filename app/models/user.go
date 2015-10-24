package models

import (
	"errors"

	g "github.com/Gr1N/revel-gorm/app"
)

var (
	ErrUserNotExist      = errors.New("User does not exist")
	ErrServiceNotExist   = errors.New("User service does not exist")
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

func CreateUserRepo(serviceID int64, name, description string, private, fork bool,
	url, homepage string) (*Repo, error) {

	repo := Repo{
		ServiceID:   serviceID,
		Name:        name,
		Description: description,
		Private:     private,
		Fork:        fork,
		URL:         url,
		Homepage:    homepage,
	}
	if err := g.DB.Create(&repo).Error; err != nil {
		return nil, err
	}

	return &repo, nil
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
		return nil, ErrServiceNotExist
	}

	var user User
	if g.DB.Model(&service).Related(&user).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserByToken(value string) (*User, error) {
	var token Token
	if g.DB.Where(&Token{
		Value: value,
	}).First(&token).RecordNotFound() {
		return nil, ErrUserTokenNotExist
	}

	var user User
	if g.DB.Model(&token).Related(&user).RecordNotFound() {
		return nil, ErrUserNotExist
	}

	return &user, nil
}

func GetUserService(id int64, serviceName string) (*Service, error) {
	var service Service
	if g.DB.Where(&Service{
		UserID: id,
		Name:   serviceName,
	}).First(&service).RecordNotFound() {
		return nil, ErrServiceNotExist
	}

	return &service, nil
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

func GetUserReposByService(id int64, serviceName string) ([]*Repo, error) {
	var service Service
	if g.DB.Where(&Service{
		UserID: id,
		Name:   serviceName,
	}).First(&service).RecordNotFound() {
		return nil, ErrServiceNotExist
	}

	var repos []*Repo
	if err := g.DB.Find(&repos, "service_id = ?", service.ID).Error; err != nil {
		return nil, err
	}

	return repos, nil
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
