package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type LinkMsg struct {
	Id int64,
	ToUserName   string    `xml: "ToUserName"`
	FromUserName string    `xml: "FromUserName"`
	CreateTime   int       `xml: "CreateTime"`
	MsgType      string    `xml: "MsgType"`
	Title        string    `xml: "Title"`
	Description  string    `xml: "Description"`
	Url          string    `xml: "Url"`
	MsgId        int64     `xml: "MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveLinkMsg(content string) string {
	var msg LinkMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *LinkMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}
