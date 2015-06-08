package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beewechat/wechat/event"
	"github.com/ckeyer/beewechat/wechat/msg"
)

func RegDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/lab204?charset=utf8")
	msg.RegDB()
	event.RegDB()
}
