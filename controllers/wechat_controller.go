package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	config "github.com/ckeyer/beewechat/conf"
	wx "github.com/ckeyer/beewechat/wechat"
	"io"
	"log"
	// "net/http"
	"sort"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//Start 微信服务器认证部分
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	tmps := []string{config.WECHAT_TOKEN, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	if tmp == signature {
		this.Ctx.WriteString(echostr)
		return
	}
	//Over 微信服务器认证部分

	//Start 微信网页认证部分
	code := this.GetString("code")
	log.Println(this.GetString("code"))
	log.Println(this.GetString("status"))
	if code != "" {
		wat := wx.GetWebAccessToken(code)
		u := wat.GetUserInfo()
		u.Insert()
	}
	//Over 微信网页认证部分

	this.Data["QRImgUrl"] = wx.GetTempTicket(120, 1214, "FUNX_HOME")
	this.Data["PageTitle"] = "FANX-HOME"
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["StaticHost"] = beego.AppConfig.String("static_host")
	this.TplNames = "funx-data.tpl"
}

func (this *MainController) Post() {
	s := fmt.Sprintf("%s", this.Ctx.Input.CopyBody())
	// log.Println(s)
	r := wx.ReceiveMsg(s)
	// log.Println(r)
	this.Ctx.WriteString(r)
}

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
