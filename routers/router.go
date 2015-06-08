package routers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})

	// beego.Router("/test", &controllers.TestController{})
	// beego.Router("/b", &controllers.BController{})
}
