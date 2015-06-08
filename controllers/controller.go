package controllers

import (
	"github.com/ckeyer/beewechat/conf"
)

var (
	config *conf.CkConfig
)

func init() {
	config = conf.NewConfig()
}
