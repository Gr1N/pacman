package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/jsonapi"
)

func (c Base) RenderJsonCreated(item jsonapi.Item) revel.Result {
	c.Response.Status = http.StatusCreated
	c.Response.Out.Header().Set("Location", item.Links.Self)

	return c.RenderJson(map[string]jsonapi.Item{
		"data": item,
	})
}

func (c Base) RenderJsonBadRequest(errors []jsonapi.Error) revel.Result {
	c.Response.Status = http.StatusBadRequest

	return c.RenderJson(map[string][]jsonapi.Error{
		"errors": errors,
	})
}
