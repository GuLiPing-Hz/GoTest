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
  ytid   int PRIMARY KEY AUTO_INCREMENT
  COMMENT '鱼塘id',
  uid    bigint      NOT NULL
  COMMENT '创建者uid',
  name   varchar(50) NOT NULL
  COMMENT '鱼塘名称',
  info   text COMMENT '鱼塘简介',
  reward int             DEFAULT 0
  COMMENT '鱼塘每日签到奖励',
  `limit`  int             DEFAULT 100
  COMMENT '人数上限',
  tm     datetime COMMENT '创建时间',
  pool   bigint          default 0
  comment '鱼塘资金',
  act    bigint          default 0
  comment '鱼塘当前活跃度，根据yt_clear_cfg来判断是否需要清零',
  ver    bigint          default 0
  COMMENT '鱼塘信息版本。每当鱼塘信息变更了，ver++'
);
ALTER TABLE yt
  COMMENT = '鱼塘表';

CREATE TABLE yt_user
(
  uid     bigint PRIMARY KEY NOT NULL,
  ytid    int                NOT NULL
  COMMENT '鱼塘id',
  tm      datetime           NOT NULL
  COMMENT '加入时间',
  yuhuo   bigint default 0
  comment '累计鱼货，离开鱼塘马上清零',
  checkin int    default 0
  comment '累计签到次数，离开鱼塘马上清零',
  endtm   datetime COMMENT '离开时间-保护期24小时，此期间无法再申请加入其它鱼塘'
);
ALTER TABLE yt_user
  COMMENT = '鱼塘用户表';

CREATE TABLE yt_bill_log
(
  ytid int      NOT NULL,
  tm   datetime NOT NULL,
  bill bigint DEFAULT 0
  COMMENT '鱼塘流水',
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
  COMMENT '0表示鱼塘签到，1表示偷取鱼货 (鱼货的收取和被偷起始时间都是早上 2:00 am)',
  optuid bigint             NOT NULL
  COMMENT '-1表示系统，0表示自己，如果type是2，表示被偷的uid，'
);
ALTER TABLE yt_coin_log
  COMMENT = '鱼塘次数日志记录';

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

create table yt_clear_cfg
(
  starttm datetime default CURRENT_TIMESTAMP not null
  comment '起始日期',
  life    smallint(6)                        not null
  comment '周期，以天为单位'
)
  comment '鱼塘清理活跃度配置。鱼塘排名是按活跃度来的。';





