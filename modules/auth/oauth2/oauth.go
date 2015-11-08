package oauth2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/franela/goreq"

	"github.com/Gr1N/pacman/modules/errors"
)

var (
	errServiceUnavailable = errors.New(
		"authorization_service_unavailable", "Authorization service temporary unavailable")
	errCodeInvalid = errors.New(
		"authorization_code_invalid", "Authorization code invalid or issued for another service")
	errTokenInvalid = errors.New(
		"access_token_invalid", "Access token invalid or expired")
)

// Config describes a typical 3-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string

	Endpoint Endpoint
}

// Endpoint contains the OAuth 2.0 provider's authorization, token and user
// endpoint URLs.
type Endpoint struct {
	AuthURL  string
	TokenURL string
	UserURL  string
}

// Token represents the crendentials used to authorize
// the requests to access protected resources on the OAuth 2.0
// provider's backend.
type Token struct {
	Access string
	Type   string
	Scopes []string
}

// User represents the user's data from authorization service.
type User struct {
	ID    int64
	Name  string
	Email string
}

type tokenJSON struct {
	Access string `json:"access_token"`
	Type   string `json:"token_type"`
	Scope  string `json:"scope"`
}

type userJSON struct {
	ID    int64  `json:"id"`
	Name  string `json:"login"`
	Email string `json:"email"`
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

// User returns the user's data from authorization service.
func (c *Config) User(token *Token) (*User, error) {
	authorization := strings.Join([]string{
		"token",
		token.Access,
	}, " ")
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    c.Endpoint.UserURL,
		Accept: "application/json",
	}.WithHeader("Authorization", authorization).Do()
	if err != nil {
		return nil, errServiceUnavailable
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return nil, errTokenInvalid
	}

	var uj userJSON

	if err := resp.Body.FromJsonTo(&uj); err != nil || uj.ID == 0 {
		return nil, errTokenInvalid
	}

	return &User{
		ID:    uj.ID,
		Name:  uj.Name,
		Email: uj.Email,
	}, nil
}
