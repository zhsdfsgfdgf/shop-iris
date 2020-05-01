package datamodels

type User struct {
	ID       int64  `json:"id" form:"ID" sql:"ID"`
	NickName string `json:"nickName" form:"nickName" sql:"nickName"`
	UserName string `json:"userName" form:"userName" sql:"userName"`
	//存后台加密的密码,-表示为空
	HashPassword string `json:"-" form:"passWord" sql:"passWord"`
}
