package event

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hoisie/redis"
	"time"
)

// 二维码扫码事件结构体
type ScanEvent struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Event        string    `xml:"Event"`
	EventKey     string    `xml:"EventKey"`
	Ticket       string    `xml:"Ticket"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
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

func (this *ScanEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
