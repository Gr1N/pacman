package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod(Base.tryAuthenticate, revel.BEFORE)
	revel.InterceptMethod(AuthSocial.checkAuthentication, revel.BEFORE)
}
