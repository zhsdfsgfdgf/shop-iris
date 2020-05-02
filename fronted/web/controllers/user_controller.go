package controllers

import (
	"shop-iris/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type UserController struct {
	//要用到上下文的地方很多,所以放在这里
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}
