package oauth2

import (
	"net/url"
	"strings"

	"github.com/Gr1N/pacman/app/modules/helpers"
)

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string

	Endpoint Endpoint
}

type Endpoint struct {
	AuthUrl  string
	TokenUrl string
}

func (c Config) AuthCodeUrl(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientId},
		"redirect_uri":  {c.RedirectUrl},
		"scope":         {strings.Join(c.Scopes, " ")},
		"state":         {state},
	}

	return helpers.JoinStrings(c.Endpoint.AuthUrl, "?", v.Encode())
}
