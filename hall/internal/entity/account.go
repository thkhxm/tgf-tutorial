package entity

import "github.com/thkhxm/tgf/db"

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/13
//***************************************************

type UserAccount struct {
	db.Model
	Account  string `orm:"pk"`
	Password string
	UserId   string
}

func (u *UserAccount) GetTableName() string {
	return "t_account"
}
