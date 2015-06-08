/*
 *  通过网页授权获取用户基本信息
**/
package wechat

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/beewechat/models/global"
	"io"
	"log"
	"strings"
)

type WebAccessToken struct {
	Access_token  string `json:"access_token"`
	Expires_in    int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Openid        string `json:"openid"`
	Scope         string `json:"scope"`
}

type WebUserInfo struct {
	Id         int64
	Openid     string   `json:"openid"`     // 用户的唯一标识
	Nickname   string   `json:"nickname"`   // 用户昵称
	Sex        int      `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string   `json:"province"`   // 用户个人资料填写的省份
	City       string   `json:"city"`       // 普通用户个人资料填写的城市
	Country    string   `json:"country"`    // 国家，如中国为CN
	Headimgurl string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小
	Privilege  []string `orm:"-"`           // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Unionid    int64    `json:"unionid"`    //
}

// 获取网页端的 AccessToken
func GetWebAccessToken(code string) *WebAccessToken {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?" +
		"appid=" + config.WECHAT_APPID +
		"&secret=" + config.WECHAT_SECRET +
		"&code=" + code +
		"&grant_type=authorization_code"
	content, status := global.HttpGet(url)
	if status >= 0 {
		log.Println(content)
		dec := json.NewDecoder(strings.NewReader(content))
		var v WebAccessToken
		for {
			if err := dec.Decode(&v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		return &v
	}
	return nil
}

// 获取用户信息
func (this *WebAccessToken) GetUserInfo() *WebUserInfo {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + this.Access_token +
		"&openid=" + this.Openid + "&lang=zh_CN"
	content, status := global.HttpGet(url)
	if status >= 0 {
		log.Println(content)
		dec := json.NewDecoder(strings.NewReader(content))
		var v WebUserInfo
		for {
			if err := dec.Decode(&v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		return &v
	}
	return nil
}

func (this *WebUserInfo) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}

// 刷新网页端 AccessToken
func (this *WebAccessToken) RefreshAccessToken() {
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + config.WECHAT_APPID +
		"&grant_type=refresh_token&refresh_token=" + this.Refresh_token
	content, status := global.HttpGet(url)
	if status >= 0 {
		log.Println(content)
		dec := json.NewDecoder(strings.NewReader(content))
		var v WebAccessToken
		for {
			if err := dec.Decode(&v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		this = &v
	}
}

/// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
