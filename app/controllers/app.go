package controllers

import (
	"github.com/revel/revel"
)

type Application struct {
	Any
}

func (c Application) Index() revel.Result {
	return c.Render()
}
