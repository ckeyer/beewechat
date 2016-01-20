package wechat

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/models/global"
	"github.com/hoisie/redis"
	"io"
	"log"
	"strconv"
	"strings"
)

type Ticket struct {
	Ticket         string `json : "ticket"`         //	获取的二维码ticket
	Expire_seconds int    `json : "expire_seconds"` //二维码的有效时间，以秒为单位。最大不超过1800。
	Url            string `json : "url"`
}

func GetTempTicket(expire_seconds int, scene_id int, scene_str string) string {
	url := `https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=` + GetAccessToken()
	data := `{"expire_seconds": ` + strconv.Itoa(expire_seconds) +
		`, "action_name": "QR_SCENE", ` +
		`"action_info": {"scene": {"scene_id": ` + strconv.Itoa(scene_id) +
		`,"scene_str":"` + scene_str + `"}}}`
	if ticket := getTicket(url, data); len(ticket) < 1 {
		return ""
	} else {
		var redcli redis.Client
		redcli.Addr = beego.AppConfig.String("redis_addr")
		redcli.Hset(ticket, "scene_id", []byte(strconv.Itoa(scene_id)))
		redcli.Expire(ticket, 120)
		return ticket
	}
}
func GetTempQrUrl(expire_seconds int, scene_id int, scene_str string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + GetTempTicket(expire_seconds, scene_id, scene_str)
}
func getTicket(url string, data string) string {
	c, status := global.HttpPost(url, data)
	if status < 0 {
		log.Fatal(status)
	}
	var v Ticket
	dec := json.NewDecoder(strings.NewReader(c))
	for {
		if err := dec.Decode(&v); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return ""
		}
	}
	return v.Ticket
}
