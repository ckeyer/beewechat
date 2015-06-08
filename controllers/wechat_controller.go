package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	wx "github.com/ckeyer/beewechat/wechat"
	"io"
	"log"
	// "net/http"
	"sort"
)

type WeChatController struct {
	beego.Controller
}

func (this *WeChatController) Get() {
	//Start 微信服务器认证部分
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	tmps := []string{config.WECHAT_TOKEN, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]

	tmp := func(data string) string {
		t := sha1.New()
		io.WriteString(t, data)
		return fmt.Sprintf("%x", t.Sum(nil))
	}(tmpStr)
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

	// this.Data["QRImgUrl"] = wx.GetTempTicket(120, 1214, "FUNX_HOME")
	// this.Data["PageTitle"] = "FANX-HOME"
	// this.Data["Website"] = "beego.me"
	// this.Data["Email"] = "astaxie@gmail.com"
	// this.TplNames = ".tpl"
}

func (this *WeChatController) Post() {
	s := fmt.Sprintf("%s", this.Ctx.Input.CopyBody())
	// log.Println(s)
	r := wx.ReceiveMsg(s)
	// log.Println(r)
	this.Ctx.WriteString(r)
}
