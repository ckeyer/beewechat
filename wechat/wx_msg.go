package wechat

import (
	"apibaiwandun/weichat/msg"
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	// "log"
	// "strings"
	// "time"
)

func RegDB() {
	orm.RegisterModel(new(msg.TextMsg), new(msg.ImageMsg), new(msg.LinkMsg), new(msg.LocationMsg), new(msg.VideoMsg), new(msg.VoiceMsg))
	// orm.RegisterModel()
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/lab204?charset=utf8")

}

type MsgType struct {
	MsgType string `xml:"MsgType"`
	Event   string `xml:"Event"`
}

func ReceiveMsg(content string) (r string) {
	r = ""
	// content = strings.Replace(content, "<![CDATA[", "", -1)
	// content = strings.Replace(content, "]]>", "", -1)

	var msgtype MsgType
	err := xml.Unmarshal([]byte(content), &msgtype)
	if err != nil {
		return
	}
	switch msgtype.MsgType {
	case "text":
		r = msg.ReceiveTextMsg(content)
	case "image":
		r = msg.ReceiveImageMsg(content)
	case "voice":
		r = msg.ReceiveVoiceMsg(content)
	case "video":
		r = msg.ReceiveVideoMsg(content)
	case "location":
		r = msg.ReceiveLocationMsg(content)
	case "link":
		r = msg.ReceiveLinkMsg(content)
	case "event":
		r = ReceiveEvent(content, &msgtype)
	default:
		r = "error"
	}
	return
}
