package datamodels

type Order struct {
	ID          int64 `sql:"ID"`
	UserId      int64 `sql:"userID"`
	ProductId   int64 `sql:"productID"`
	OrderStatus int64 `sql:"orderStatus"`
}

//订单状态,等待付款,成功状态,失败
const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)
