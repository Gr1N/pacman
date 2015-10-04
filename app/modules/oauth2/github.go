package oauth2

import (
	"strings"

	"github.com/revel/revel"
)

var (
	GitHub *Config
)

func init() {
	revel.OnAppStart(initGitHub)
}

func initGitHub() {
	clientID, _ := revel.Config.String("auth.github.client_id")
	clientSecret, _ := revel.Config.String("auth.github.client_secret")
	redirectUrl, _ := revel.Config.String("auth.github.redirect_url")
	scopes, _ := revel.Config.String("auth.github.scopes")

	GitHub = &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectUrl:  redirectUrl,
		Scopes:       strings.Split(scopes, ","),

		Endpoint: Endpoint{
			AuthUrl:  "https://github.com/login/oauth/authorize",
			TokenUrl: "https://github.com/login/oauth/access_token",
		},
	}
}