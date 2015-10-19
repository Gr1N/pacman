package tests

import (
	"net/http"
	"net/url"

	"github.com/revel/revel"
	"github.com/revel/revel/testing"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/routes"
)

type ControllersTokenTestSuite struct {
	testing.TestSuite

	user *models.User
}

func (t *ControllersTokenTestSuite) Before() {
}

func (t *ControllersTokenTestSuite) After() {
	detachUser(t)
}

func (t *ControllersTokenTestSuite) getSession() revel.Session {
	return t.Session
}

func (t *ControllersTokenTestSuite) getUser() *models.User {
	return t.user
}

func (t *ControllersTokenTestSuite) setUser(user *models.User) {
	t.user = user
}

func (t *ControllersTokenTestSuite) TestCreateUnauthorized() {
	t.PostForm(routes.Token.Create(), url.Values{})
	t.AssertStatus(http.StatusOK)
	t.AssertContentType(contentTypeTextHtml)
}

func (t *ControllersTokenTestSuite) TestCreateBadRequest() {
	attachUser(t)

	t.PostForm(routes.Token.Create(), url.Values{})
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType(contentTypeApplicationJson)

	// r := getResultError(t.ResponseBody)
	// revel.INFO.Println(r.Errors[0].Detail)
}

func (t *ControllersTokenTestSuite) TestCreateCreated() {
}

func (t *ControllersTokenTestSuite) TestReadAllUnauthorized() {
}

func (t *ControllersTokenTestSuite) TestReadAllOk() {
}

func (t *ControllersTokenTestSuite) TestReadUnauthorized() {
}

func (t *ControllersTokenTestSuite) TestReadNotFound() {
}

func (t *ControllersTokenTestSuite) TestReadOk() {
}

func (t *ControllersTokenTestSuite) TestDeleteNoContent() {
}
