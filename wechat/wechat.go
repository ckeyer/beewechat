package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beewechat/wechat/event"
	"github.com/ckeyer/beewechat/wechat/msg"
	_ "github.com/go-sql-driver/mysql"
)

func RegDB() {
	msg.RegDB()
	event.RegDB()
	orm.RegisterModel(new(WebUserInfo))
}
