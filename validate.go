package main

import (
	"errors"
	"fmt"
	"net/http"
	"shop-iris/common"
)

//执行正常业务逻辑
func Check(w http.ResponseWriter, r *http.Request) {
	//执行正常业务逻辑
	fmt.Println("验证成功,执行check")
}

//统一验证拦截器，每个接口都需要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("拦截器进行验证")
	return errors.New("验证失败")
}

func main() {

	//1、过滤器
	filter := common.NewFilter()
	//注册拦截器,每次访问check接口,都会使用拦截器
	filter.RegisterFilterUri("/check", Auth)
	//2、启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	//启动服务
	http.ListenAndServe(":8083", nil)
}
