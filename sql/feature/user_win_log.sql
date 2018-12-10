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


