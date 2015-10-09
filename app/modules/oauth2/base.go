package oauth2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/franela/goreq"

	"github.com/revel/revel"
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
	UserUrl  string
}

type Token struct {
	Access string
	Type   string
	Scopes []string
}

type User struct {
	Id    int64
	Name  string
	Email string
}

type tokenJSON struct {
	Access string `json:"access_token"`
	Type   string `json:"token_type"`
	Scope  string `json:"scope"`
}

type userJSON struct {
	Id    int64  `json:"id"`
	Name  string `json:"login"`
	Email string `json:"email"`
}

func (c Config) AuthCodeUrl(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientId},
		"redirect_uri":  {c.RedirectUrl},
		"scope":         {strings.Join(c.Scopes, " ")},
		"state":         {state},
	}

	return strings.Join([]string{
		c.Endpoint.AuthUrl,
		v.Encode(),
	}, "?")
}

func (c Config) Exchange(code string) *Token {
	body := map[string]string{
		"grant_type":   "authorization_code",
		"code":         code,
		"redirect_uri": c.RedirectUrl,
	}
	resp, err := goreq.Request{
		Method:            "POST",
		Uri:               c.Endpoint.TokenUrl,
		Body:              body,
		ContentType:       "application/json",
		Accept:            "application/json",
		BasicAuthUsername: c.ClientId,
		BasicAuthPassword: c.ClientSecret,
	}.Do()
	if err != nil {
		revel.ERROR.Printf("Got error (%v), while fetching token", err)
		return nil
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		revel.ERROR.Printf("Got unexpected status (%v) while fetching token",
			status)
		return nil
	}

	var tj tokenJSON

	if err := resp.Body.FromJsonTo(&tj); err != nil || tj.Access == "" {
		revel.ERROR.Printf("AccessToken not found in token response")
		return nil
	}

	return &Token{
		Access: tj.Access,
		Type:   tj.Type,
		Scopes: strings.Split(tj.Scope, ","),
	}
}

func (c Config) User(token *Token) *User {
	authorization := strings.Join([]string{
		"token",
		token.Access,
	}, " ")
	resp, err := goreq.Request{
		Method: "GET",
		Uri:    c.Endpoint.UserUrl,
		Accept: "application/json",
	}.WithHeader("Authorization", authorization).Do()
	if err != nil {
		revel.ERROR.Printf("Got error (%v), while fetching user data", err)
		return nil
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		revel.ERROR.Printf("Got unexpected status (%v) while user data", status)
		return nil
	}

	var uj userJSON

	if err := resp.Body.FromJsonTo(&uj); err != nil || uj.Id == 0 {
		revel.ERROR.Printf("User Id not found in response")
		return nil
	}

	return &User{
		Id:    uj.Id,
		Name:  uj.Name,
		Email: uj.Email,
	}
}
