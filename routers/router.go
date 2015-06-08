package routers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/controllers"
)

func init() {
	beego.Router("/", &controllers.WeChatController{})
	beego.Router("/home", &controllers.WeChatController{})

	// beego.Router("/test", &controllers.TestController{})
	// beego.Router("/b", &controllers.BController{})
}
