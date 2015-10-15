package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod(Any.attachUser, revel.BEFORE)
	revel.InterceptMethod(NotAuthenticated.attachUser, revel.BEFORE)
	revel.InterceptMethod(NotAuthenticated.checkUser, revel.BEFORE)
	revel.InterceptMethod(AnyAuthenticated.attachUser, revel.BEFORE)
	revel.InterceptMethod(AnyAuthenticated.checkUser, revel.BEFORE)
	revel.InterceptMethod(SessionAuthenticated.attachUser, revel.BEFORE)
	revel.InterceptMethod(SessionAuthenticated.checkUser, revel.BEFORE)
	revel.InterceptMethod(TokenAuthenticated.attachUser, revel.BEFORE)
	revel.InterceptMethod(TokenAuthenticated.checkUser, revel.BEFORE)
}
