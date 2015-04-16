package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

// 点击事件结构体
type ClickEvent struct {
	Id           int64
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
}

func ReceiveClickvent(content string) string {
	var msg ClickEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
