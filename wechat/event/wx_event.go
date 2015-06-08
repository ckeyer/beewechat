package event

import (
	"github.com/astaxie/beego/orm"
)

func RegDB() {
	orm.RegisterModel(new(ClickEvent),
		new(LocationEvent),
		new(ScanEvent),
		new(ScribeEvent),
		new(SubscribeEvent),
		new(ViewEvent))
}

func ReceiveEvent(content string, msgtype string) string {
	switch msgtype {
	case "subscribe":
		return ReceiveSubscribeEvent(content)
	case "unsubscribe":
		return ReceiveUnsubscribeEvent(content)
	case "SCAN":
		return ReceiveScanEvent(content)
	case "LOCATION":
		return ReceiveLocationvent(content)
	case "CLICK":
		return ReceiveClickvent(content)
	case "VIEW":
		return ReceiveViewEvent(content)
	}
	return ""
}
