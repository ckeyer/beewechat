package msg

import (
	"github.com/astaxie/beego/orm"
)

func RegDB() {
	orm.RegisterModel(new(TextMsg),
		new(ImageMsg),
		new(LinkMsg),
		new(LocationMsg),
		new(VideoMsg),
		new(VoiceMsg))
}

func ReceiveMsg(content string, msgtype string) (r string) {
	r = ""
	switch msgtype {
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
	default:
		r = "error"
	}
	return
}
