package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

type SubscribeEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     int32  `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

func ReceiveUnsubscribeEvent(content string) string {
	var msg ScribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
