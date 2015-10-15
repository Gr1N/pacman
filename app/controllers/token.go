package controllers

import (
	"net/http"

	"github.com/revel/revel"

	autht "github.com/Gr1N/pacman/app/modules/auth/token"
)

type Token struct {
	SessionAuthenticated
}

func (c Token) Create() revel.Result {
	var audience string

	c.Params.Bind(&audience, "audience")

	if err := autht.ValidateTokenRequest(audience, c.Validation); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJson(map[string]interface{}{
			"error": err.Error(),
		})
	}

	user := c.getUser()
	token := autht.FinishTokenRequest(user.Id, audience)

	return c.RenderJson(map[string]interface{}{
		"audience":   token.Audience,
		"token":      token.Value,
		"created_at": token.CreatedAt.Unix(),
	})
}
