CREATE TABLE IF NOT EXISTS Buyu.user_win_log (
  id   BIGINT PRIMARY KEY              NOT NULL AUTO_INCREMENT,
  uid  BIGINT                          NOT NULL
  COMMENT '用户ID',
  time DATETIME                                 DEFAULT NOW()
  COMMENT '纪录时间',
  win  bigint                                   default 0
  comment '那个时间的输赢值'
)
  COMMENT '用户在线实时纪录输赢值'
  DEFAULT CHARSET = utf8mb4;

create index uid
  on Buyu.user_win_log (uid);


CREATE TABLE IF NOT EXISTS Buyu.user_wealth_log
(
  id     bigint PRIMARY KEY AUTO_INCREMENT,
  uid    bigint NOT NULL
  COMMENT '用户id',
  wealth bigint             DEFAULT 0
  COMMENT '用户那个时间点的财富值',
  time   datetime           DEFAULT now()
  COMMENT '记录时间'
);
ALTER TABLE user_wealth_log
  COMMENT = '用户财富变化表，统计金币+龙珠';
create index user_wealth_log_uid_index
  on user_wealth_log (uid);

ALTER TABLE user_win_log ADD wealth bigint DEFAULT 0 NULL COMMENT '那个时间的财富值';
ALTER TABLE user_win_log COMMENT = '用户在线实时纪录输赢值,财富值';
