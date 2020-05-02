package controllers

import (
	"shop-iris/datamodels"
	"shop-iris/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type UserController struct {
	//要用到上下文的地方很多,所以放在这里
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

//注册页面
func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

/*
	映射表单
		一.创建结构体,映射
			1.创建product实例
			2.通过form映射
		二.适用于字段比较少
			单独拿出来处理
*/
func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	user := &datamodels.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}
	_, err := c.Service.AddUser(user)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		//记得return
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}
