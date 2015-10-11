package controllers

import (
	"github.com/revel/revel"

	gorm "github.com/Gr1N/revel-gorm/app/controllers"
)

func init() {
	revel.InterceptMethod((*gorm.TransactionalController).Begin, revel.BEFORE)
	revel.InterceptMethod(Application.tryAuthenticate, revel.BEFORE)
	revel.InterceptMethod(AuthSocial.checkAuthentication, revel.BEFORE)

	revel.InterceptMethod((*gorm.TransactionalController).Commit, revel.AFTER)

	revel.InterceptMethod((*gorm.TransactionalController).Rollback, revel.FINALLY)
}
