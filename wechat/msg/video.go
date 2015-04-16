package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type VideoMsg struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	MediaId      int       `xml:"MediaId"`
	ThumbMediaId string    `xml:"ThumbMediaId"`
	MsgId        int64     `xml:"MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveVideoMsg(content string) string {
	var msg VideoMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *VideoMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}
