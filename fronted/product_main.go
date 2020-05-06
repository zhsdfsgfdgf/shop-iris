package main

import (
	"github.com/kataras/iris"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置模板
	//访问方式http://127.0.0.1:8082/public/js/bootstrap.min.js
	app.StaticWeb("/public", "./web/public")
	//3.访问生成好的html静态文件
	//http://127.0.0.1:8082/html/htmlProduct.html
	app.StaticWeb("/html", "./web/htmlProductShow")
	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
