package datamodels

type Product struct {
	//sql标签用来给数据库用,shop标签用来映射表单
	ID           int64  `json:"id" sql:"ID" shop:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" shop:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" shop:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" shop:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" shop:"ProductUrl"`
}
