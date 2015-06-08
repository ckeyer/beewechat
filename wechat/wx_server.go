/*
 * 与腾讯服务器的相关交互
**/
package wechat

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"funxdata/models/global"
	"github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/wechat/event"
	"github.com/ckeyer/beewechat/wechat/msg"
	"github.com/hoisie/redis"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	//  redis中微信 AccessToken 的key
	REDIS_KEY_WC_ACCESS_TOKEN = "wx_AccessToken"
)

type AccessToken struct {
	Access_token string `json: "access_token"`
	Expires_in   int64  `json:"expires_in"`
}

type MsgType struct {
	MsgType string `xml:"MsgType"`
	Event   string `xml:"Event"`
}

func (this *AccessToken) Decode(jsonstr []byte) error {
	return json.Unmarshal(jsonstr, &this)
}

// 获取微信的AppID
func GetAppID() string {
	return beego.AppConfig.String("appId")
}

// 获取微信的aApp Secret
func GetAppSecret() string {
	return beego.AppConfig.String("app_secret")
}

// 更新微信的AccessToken到Redis中 key=REDIS_KEY_WC_ACCESS_TOKEN
func UpdateAccessToken() (expires_in int, err error) {
	appid := GetAppID()
	secret := GetAppSecret()
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid +
		"&secret=" + secret
	if c, status := global.HttpGet(url); status < 0 {
		err = errors.New("access_token 获取异常 " + strconv.Itoa(status))
	} else {
		log.Println(c)
		var v AccessToken
		dec := json.NewDecoder(strings.NewReader(c))
		for {
			if err = dec.Decode(&v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
				return
			}
		}

		var redcli redis.Client
		redcli.Addr = beego.AppConfig.String("redis_addr")
		expires_in = (int)(v.Expires_in)
		err = redcli.Setex(REDIS_KEY_WC_ACCESS_TOKEN, v.Expires_in, []byte(v.Access_token))
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("Successful: get AccessToken ")
		}
	}
	return
}

func AutoGetAccessToken() {
	ei, err := UpdateAccessToken()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	outtime := (time.Duration)(ei-100) * time.Second
	go time.AfterFunc(outtime, AutoGetAccessToken)
}

func GetAccessToken() string {
	var redcli redis.Client
	redcli.Addr = beego.AppConfig.String("redis_addr")
	b, e := redcli.Get(REDIS_KEY_WC_ACCESS_TOKEN)
	if e != nil {
		log.Println(e.Error())
	}
	return fmt.Sprintf("%s", b)
}

func ReceiveMsg(content string) (r string) {
	r = ""

	var msgtype MsgType
	err := xml.Unmarshal([]byte(content), &msgtype)
	if err != nil {
		return
	}
	switch msgtype.MsgType {
	// case "text", "image", "voice", "video", "location", "link":
	case "event":
		r = event.ReceiveEvent(content, msgtype.Event)
	default:
		r = msg.ReceiveMsg(content, msgtype.Event)
	}
	return
}
