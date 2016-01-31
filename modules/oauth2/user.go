package oauth2

import (
	"net/http"
	"strings"

	"github.com/franela/goreq"
)

// User represents the user's data from authorization service.
type User struct {
	ID    int64
	Name  string
	Email string
}

type userJSON struct {
	ID    int64  `json:"id"`
	Name  string `json:"login"`
	Email string `json:"email"`
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
