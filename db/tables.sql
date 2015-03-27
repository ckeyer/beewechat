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

