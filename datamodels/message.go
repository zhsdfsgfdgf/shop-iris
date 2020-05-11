package datamodels

//简单的消息体
type Message struct {
	ProductID int64
	UserID    int64
}

//创建结构体
func NewMessage(userId int64, productId int64) *Message {
	return &Message{UserID: userId, ProductID: productId}
}
