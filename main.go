package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beeweichat/models"
	_ "lab204/routers"
)

func init() {
	models.RegDB()
}

func main() {
	//start ORM debug
	orm.Debug = true
	//create table
	orm.RunSyncdb("default", true, true)
	beego.Run()
}
