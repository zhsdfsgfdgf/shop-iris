package controllers

import (
	"shop-iris/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductByID(3)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}
