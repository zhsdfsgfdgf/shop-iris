## Go商城秒杀项目
特性：
一致性哈希算法实现分布式验证；RabbitMQ进行下单异步处理，流量削峰；Go接口实现数量控制；AES双向加密Cookie；页面静态化，CDN部署；前端限流；wrk压测等。
## 项目流程
  ![image](https://github.com/zhsdfsgfdgf/istio-note/blob/master/images/miaosha.png)

## 本地开发
```go
# 获取代码
git clone https://github.com/zhsdfsgfdgf/shop-iris.git

# 进入工作路径
cd ./shop-iris

# 编译项目
# 注意: 要启用go module支持
go build

# 修改配置 
Mysql配置	/shop-iris/common/mysql.go
RabbitMQ配置	/shop-iris/rabbitmq/rabbitmq.go

# 交叉编译
GOOS=linux GOARCH=amd64 go build validate.go
```
