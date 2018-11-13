#新增邀请码配置
CREATE TABLE IF NOT EXISTS Buyu.invite_cfg (
  #定义字段名 类型 默认值 主键 是否可空 自动增加
  id           BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT
  COMMENT '自增字段',
  inviterGift  INT                NOT NULL
  COMMENT '邀请人奖励金币',
  inviteesGift INT                NOT NULL
  COMMENT '被邀请人奖励金币',
  time         DATETIME           NOT NULL
  COMMENT '指定时间之后注册的用户才能填'
)
  COMMENT '邀请活动配置'
  DEFAULT CHARSET = utf8mb4;
INSERT INTO Buyu.invite_cfg (inviterGift, inviteesGift, time) VALUES (50000, 50000, CURRENT_DATE());

#新增邀请码配置
CREATE TABLE IF NOT EXISTS Buyu.invite_log (
  uid  bigint(10) PRIMARY KEY        NOT NULL
  COMMENT '被邀请人',
  code bigint(10)                    NOT NULL
  COMMENT '邀请码/邀请人ID',
  time DATETIME DEFAULT NOW()
  COMMENT '邀请时间'
)
  COMMENT '邀请记录'
  DEFAULT CHARSET = utf8mb4;

CREATE INDEX uid
  ON Buyu.invite_log (uid);

#创建索引邀请码
CREATE INDEX code
  ON Buyu.invite_log (code);

# insert into invite_log (uid, code) values (1,1)