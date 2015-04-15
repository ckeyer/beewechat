package msg

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

var CKLog *logs.BeeLogger

func init() {
	CKLog = beego.BeeLogger
	CKLog.SetLevel(beego.LevelDebug)
}

type TextMsg struct {
	Id           int64     `orm:"pri"`
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Content      string    `xml:"Content"`
	MsgId        int64     `xml:"MsgId"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *TextMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
	return nil
}

func ReceiveTextMsg(content string) string {
	var msg TextMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	msg.Insert()
	CKLog.Debug(msg.Content)
	return msg.ReplyTextMsg(`/::D/::D
服务器维护中
/::D/::D`)
}
func (this *TextMsg) ReplyTextMsg(data string) string {
	xmldata := `<xml>
<ToUserName><![CDATA[` + this.FromUserName + `]]></ToUserName>
<FromUserName><![CDATA[` + this.ToUserName + `]]></FromUserName>
<CreateTime>` + fmt.Sprint((time.Now().Unix())) + `</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[` + data + `]]></Content>
</xml>`
	CKLog.Info("回复：%s", data)
	return xmldata
}

// func FindUserByName(uname string) (*TextMsg, error) {
// 	o := orm.NewOrm()
// 	user := new(User)
// 	qs := o.QueryTable("user")
// 	err := qs.Filter("username", uname).One(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
