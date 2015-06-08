package event

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type ViewEvent struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Event        string    `xml:"Event"`
	EventKey     string    `xml:"EventKey"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveViewEvent(content string) string {
	var msg ViewEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *ViewEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
