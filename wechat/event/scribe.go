package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

type ScribeEvent struct {
	Id           int64
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
}

func ReceiveSubscribeEvent(content string) string {
	var msg SubscribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
