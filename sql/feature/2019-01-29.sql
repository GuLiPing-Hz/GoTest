ALTER TABLE user
  MODIFY ID_valid varchar(100) COMMENT '推广渠道需要的设备码 android:IMEI;ios:IDFA';

ALTER TABLE user_stat
  CHANGE lucky_room1 fuzi int NOT NULL DEFAULT 0
COMMENT '集福活动当前的褔字数量';

ALTER TABLE user_stat
  CHANGE lucky_room2 fuzi_total int NOT NULL DEFAULT 0
COMMENT '集福活动累计获得的褔字数量';

update user_stat
set user_stat.fuzi = 0;
update user_stat
set user_stat.fuzi_total = 0;

ALTER TABLE user_props_log
  MODIFY optUid bigint(20) NOT NULL
  COMMENT '操作者ID，0表示系统，其他表示用户ID，如果是sendType是红包提现（61），
这个表示订单ID；如果是是sendType是商店兑换(60)或者新年集福活动(61)，这个表示兑换ID';

ALTER TABLE coin_log
  MODIFY coin_change bigint(20) COMMENT '消耗的金币总数';

CREATE TABLE fu_cfg
(
  id   smallint PRIMARY KEY NOT NULL,
  cost int                  NOT NULL
  COMMENT '需要消耗的褔字数量',
  type smallint             NOT NULL
  COMMENT '能兑换的物品类型，参见trello中的物品字段说明',
  cnt  int DEFAULT 0        NOT NULL
  COMMENT '兑换的物品数量，红包券单位分'
);
CREATE UNIQUE INDEX fu_cfg_id_uindex
  ON fu_cfg (id);
ALTER TABLE fu_cfg
  COMMENT = '新年集福活动配置,兑换记录参见表user_props_log';

INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('1', '8', '16', '20');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('2', '28', '13', '168');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('3', '68', '32', '30');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('4', '88', '40', '180');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('5', '288', '10', '180000');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('6', '488', '43', '1');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('7', '888', '3', '1');
INSERT INTO `fu_cfg` (`id`, `cost`, `type`, `cnt`) VALUES ('8', '1688', '40', '8800');

INSERT INTO `activities` VALUES
  (8, '集福赢红包', '集福赢红包', 1, '', '', '2019-01-01 00:00:00', '2019-12-31 00:00:00', 8, '2019-01-01 00:00:00', 'N', 'Y');

