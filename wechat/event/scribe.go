package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type ScribeEvent struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Event        string    `xml:"Event"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveSubscribeEvent(content string) string {
	var msg SubscribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *ScribeEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
