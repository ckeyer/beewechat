package msg

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beewechat/wechat/event"
)

func RegDB() {
	orm.RegisterModel(new(TextMsg),
		new(ImageMsg),
		new(LinkMsg),
		new(LocationMsg),
		new(VideoMsg),
		new(VoiceMsg))
}

type MsgType struct {
	MsgType string `xml:"MsgType"`
	Event   string `xml:"Event"`
}

func ReceiveMsg(content string) (r string) {
	r = ""

	var msgtype MsgType
	err := xml.Unmarshal([]byte(content), &msgtype)
	if err != nil {
		return
	}
	switch msgtype.MsgType {
	case "text":
		r = ReceiveTextMsg(content)
	case "image":
		r = ReceiveImageMsg(content)
	case "voice":
		r = ReceiveVoiceMsg(content)
	case "video":
		r = ReceiveVideoMsg(content)
	case "location":
		r = ReceiveLocationMsg(content)
	case "link":
		r = ReceiveLinkMsg(content)
	case "event":
		r = event.ReceiveEvent(content, msgtype.Event)
	default:
		r = "error"
	}
	return
}
