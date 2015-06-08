package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beewechat/conf"
	"github.com/ckeyer/beewechat/wechat/event"
	"github.com/ckeyer/beewechat/wechat/msg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
)

var (
	redcli redis.Client
	config *conf.CkConfig
)

func init() {
	config = conf.NewConfig()
	redcli.Addr = config.REDIS_ADDR
}

func RegDB() {
	msg.RegDB()
	event.RegDB()
	orm.RegisterModel(new(WebUserInfo))
}
