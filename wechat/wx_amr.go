package wechat

type WX_AMR struct {
	Id           int64
	ToUserName   string `xml : "ToUserName"`   //开发者微信号
	FromUserName string `xml : "FromUserName"` //发送方帐号（一个OpenID）
	CreateTime   string `xml : "CreateTime"`   //消息创建时间 （整型）
	MsgType      string `xml : "MsgType"`      //语音为voice
	MediaID      string `xml : "MediaID"`      //语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format       string `xml : "Format"`       //语音格式：amr
	Recognition  string `xml : "Recognition"`  //语音识别结果，UTF8编码
	MsgID        int64  `xml : "MsgID"`        //消息id，64位整型
}
