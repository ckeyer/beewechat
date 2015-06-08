package event

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
	"time"
)

// 地理位置推送事件结构体
type LocationEvent struct {
	Id           int64
	ToUserName   string    `xml:"ToUserName"`
	FromUserName string    `xml:"FromUserName"`
	CreateTime   int       `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Event        string    `xml:"Event"`
	Latitude     float64   `xml:"Latitude"`
	Longitude    float64   `xml:"Longitude"`
	Precision    int       `xml:"Precision"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
}

func ReceiveLocationvent(content string) string {
	var msg LocationEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (this *LocationEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}
