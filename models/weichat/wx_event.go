package weichat

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"time"
)

type ScribeEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
}
type SubscribeEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     int32  `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

// 二维码扫码事件结构体
type ScanEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

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

// 点击事件结构体
type ClickEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
}
type ViewEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
}

func ReceiveEvent(content string, msgtype *MsgType) string {
	switch msgtype.Event {
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

func ReceiveSubscribeEvent(content string) string {
	var msg SubscribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveUnsubscribeEvent(content string) string {
	var msg ScribeEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveScanEvent(content string) string {
	var this ScanEvent
	fmt.Println(content)
	err := xml.Unmarshal([]byte(content), &this)
	if err != nil {
		return ""
	}
	data := "请联系系统管理员进行身份认证"
	if "101" == this.EventKey &&
		"oecJ3jhN5usPBQMIXqc9bVP0toi4" == this.Ticket {
		data = "吼吼吼"
		var redcli redis.Client
		redcli.Addr = beego.AppConfig.String("redis_addr")
		redcli.Hset(this.Ticket, "scan", []byte("true"))
	}
	fmt.Println(data)
	rcontent := `<xml>
<ToUserName><![CDATA[` + this.FromUserName + `]]></ToUserName>
<FromUserName><![CDATA[` + this.ToUserName + `]]></FromUserName>
<CreateTime>` + fmt.Sprint((time.Now().Unix())) + `</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[` + data + `]]></Content>
</xml>`
	return rcontent
}
func ReceiveLocationvent(content string) string {
	var msg LocationEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveClickvent(content string) string {
	var msg ClickEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveViewEvent(content string) string {
	var msg ViewEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
