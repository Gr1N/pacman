package oauth2

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/franela/goreq"
)

var (
	ErrCodeInvalid        = errors.New("Authorization code invalid")
	ErrAccessTokenInvalid = errors.New("AccessToken invalid")
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string

	Endpoint Endpoint
}

type Endpoint struct {
	AuthURL  string
	TokenURL string
	UserURL  string
}

type Token struct {
	Access string
	Type   string
	Scopes []string
}

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

func (c Config) AuthCodeURL(state string) string {
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

func (c Config) Exchange(code string) (*Token, error) {
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
		return nil, err
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return nil, ErrCodeInvalid
	}

	var tj tokenJSON

	if err := resp.Body.FromJsonTo(&tj); err != nil || tj.Access == "" {
		return nil, ErrCodeInvalid
	}

	return &Token{
		Access: tj.Access,
		Type:   tj.Type,
		Scopes: strings.Split(tj.Scope, ","),
	}, nil
}

func (c Config) User(token *Token) (*User, error) {
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
		return nil, err
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return nil, ErrAccessTokenInvalid
	}

	var uj userJSON

	if err := resp.Body.FromJsonTo(&uj); err != nil || uj.ID == 0 {
		return nil, ErrAccessTokenInvalid
	}

	return &User{
		ID:    uj.ID,
		Name:  uj.Name,
		Email: uj.Email,
	}, nil
}
