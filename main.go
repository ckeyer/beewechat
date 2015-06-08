package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/ckeyer/beewechat/models"
	config "github.com/ckeyer/beewechat/conf"
	_ "github.com/ckeyer/beewechat/routers"
	"github.com/ckeyer/beewechat/wechat"
)

func init() {
	force := false   // 强制重新建表
	verbose := false // 打印SQL语句
	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", config.WX_MYSQL_CONN)
	wechat.RegDB()
	orm.RunSyncdb("default", force, verbose)
}

func main() {
	beego.Run()
}
