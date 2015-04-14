package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type ImageMsg struct {
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	PicUrl       string    `xml:"PicUrl"`
	MediaId      int       `xml:"MediaId"`
	MsgId        int64     `xml:"MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *ImageMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}

func ReceiveImageMsg(content string) string {
	var msg ImageMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
