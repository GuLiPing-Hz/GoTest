ALTER TABLE user_stat
  CHANGE user_win hbq int DEFAULT '0'
COMMENT '红包券，兑换微信红包，单位分';
ALTER TABLE user_stat
  CHANGE winstat hb int DEFAULT '0'
COMMENT '微信红包，可提现，单位分';
ALTER TABLE user_stat
  MODIFY lucky_flag int NOT NULL DEFAULT '0'
  COMMENT '00000000000000000000000000000  从右往左 第一位表示是否已经兑换过2元红包';
ALTER TABLE user
  CHANGE phone_valid guide_flag tinyint DEFAULT 0
COMMENT '是否已经领取0.3元新手引导礼包。 0未领取，1领取';
update user_stat
set lucky_flag = 0;

update user_stat
set hbq = 0;
update user_stat
set hb = 0;

ALTER TABLE user_props_log
  MODIFY type int(11) DEFAULT '-1'
  COMMENT '见trello 字段说明物品类型';

CREATE TABLE hbq_cfg
(
  id          int PRIMARY KEY AUTO_INCREMENT,
  hbqid       int           NOT NULL
  COMMENT '红包券ID',
  val         int             DEFAULT 1
  COMMENT '红包券金额，单位分',
  probability int DEFAULT 1 NOT NULL
  COMMENT '抽奖概率'
);
ALTER TABLE hbq_cfg
  COMMENT = '红包券配置信息表';

INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('1', '10', '333');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('1', '20', '334');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('1', '30', '333');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('2', '100', '333');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('2', '200', '334');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('2', '300', '333');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('3', '600', '330');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('3', '1200', '330');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('3', '1800', '330');
INSERT INTO `hbq_cfg` (`hbqid`, `val`, `probability`) VALUES ('3', '8800', '10');

INSERT INTO `wares_cfg` VALUES
  ('tx5', 3, 0, '', -5, 0, 0, 0, 0, 0, 0, 0),
  ('tx10', 3, 0, '', -10, 0, 0, 0, 0, 0, 0, 0),
  ('tx20', 3, 0, '', -20, 0, 0, 0, 0, 0, 0, 0),
  ('tx50', 3, 0, '', -50, 0, 0, 0, 0, 0, 0, 0),
  ('tx100', 3, 0, '', -100, 0, 0, 0, 0, 0, 0, 0);

CREATE TABLE hbq_dui_cfg
(
  id    int PRIMARY KEY,
  price int COMMENT '需要花费的红包券数量，单位分',
  type  int COMMENT '兑换的物品类型，参见trello物品类型',
  cnt   int COMMENT '兑换的数量，如果是红包(41)，单位是分'
);
ALTER TABLE hbq_dui_cfg
  COMMENT = '红包券兑换配置信息';

INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('1', '200', '41', '200');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('2', '500', '41', '500');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('3', '2000', '41', '2000');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('4', '5000', '41', '5000');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('5', '580', '10', '58000');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('6', '5000', '10', '600000');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('7', '580', '13', '58');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('8', '5000', '13', '600');
INSERT INTO `hbq_dui_cfg` (`id`, price, `type`, `cnt`) VALUES ('9', '200', '10', '20000');

ALTER TABLE user_props_log
  MODIFY optUid bigint(20) NOT NULL
  COMMENT '操作者ID，0表示系统，其他表示用户ID，如果是sendType是红包提现（61），
这个表示订单ID；如果是是sendType是商店兑换(60)，这个表示兑换的物品类型';

ALTER TABLE user
  MODIFY phone_valid tinyint DEFAULT 0
  COMMENT '是否已经领取0.3元新手引导礼包。 0未领取，1领取';
update user
set phone_valid = 0;

ALTER TABLE user
  MODIFY type tinyint(4) DEFAULT '0'
  COMMENT '用户类型，0.普通用户1.普通推广员2.金牌推广员3.机器人，4捕鱼机器人';



