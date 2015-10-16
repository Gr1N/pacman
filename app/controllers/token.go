package controllers

import (
	"github.com/revel/revel"

	autht "github.com/Gr1N/pacman/app/modules/auth/token"
	"github.com/Gr1N/pacman/app/modules/jsonapi"
)

type Token struct {
	SessionAuthenticated
}

type tokenItemAttrs struct {
	Audience string `json:"audience"`
	Value    string `json:"value"`
	Created  int64  `json:"created"`
}

func (c Token) Create() revel.Result {
	var audience string

	c.Params.Bind(&audience, "audience")

	if err := autht.ValidateTokenRequest(audience, c.Validation); err != nil {
		return c.RenderJsonBadRequest([]jsonapi.Error{{
			Detail: err.Error(),
		}})
	}

	user := c.getUser()
	token := autht.FinishTokenRequest(user.Id, audience)

	location := "TBD"

	return c.RenderJsonCreated(jsonapi.Item{
		Type: "tokens",
		Id:   token.Id,
		Attributes: tokenItemAttrs{
			Audience: token.Audience,
			Value:    token.Value,
			Created:  token.CreatedAt.Unix(),
		},
		Links: jsonapi.ItemLinks{
			Self: location,
		},
	})
}
