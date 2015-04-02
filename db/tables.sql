CREATE DATABASE db_wx;

--- 来自微信的用户基本信息
use db_wx;
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
	created datetime DEFAULT NULL,
	updated datetime DEFAULT NULL,
	PRIMARY KEY(id ,openid )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TRIGGER IF EXISTS tg_insert_wxuserinfo;
DELIMITER $$
create trigger tg_insert_wxuserinfo before insert on tb_wx_user_info
FOR EACH ROW
BEGIN
set NEW.created =NOW();
set NEW.updated =NOW();
 END$$
DELIMITER ;

DROP TRIGGER IF EXISTS tg_update_wxuserinfo;
create trigger tg_update_wxuserinfo before update on tb_wx_user_info
FOR EACH ROW
set NEW.updated =NOW();

DROP TABLE IF EXISTS tb_wx_amr ;
CREATE TABLE tb_wx_amr(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	MediaID      varchar(32) ,
	Format       varchar(32) ,
	Recognition  varchar(32) ,
	MsgID        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_scribe_event ;
CREATE TABLE tb_wx_scribe_event(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	Event        varchar(32) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_text_msg ;
CREATE TABLE tb_wx_text_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	Event        varchar(32) ,
	Content     varchar(256) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_image_msg ;
CREATE TABLE tb_wx_image_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	PicUrl      varchar(64) ,
	MediaId      number(12) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS tb_wx_voice_msg ;
CREATE TABLE tb_wx_voice_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	Format      varchar(64) ,
	MediaId      number(12) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_video_msg ;
CREATE TABLE tb_wx_video_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	ThumbMediaId      varchar(64) ,
	MediaId      number(12) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_location_msg ;
CREATE TABLE tb_wx_location_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	Location_X      DOUBLE ,
	Location_Y  DOUBLE,
	Scale	number(8),
	Label 	 varchar(32) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS tb_wx_link_msg ;
CREATE TABLE tb_wx_link_msg(
	id int  AUTO_INCREMENT,
	ToUserName   varchar(32) ,
	FromUserName varchar(32) ,
	CreateTime   number(12) ,
	MsgType      varchar(32) ,
	Title varchar(32) ,
	Description varchar(256) ,
	Url      varchar(64) ,
	MsgId        number(12) ,
	PRIMARY KEY(id  )
) AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

