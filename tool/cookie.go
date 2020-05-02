package tool

import (
	"net/http"

	"github.com/kataras/iris"
)

//设置全局cookie
func GlobalCookie(ctx iris.Context, name string, value string) {
	ctx.SetCookie(&http.Cookie{Name: name, Value: value, Path: "/"})
}
