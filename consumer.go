package main

import (
	"fmt"
	"shop-iris/common"
	"shop-iris/rabbitmq"
	"shop-iris/repositories"
	"shop-iris/services"
)

func main() {
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
	}
	//创建product数据库操作实例
	product := repositories.NewProductManager("product", db)
	//创建product serivce
	productService := services.NewProductService(product)
	//创建Order数据库实例
	order := repositories.NewOrderMangerRepository("orders", db)
	//创建order Service
	orderService := services.NewOrderService(order)

	rabbitmqConsumeSimple := rabbitmq.NewRabbitMQSimple("shopProduct")
	rabbitmqConsumeSimple.ConsumeSimple(orderService, productService)
}
