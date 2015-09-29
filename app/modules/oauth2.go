package modules

import (
	"bytes"
	"net/url"
	"strings"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string

	Endpoint Endpoint
}

type Endpoint struct {
	AuthUrl  string
	TokenUrl string
}

var (
	GitHub = &Config{
		ClientID:     "fake",
		ClientSecret: "fake",
		RedirectUrl:  "http://localhost",
		Scopes:       []string{},

		Endpoint: Endpoint{
			AuthUrl:  "https://github.com/login/oauth/authorize",
			TokenUrl: "https://github.com/login/oauth/access_token",
		},
	}
)

func (c *Config) AuthCodeUrl(state string) string {
	var buf bytes.Buffer

	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
		"redirect_uri":  {c.RedirectUrl},
		"scope":         {strings.Join(c.Scopes, " ")},
		"state":         {state},
	}

	buf.WriteString(c.Endpoint.AuthUrl)
	buf.WriteString("?")
	buf.WriteString(v.Encode())

	return buf.String()
}
