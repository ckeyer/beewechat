package conf

import (
	"github.com/astaxie/beego"
)

type CkConfig struct {
	WX_MYSQL_CONN string
	WECHAT_TOKEN  string
	WECHAT_APPID  string
	WECHAT_SECRET string
	REDIS_ADDR    string
}

func NewConfig() *CkConfig {
	return &CkConfig{
		WX_MYSQL_CONN: beego.AppConfig.String("wx_mysql_connstr"),
		WECHAT_TOKEN:  beego.AppConfig.String("Token"),
		WECHAT_APPID:  beego.AppConfig.String("appid"),
		WECHAT_SECRET: beego.AppConfig.String("app_secret"),
		REDIS_ADDR:    beego.AppConfig.String("redis_addr"),
	}
}
