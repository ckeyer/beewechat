package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

// 地理位置推送事件结构体
type LocationEvent struct {
	ToUserName   string  `xml:"ToUserName"`
	FromUserName string  `xml:"FromUserName"`
	CreateTime   int     `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	Event        string  `xml:"Event"`
	Latitude     float64 `xml:"Latitude"`
	Longitude    float64 `xml:"Longitude"`
	Precision    int     `xml:"Precision"`
}

func ReceiveLocationvent(content string) string {
	var msg LocationEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
