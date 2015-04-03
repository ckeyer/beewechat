package weichat

import (
	"encoding/xml"
	"fmt"
	"log"
	// "strings"
	"time"
)

type MsgType struct {
	MsgType string `xml:"MsgType"`
	Event   string `xml:"Event"`
}
type TextMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}
type ImageMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	PicUrl       string `xml:"PicUrl"`
	MediaId      int    `xml:"MediaId"`
	MsgId        int64  `xml:"MsgId"`
}
type VoiceMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MediaId      string `xml:"MediaId"`
	Format       string `xml:"Format"`
	MsgId        int64  `xml:"MsgId"`
}
type VideoMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MediaId      int    `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        int64  `xml:"MsgId"`
}
type LocationMsg struct {
	ToUserName   string  `xml: "ToUserName"`
	FromUserName string  `xml: "FromUserName"`
	CreateTime   int     `xml: "CreateTime"`
	MsgType      string  `xml: "MsgType"`
	Location_X   float64 `xml: "Location_X"`
	Location_Y   float64 `xml: "Location_Y"`
	Scale        int     `xml: "Scale"`
	Label        string  `xml: "Label"`
	MsgId        int64   `xml: "MsgId"`
}
type LinkMsg struct {
	ToUserName   string `xml: "ToUserName"`
	FromUserName string `xml: "FromUserName"`
	CreateTime   int    `xml: "CreateTime"`
	MsgType      string `xml: "MsgType"`
	Title        string `xml: "Title"`
	Description  string `xml: "Description"`
	Url          string `xml: "Url"`
	MsgId        int64  `xml: "MsgId"`
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
		r = ReceiveTextMsg(content)
	case "image":
		r = ReceiveImageMsg(content)
	case "voice":
		r = ReceiveVoiceMsg(content)
	case "video":
		r = ReceivevIdeoMsg(content)
	case "location":
		r = ReceiveLocationMsg(content)
	case "link":
		r = ReceiveLinkMsg(content)
	case "event":
		r = ReceiveEvent(content, &msgtype)
	default:
		r = "error"
	}
	return
}
func ReceiveTextMsg(content string) string {
	var msg TextMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	log.Println(msg.Content)
	return msg.ReplyTextMsg(`/::D/::D
服务器维护中
/::D/::D`)
}
func (this *TextMsg) ReplyTextMsg(data string) string {
	xmldata := `<xml>
<ToUserName><![CDATA[` + this.FromUserName + `]]></ToUserName>
<FromUserName><![CDATA[` + this.ToUserName + `]]></FromUserName>
<CreateTime>` + fmt.Sprint((time.Now().Unix())) + `</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[` + data + `]]></Content>
</xml>`
	log.Println("回复：", data)
	return xmldata
}
func ReceiveImageMsg(content string) string {
	var msg ImageMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveVoiceMsg(content string) string {
	var msg VoiceMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceivevIdeoMsg(content string) string {
	var msg VideoMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveLocationMsg(content string) string {
	var msg LocationMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
func ReceiveLinkMsg(content string) string {
	var msg LinkMsg
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
