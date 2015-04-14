package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

type VoiceMsg struct {
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	MediaId      string    `xml:"MediaId"`
	Format       string    `xml:"Format"`
	MsgId        int64     `xml:"MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveVoiceMsg(content string) string {
	var msg VoiceMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *VoiceMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}
