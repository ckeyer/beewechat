package weichat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"funxdata/models/global"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"io"
	"log"
	// "errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type WXOpenIDList struct {
	Openid []string `json:openid`
}
type WXUserList struct {
	Total       int          `json:total`
	Count       int          `json:count`
	Data        WXOpenIDList `json:data`
	Next_openid string       `json:next_openid`
}
type WXUserInfo struct {
	Subscribe      int    `json:"subscribe"`      //	用户是否订阅该公众号标识
	Openid         string `json:"openid"`         //	用户的标识，对当前公众号唯一
	Nickname       string `json:"nickname"`       //	用户的昵称
	Sex            int    `json:"sex"`            //	用户的性别，值为1时是男性，值为2时是女性
	City           string `json:"city"`           //	用户所在城市
	Country        string `json:"country"`        //	用户所在国家
	Province       string `json:"province"`       //	用户所在省份
	Language       string `json:"language"`       //	用户的语言，简体中文为zh_CN
	Headimgurl     string `json:"headimgurl"`     //	用户头像
	Subscribe_time int64  `json:"subscribe_time"` // 	用户关注时间，为时间戳
	Unionid        int    `json:"unionid"`        // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

///	获取用户列表并报错到redis数据库中 KEY=wx_UserList
func GetUserList() {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + GetAccessToken()
	key := "wx_UserList"
	var v WXUserList
	c, status := global.HttpGet(url)
	if status >= 0 {
		dec := json.NewDecoder(strings.NewReader(c))
		for {
			if err := dec.Decode(&v); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}

		var redcli redis.Client
		redcli.Addr = beego.AppConfig.String("redis_addr")

		if ok, err := redcli.Exists(key); err != nil {
			log.Println(err.Error())
		} else {
			if ok {
				redcli.Del(key)
			}
		}
		for _, value := range v.Data.Openid {
			redcli.Rpush(key, []byte(value))
		}
		log.Println("Successful: Add Userlist ")
	}
}

/// 获取用户基本信息，并保存到mysql 表tb_wx_user_info 中
func GetUserInfo(openid string) (this *WXUserInfo) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + GetAccessToken() + "&openid=" + openid + "&lang=zh_CN"
	c, status := global.HttpGet(url)
	this = &WXUserInfo{}
	if status >= 0 {
		dec := json.NewDecoder(strings.NewReader(c))
		for {
			if err := dec.Decode(this); err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
			}
		}
	}
	this.saveUserInfo()
	return
}
func (this *WXUserInfo) saveUserInfo() bool {
	db, err := connectDB()
	defer db.Close()
	if err != nil {
		return false
	} else {
		sql := "insert into tb_wx_user_info ( subscribe, openid, nickname, sex, " +
			"city, country, province, language, headimgurl, subscribe_time, unionid)" +
			" values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		_, e := db.Exec(sql, this.Subscribe, this.Openid, this.Nickname, this.Sex, this.City,
			this.Country, this.Province, this.Language, this.Headimgurl,
			this.Subscribe_time, this.Unionid)
		if e != nil {
			sql := "update tb_wx_user_info set subscribe=? , nickname=?, " +
				"sex=?, city=?, country=?, province=?, language=?, headimgurl=?, " +
				"subscribe_time=?, unionid=? where openid=? "
			r, e2 := db.Exec(sql, this.Subscribe, this.Nickname, this.Sex, this.City,
				this.Country, this.Province, this.Language, this.Headimgurl,
				this.Subscribe_time, this.Unionid, this.Openid)
			if e2 != nil {
				log.Println(e2)
				return false
			} else {
				log.Println(r.LastInsertId())
				// log.Println(r.RowsAffected())
			}
		} else {
			// log.Println(r.LastInsertId())
			// log.Println(r.RowsAffected())

		}
	}
	return true
}
func connectDB() (db *sql.DB, err error) {
	connStr := beego.AppConfig.String("wx_mysql_connstr")
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("database initialize error : ", err.Error())
		// db = db
	}
	return
}
func GetUserListInfo() {
	var redcli redis.Client
	redcli.Addr = beego.AppConfig.String("redis_addr")
	key := "wx_UserList"
	b, _ := redcli.Lrange(key, 0, -1)
	for _, v := range b {
		GetUserInfo(fmt.Sprintf("%s", v))
	}
}
func Test() {

}
