#停服，把redis数据推送到数据库，清理redis的usr_数据


#把老玩家的龙珠红包引导置为已完成
update user_stat
set guide_flag = guide_flag | 2;

alter table notice modify mail_giftid varchar(100) null comment '如果是邮件型的消息带有礼物id';

create index notice_mail_giftid_index
	on notice (mail_giftid);

alter table mail_reward modify mail_giftid varchar(100) null comment '礼物编号';

create index mail_reward_mail_giftid_index
	on mail_reward (mail_giftid);

ALTER TABLE cdkey_log2 ADD uuid varchar(60) NULL COMMENT '设备码';

ALTER TABLE ylc_cfg MODIFY gameId int(11) COMMENT '游戏ID 1捕鱼 2英雄传 3渔乐，4转盘';
INSERT INTO Buyu.ylc_cfg (id, gameId, vipLimit, levelLimit) VALUES (3, 4, 0, 0);

