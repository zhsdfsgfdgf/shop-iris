package common

import (
	"net/http"
)

//声明一个新的数据类型（函数类型）拦截器函数
type FilterHandle func(rw http.ResponseWriter, req *http.Request) error

//拦截器结构体
type Filter struct {
	//用来存储需要拦截的URI
	filterMap map[string]FilterHandle
}

//Filter初始化函数
func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]FilterHandle)}
}

//注册拦截器
func (f *Filter) RegisterFilterUri(uri string, handler FilterHandle) {
	//将uri绑定到handle上进行执行
	f.filterMap[uri] = handler
}

//根据Uri获取对应的handle
func (f *Filter) GetFilterHandle(uri string) FilterHandle {
	return f.filterMap[uri]
}

//声明新的函数类型,用来处理请求
type WebHandle func(rw http.ResponseWriter, req *http.Request)

//执行拦截器，返回函数类型
func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		//在正常执行业务逻辑之前,会对注册的uri进行判断,看是否存在拦截
		for path, handle := range f.filterMap {
			if path == r.RequestURI {
				//执行拦截业务逻辑
				err := handle(rw, r)
				if err != nil {
					rw.Write([]byte(err.Error()))
					return
				}
				//跳出循环
				break
			}
		}
		//执行正常注册的函数
		webHandle(rw, r)
	}
}
