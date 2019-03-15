#下面这个sql语句需要服务器停止运行后才能执行。。否则渔乐机器人还是会自动加钱。。
UPDATE user_stat
set coin = 100000
WHERE uid in (SELECT uid
              FROM Buyu.user
              WHERE user.type = 3);

#修改周月卡逻辑，
ALTER TABLE user
  CHANGE ufreeze week_card_utc bigint(20) DEFAULT 0
COMMENT '周卡截止时间，utc时间戳，单位秒';
ALTER TABLE user
  CHANGE plat_src month_card_utc bigint(20) DEFAULT 0
COMMENT '月卡截止时间，utc时间戳，单位秒';
update user
set user.week_card_utc = 0, user.month_card_utc = 0;
update user
  inner join card_log on user.uid = card_log.uid
set week_card_utc = end_tm
where wares_id = 'lailai.fish.sevenday' and state = 1;
update user
  inner join card_log on user.uid = card_log.uid
set month_card_utc = end_tm
where wares_id = 'lailai.fish.thirtyday' and state = 1;

ALTER TABLE user_stat
  MODIFY daily_reward int(11) NOT NULL DEFAULT '0'
  COMMENT '废弃。。。字段。';

CREATE TABLE cdkey_cfg
(
  id      char(15) PRIMARY KEY NOT NULL,
  endtime datetime             NOT NULL
  COMMENT '该CDKey截止时间',
  gift    text                 NOT NULL
  COMMENT '礼物,范例[[物品类型,物品数量],] 是一个json数据结构'
);

CREATE TABLE cdkey_log2
(
  id   char(15) NOT NULL
  COMMENT 'cdkey 编码',
  uid  bigint   NOT NULL,
  time datetime NOT NULL
  COMMENT '领取时间',
  CONSTRAINT cdkey_log2_id_uid_pk PRIMARY KEY (id, uid)
);

drop table yule_betlog;