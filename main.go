package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/ckeyer/beewechat/models"
	"github.com/ckeyer/beewechat/wechat"
)

func init() {
	wechat.RegDB()
}

func main() {
	//start ORM debug
	orm.Debug = true
	//create table
	orm.RunSyncdb("default", true, true)
	beego.Run()
}
