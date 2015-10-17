package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
	autht "github.com/Gr1N/pacman/app/modules/auth/token"
	"github.com/Gr1N/pacman/app/modules/jsonapi"
)

type Token struct {
	SessionAuthenticated
}

type tokenWithValueItemAttrs struct {
	Audience string `json:"audience"`
	Value    string `json:"value"`
	Created  int64  `json:"created"`
}

type tokenItemAttrs struct {
	Audience string `json:"audience"`
	Created  int64  `json:"created"`
}

func (c Token) Create() revel.Result {
	var audience string

	c.Params.Bind(&audience, "audience")

	if err := autht.ValidateTokenRequest(audience, c.Validation); err != nil {
		return c.RenderJsonBadRequest([]*jsonapi.Error{{
			Detail: err.Error(),
		}})
	}

	user := c.getUser()
	token := autht.FinishTokenRequest(user.Id, audience)

	return c.RenderJsonCreated(c.item(token, true))
}

func (c Token) ReadAll() revel.Result {
	user := c.getUser()
	tokens := models.GetUserTokens(user.Id)

	items := make([]*jsonapi.Item, len(tokens))
	for i := range items {
		items[i] = c.item(tokens[i], false)
	}

	return c.RenderJsonOk(items)
}

func (c Token) Read(id int64) revel.Result {
	user := c.getUser()

	if token, err := models.GetUserToken(user.Id, id); err == nil {
		item := c.item(token, false)
		return c.RenderJsonOk([]*jsonapi.Item{item})
	}

	return c.RenderNotFound()
}

func (c Token) Delete(id int64) revel.Result {
	user := c.getUser()
	models.DeleteUserToken(user.Id, id)

	return c.RenderNoContent()
}

func (c Token) item(token *models.Token, withValue bool) *jsonapi.Item {
	var attributes interface{}

	if withValue {
		attributes = tokenWithValueItemAttrs{
			Audience: token.Audience,
			Value:    token.Value,
			Created:  token.CreatedAt.Unix(),
		}
	} else {
		attributes = tokenItemAttrs{
			Audience: token.Audience,
			Created:  token.CreatedAt.Unix(),
		}
	}

	return &jsonapi.Item{
		Type:       "tokens",
		Id:         token.Id,
		Attributes: attributes,
		Links: jsonapi.ItemLinks{
			Self: token.Url(),
		},
	}
}
