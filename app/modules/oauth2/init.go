package oauth2

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(func() {
		initGitHub()
	})
}
