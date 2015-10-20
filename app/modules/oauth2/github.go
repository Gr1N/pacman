package oauth2

import (
	"strings"

	"github.com/revel/revel"
)

var (
	GitHub *Config
)

func initGitHub() {
	clientID, _ := revel.Config.String("auth.github.client_id")
	clientSecret, _ := revel.Config.String("auth.github.client_secret")
	redirectURL, _ := revel.Config.String("auth.github.redirect_url")
	scopes, _ := revel.Config.String("auth.github.scopes")

	GitHub = &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       strings.Split(scopes, ","),

		Endpoint: Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
			UserURL:  "https://api.github.com/user",
		},
	}
}
