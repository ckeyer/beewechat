package wechat

import (
	// "encoding/xml"
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/wechat/event"
	// "github.com/hoisie/redis"
	// "time"
)

func ReceiveEvent(content string, msgtype *MsgType) string {
	switch msgtype.Event {
	case "subscribe":
		return event.ReceiveSubscribeEvent(content)
	case "unsubscribe":
		return event.ReceiveUnsubscribeEvent(content)
	case "SCAN":
		return event.ReceiveScanEvent(content)
	case "LOCATION":
		return event.ReceiveLocationvent(content)
	case "CLICK":
		return event.ReceiveClickvent(content)
	case "VIEW":
		return event.ReceiveViewEvent(content)
	}
	return ""
}
