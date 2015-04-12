CREATE DATABASE db_weichat;

--- 来自微信的用户基本信息
use db_weichat;
DROP TABLE IF EXISTS tb_wx_user_info ;
CREATE TABLE tb_wx_user_info (
	id int  AUTO_INCREMENT,
	subscribe   int DEFAULT 0,
	openid   varchar(28)  unique,
	nickname   varchar(64) DEFAULT "",
	sex   int DEFAULT 0,
	city   varchar(64) DEFAULT "",
	country   varchar(64) DEFAULT "",
	province   varchar(64) DEFAULT "",
	language   varchar(16) DEFAULT "",
	headimgurl   varchar(128) DEFAULT "",
	subscribe_time   timestamp DEFAULT 0,
	Unionid        int  DEFAULT 0,
	PRIMARY KEY(id ,openid )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_amr ;
CREATE TABLE tb_wx_amr(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	MediaID      varchar(32) ,
	Format       varchar(32) ,
	Recognition  varchar(32) ,
	MsgID        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_scribe_event ;
CREATE TABLE tb_wx_scribe_event(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	Event        varchar(32) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_text_msg ;
CREATE TABLE tb_wx_text_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	Event        varchar(32) ,
	Content     varchar(256) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_image_msg ;
CREATE TABLE tb_wx_image_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	PicUrl      varchar(64) ,
	MediaId      int(12) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS tb_wx_voice_msg ;
CREATE TABLE tb_wx_voice_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	Format      varchar(64) ,
	MediaId      int(12) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_video_msg ;
CREATE TABLE tb_wx_video_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	ThumbMediaId      varchar(64) ,
	MediaId      int(12) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_location_msg ;
CREATE TABLE tb_wx_location_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	Location_X      DOUBLE ,
	Location_Y  DOUBLE,
	Scale	int(8),
	Label 	 varchar(32) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_link_msg ;
CREATE TABLE tb_wx_link_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   int(12) ,
	MsgType      varchar(32) ,
	Title varchar(32) ,
	Description varchar(256) ,
	Url      varchar(64) ,
	MsgId        int(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;



insert into tb_wx_amr(ToUserName, FromUserName, CreateTime, MsgType, MediaID, Format, Recognition, MsgID)
	values("asb","abd",1234567,"abc","asdfb","asdfb","asdf",123456);

insert into tb_wx_scribe_event(ToUserName,FromUserName,CreateTime,MsgType,Event)
	values("asb","abd",1234567,"abc","asdfb");

insert into tb_wx_text_msg(ToUserName, FromUserName, CreateTime, MsgType, Event, Content, MsgId)
	values("asb","abd",1234567,"abc","asdfb","asdfb",123456);

insert into tb_wx_image_msg(ToUserName, FromUserName, CreateTime, MsgType, PicUrl, MediaId, MsgId)
	values("asb","abd",1234567,"abc","asdfb",123456,123456);

insert into tb_wx_voice_msg(ToUserName, FromUserName, CreateTime, MsgType, Format, MediaId, MsgId)
	values("asb","abd",1234567,"abc","asdfb",123456,123456);

insert into tb_wx_video_msg(ToUserName, FromUserName, CreateTime, MsgType, ThumbMediaId, MediaId, MsgId)
	values("asb","abd",1234567,"abc","asdfb",123456,123456);

insert into tb_wx_location_msg(ToUserName, FromUserName, CreateTime, MsgType, Location_X, Location_Y, Scale, Label, MsgId)
	values("asb","abd",1234567,"abc",123456,123456,123456,"abc",123456);

insert into tb_wx_link_msg(ToUserName, FromUserName, CreateTime, MsgType, Title, Description, Url, MsgId)
	values("asb","abd",1234567,"abc","asdfb","abc","asdfb",123456);

