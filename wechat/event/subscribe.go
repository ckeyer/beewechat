package event

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type SubscribeEvent struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Event        string    `xml:"Event"`
	EventKey     int32     `xml:"EventKey"`
	Ticket       string    `xml:"Ticket"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveUnsubscribeEvent(content string) string {
	var msg ScribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *SubscribeEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
