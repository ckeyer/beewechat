/*
* 文件描述: 微信JS-SDK的相关服务
* 创建日期: 2015/3/18
* 作者:  ckeyer
* 功能:  微信SDK初始化
**/
package jsapi

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ckeyer/beewechat/models/global"
	"github.com/hoisie/redis"
	logpkg "log"
	"os"
	"time"
)

var (
	redcli redis.Client
	log    *logpkg.Logger
)

const (
	// redis 中 微信 jsapi_ticket的 key
	REDIS_KEY_WC_JSAPI_TICKET = "wx_JsapiTicket"
	// redis中 微信JS-sdk的 NONCESTR key  一个随机字符串
	REDIS_KEY_WC_JSAPI_NONCESTR = "wx_jsapiNoncestr"
	// redis中 微信JS-sdk的 timestamp key  一个随机字符串
	REDIS_KEY_WC_JSAPI_TIME_STAMP = "wx_jsapiTimestamp"
)

// 用于解析腾讯服务器发来的 JS_SDK ticket信息的结构体
type JsapiTicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

func init() {
	redcli.Addr = beego.AppConfig.String("redis_addr")
	log = log.New(os.Stdout, "CKEYER - ", log.LstdFlags)
}

// 获取jsasp_ticket 字符串/
func GetJsApiTicket() string {
	ts := getJsApiTicketFromLocal()
	if ts != "" {
		return ts
	} else {
		ots, err := getJsapiTicketFromServer()
		if err != nil {
			log.Println("jsapi - " + err.Error())
			return ""
		} else {
			ots.saveToRedis()
			return ots.Ticket
		}
	}
}

// 从本地redis中获取 jsasp_ticket 字符串
func getJsApiTicketFromLocal() (ts string) {

	bs, err := redcli.Get(REDIS_KEY_WC_JSAPI_TICKET)
	if err != nil {
		return ""
	} else {
		ts = string(bs)
	}
	return
}

// 从腾讯服务器获取获取 jsapi_ticket 对象
func getJsapiTicketFromServer() (jsapi_ticket *JsapiTicket, err error) {
	url := `https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=` + GetAccessToken() + `&type=jsapi`
	jsonBlob, status := global.HttpGetToBytes(url)
	// log.Println(string(jsonBlob))
	if status > 0 {
		err = json.Unmarshal(jsonBlob, &jsapi_ticket)
		if err != nil {
			log.Println(err)
		} else {
			// log.Println(jsapi_ticket.Ticket)
			return
		}
	}
	return
}

// 将 jsapi_ticket 存储到redis中
func (this *JsapiTicket) saveToRedis() {

	err := redcli.Setex(REDIS_KEY_WC_JSAPI_TICKET, this.ExpiresIn, []byte(this.Ticket))
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Successful: get JS-SDK Ticket from server")
	}
}

// 获取 jsapi的noncestr 随机字符串
// 如果redis中存在, 则直接取出, 如果不存在, 则随机产生, 并保存到redis中, 过期时间7200秒
func GetJsApiNoncestr() string {
	bs, err := redcli.Get(REDIS_KEY_WC_JSAPI_NONCESTR)
	if err != nil {
		noncestr := global.GetRandString(16)
		err = redcli.Setex(REDIS_KEY_WC_JSAPI_NONCESTR, 7200, []byte(noncestr))
		if err != nil {
			log.Println("jsapi_noncestr-" + err.Error())
		}
		return noncestr
	}
	return string(bs)
}

// 获取 jsapi的timestamp 随机字符串
// 如果redis中存在, 则直接取出, 如果不存在, 则随机产生, 并保存到redis中, 过期时间7200秒
func GetJsApiTimeStamp() string {
	bs, err := redcli.Get(REDIS_KEY_WC_JSAPI_TIME_STAMP)
	if err != nil {
		timestamp := fmt.Sprint(time.Now().Unix())
		err = redcli.Setex(REDIS_KEY_WC_JSAPI_TIME_STAMP, 7200, []byte(timestamp))
		if err != nil {
			log.Println("jsapi_noncestr-" + err.Error())
		}
		return timestamp
	}
	return fmt.Sprintf("%s", bs)
}

// 得到用于JS_SDK的signature
func GetJsApiSignature(url string) string {
	ticket := GetJsApiTicket()
	timestamp := GetJsApiTimeStamp()
	nonce := GetJsApiNoncestr()
	tmpStr := `jsapi_ticket=` + ticket +
		`&noncestr=` + nonce +
		`&timestamp=` + timestamp +
		`&url=` + url
	return global.GetSHA1(tmpStr)
}
