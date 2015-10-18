package controllers

import (
	"net/http"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/jsonapi"
)

func (c Base) RenderJsonOk(items []*jsonapi.Item) revel.Result {
	c.Response.Status = http.StatusOK

	return c.RenderJson(jsonapi.ResultCollection{
		Data: items,
	})
}

func (c Base) RenderJsonCreated(item *jsonapi.Item) revel.Result {
	c.Response.Status = http.StatusCreated
	c.Response.Out.Header().Set("Location", item.Links.Self)

	return c.RenderJson(jsonapi.ResultIndividual{
		Data: item,
	})
}

func (c Base) RenderJsonBadRequest(errors []*jsonapi.Error) revel.Result {
	c.Response.Status = http.StatusBadRequest

	return c.RenderJson(jsonapi.ResultError{
		Errors: errors,
	})
}

func (c Base) RenderNoContent() revel.Result {
	c.Response.Status = http.StatusNoContent

	return c.RenderText("")
}

func (c Base) RenderNotFound() revel.Result {
	c.Response.Status = http.StatusNotFound

	return c.RenderText("")
}
