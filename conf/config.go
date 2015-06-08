package conf

import (
	"github.com/astaxie/beego"
)

var (
	WX_MYSQL_CONN = beego.AppConfig.String("wx_mysql_connstr")
	WECHAT_TOKEN  = beego.AppConfig.String("Token")
	WECHAT_APPID  = beego.AppConfig.String("appid")
	WECHAT_SECRET = beego.AppConfig.String("app_secret")
)
