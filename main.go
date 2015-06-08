package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/ckeyer/beewechat/models"
	"github.com/ckeyer/beewechat/wechat"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/lab204?charset=utf8")
	wechat.RegDB()
	orm.RunSyncdb("default", true, true)
}

func main() {
	beego.Run()
}
