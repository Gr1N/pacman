package tests

import (
	"github.com/revel/revel/testing"

	"github.com/Gr1N/pacman/app/routes"
)

type ControllersAppTestSuite struct {
	testing.TestSuite
}

func (t *ControllersAppTestSuite) TestIndex() {
	t.Get(routes.Application.Index())
	t.AssertOk()
	t.AssertContentType(contentTypeTextHTML)
}
