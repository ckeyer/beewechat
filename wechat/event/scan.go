package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

// 二维码扫码事件结构体
type ScanEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

func ReceiveScanEvent(content string) string {
	var this ScanEvent
	fmt.Println(content)
	err := xml.Unmarshal([]byte(content), &this)
	if err != nil {
		return ""
	}
	data := "请联系系统管理员进行身份认证"
	if "101" == this.EventKey &&
		"oecJ3jhN5usPBQMIXqc9bVP0toi4" == this.Ticket {
		data = "吼吼吼"
		var redcli redis.Client
		redcli.Addr = beego.AppConfig.String("redis_addr")
		redcli.Hset(this.Ticket, "scan", []byte("true"))
	}
	fmt.Println(data)
	rcontent := `<xml>
<ToUserName><![CDATA[` + this.FromUserName + `]]></ToUserName>
<FromUserName><![CDATA[` + this.ToUserName + `]]></FromUserName>
<CreateTime>` + fmt.Sprint((time.Now().Unix())) + `</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[` + data + `]]></Content>
</xml>`
	return rcontent
}
