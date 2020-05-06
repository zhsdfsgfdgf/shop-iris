package controllers

import (
	"fmt"
	"shop-iris/datamodels"
	"shop-iris/encrypt"
	"shop-iris/services"
	"shop-iris/tool"
	"strconv"

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
func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "/user/login.html",
	}
}

//除了可以通过redirect跳转,也可以用mvc.Response
func (c *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//2、验证账号密码正确
	user, isOk := c.Service.IsPwdSuccess(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	//3、写入用户ID到cookie中
	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	uidByte := []byte(strconv.FormatInt(user.ID, 10))
	uidString, err := encrypt.EnPwdCode(uidByte)
	if err != nil {
		fmt.Println(err)
	}
	// c.Session.Set("userID", strconv.FormatInt(user.ID, 10))
	//写入到用户浏览器
	tool.GlobalCookie(c.Ctx, "sign", uidString)
	return mvc.Response{
		Path: "/product/",
	}

}
