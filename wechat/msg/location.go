package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type LocationMsg struct {
	ToUserName   string    `xml: "ToUserName"`
	FromUserName string    `xml: "FromUserName"`
	CreateTime   int       `xml: "CreateTime"`
	MsgType      string    `xml: "MsgType"`
	Location_X   float64   `xml: "Location_X"`
	Location_Y   float64   `xml: "Location_Y"`
	Scale        int       `xml: "Scale"`
	Label        string    `xml: "Label"`
	MsgId        int64     `xml: "MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveLocationMsg(content string) string {
	var msg LocationMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *LocationMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}
