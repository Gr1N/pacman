package controllers

import (
	"strconv"

	"github.com/revel/revel"

	gorm "github.com/Gr1N/revel-gorm/app/controllers"

	"github.com/Gr1N/pacman/app/models"
)

type Application struct {
	gorm.TransactionalController
}

func (c Application) tryAuthenticate() revel.Result {
	if user := c.withUser(); user != nil {
		c.RenderArgs["user"] = user
	}

	return nil
}

func (c Application) withUser() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}

	if userId, ok := c.Session["user_id"]; ok {
		if id, err := strconv.ParseInt(userId, 10, 64); err == nil {
			return c.getUser(id)
		}
	}

	return nil
}

func (c Application) getUser(userId int64) *models.User {
	var user models.User

	c.Txn.First(&user, userId)
	if user.Id != 0 {
		return &user
	}

	return nil
}

func (c Application) Index() revel.Result {
	return c.Render()
}
