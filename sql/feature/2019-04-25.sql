CREATE TABLE msg
(
  id     bigint PRIMARY KEY AUTO_INCREMENT,
  type   int               NOT NULL
  COMMENT '消息类型',
  uid    bigint            NOT NULL,
  fromid bigint DEFAULT -1 NOT NULL
  COMMENT '-1表示系统，0表示自己，其他根据type定义，可能为uid，也可能为ytid。。',
  tm     datetime          NOT NULL,
  msg    text COMMENT '消息内容。'
);
ALTER TABLE msg
  COMMENT = '消息表';

CREATE TABLE yt
(
  ytid         int PRIMARY KEY AUTO_INCREMENT
  COMMENT '鱼塘id',
  uid          bigint      NOT NULL
  COMMENT '创建者uid',
  name         varchar(50) NOT NULL
  COMMENT '鱼塘名称',
  info         text COMMENT '鱼塘简介',
  reward       int             DEFAULT 0
  COMMENT '鱼塘每日签到奖励',
  limit_people int             DEFAULT 100
  COMMENT '人数上限',
  tm           datetime COMMENT '创建时间'
);
ALTER TABLE yt
  COMMENT = '鱼塘表';

CREATE TABLE yt_user
(
  uid   bigint PRIMARY KEY NOT NULL,
  ytid  int                NOT NULL
  COMMENT '鱼塘id',
  tm    datetime           NOT NULL
  COMMENT '加入时间',
  endtm datetime COMMENT '离开时间'
);
ALTER TABLE yt_user
  COMMENT = '鱼塘用户表';

CREATE TABLE yt_bill_log
(
  ytid int      NOT NULL,
  tm   datetime NOT NULL,
  bill bigint DEFAULT 0,
  CONSTRAINT yt_bill_log_ytid_tm_pk PRIMARY KEY (ytid, tm)
);
ALTER TABLE yt_bill_log
  COMMENT = '鱼塘每日流水表';

CREATE TABLE yt_coin_log
(
  id     bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  tm     datetime           NOT NULL,
  uid    bigint             NOT NULL,
  ytid   int                NOT NULL,
  reward int                NOT NULL
  COMMENT '奖励信息。或者被偷信息。',
  type   smallint           NOT NULL
  COMMENT '0 表示鱼塘签到，1表示收取鱼货，2表示偷取鱼货 (鱼货的收取和被偷起始时间都是早上 2:00 am)',
  optuid bigint             NOT NULL
  COMMENT '-1表示系统，0表示自己，如果type是2，表示被偷的uid，'
);
ALTER TABLE yt_coin_log
  COMMENT = '鱼塘奖励日志';

CREATE TABLE yt_fc_cfg
(
  rank int PRIMARY KEY NOT NULL
  COMMENT '排名。',
  coin int DEFAULT 0   NOT NULL
  COMMENT '每日签到奖励金币增加额度'
);
CREATE UNIQUE INDEX yt_fc_cfg_rank_uindex
  ON yt_fc_cfg (rank);
ALTER TABLE yt_fc_cfg
  COMMENT = '鱼塘扶持奖励列表';



