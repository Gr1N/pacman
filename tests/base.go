package tests

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/modules/jsonapi"
)

const (
	contentTypeTextHtml        = "text/html; charset=utf-8"
	contentTypeApplicationJson = "application/json; charset=utf-8"
)

type TestSuiteWithUser interface {
	getSession() revel.Session
	getUser() *models.User
	setUser(user *models.User)
}

func attachUser(t TestSuiteWithUser) {
	if t.getUser() == nil {
		user, _ := models.CreateUser()

		t.getSession()["user_id"] = strconv.FormatInt(user.Id, 10)
		t.setUser(user)
	}
}

func detachUser(t TestSuiteWithUser) {
	if t.getUser() != nil {
		models.DeleteUser(t.getUser().Id)

		delete(t.getSession(), "user_id")
		t.setUser(nil)
	}
}

func getResultError(body []byte) *jsonapi.ResultError {
	var r jsonapi.ResultError
	if err := json.Unmarshal(body, &r); err != nil {
		panic(err)
	}

	return &r
}
