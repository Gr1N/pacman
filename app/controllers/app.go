package controllers

import (
	"github.com/revel/revel"
)

type Application struct {
	Base
}

func (c Application) Index() revel.Result {
	return c.Render()
}
