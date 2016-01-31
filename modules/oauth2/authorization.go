package oauth2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/franela/goreq"
)

// Token represents the crendentials used to authorize
// the requests to access protected resources on the OAuth 2.0
// provider's backend.
type Token struct {
	Access string
	Type   string
	Scopes []string
}

type tokenJSON struct {
	Access string `json:"access_token"`
	Type   string `json:"token_type"`
	Scope  string `json:"scope"`
}

// AuthCodeURL constructs authorization url according to the config.
func (c *Config) AuthCodeURL(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
		"redirect_uri":  {c.RedirectURL},
		"scope":         {strings.Join(c.Scopes, " ")},
		"state":         {state},
	}

	return strings.Join([]string{
		c.Endpoint.AuthURL,
		v.Encode(),
	}, "?")
}

// Exchange exchanges authorization code to the access token.
func (c *Config) Exchange(code string) (*Token, error) {
	body := map[string]string{
		"grant_type":   "authorization_code",
		"code":         code,
		"redirect_uri": c.RedirectURL,
	}
	resp, err := goreq.Request{
		Method:            "POST",
		Uri:               c.Endpoint.TokenURL,
		Body:              body,
		ContentType:       "application/json",
		Accept:            "application/json",
		BasicAuthUsername: c.ClientID,
		BasicAuthPassword: c.ClientSecret,
	}.Do()
	if err != nil {
		return nil, errServiceUnavailable
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return nil, errCodeInvalid
	}

	var tj tokenJSON

	if err := resp.Body.FromJsonTo(&tj); err != nil || tj.Access == "" {
		return nil, errCodeInvalid
	}

	return &Token{
		Access: tj.Access,
		Type:   tj.Type,
		Scopes: strings.Split(tj.Scope, ","),
	}, nil
}
