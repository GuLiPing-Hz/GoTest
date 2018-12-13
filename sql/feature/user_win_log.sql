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

ALTER TABLE user_win_log
  ADD wealth bigint DEFAULT 0 NULL
COMMENT '那个时间的财富值';
ALTER TABLE user_win_log
  COMMENT = '用户在线实时纪录输赢值,财富值';

#道具单独领出来一个记录表
CREATE TABLE user_props_log
(
  id        bigint PRIMARY KEY AUTO_INCREMENT,
  uid       bigint NOT NULL,
  type      int                DEFAULT -1
  COMMENT '道具类型
    -1 未知道具
    jinbi: 10,//金币，
    zuanshi: 13,//钻石
    //上面两个单独放到coin_log中统计

    zhenzhu: 0,//珍珠，
    longzhu1: 1,//一星龙珠，
    longzhu2: 2,//二星龙珠，
    longzhu3: 3,//三星龙珠，
    laba: 11,//喇叭，
    longzhucard: 12,//龙珠卡
    propfrozen: 15,//冰冻道具
    propLock: 16,//锁定道具
    propDGold: 17,//双倍金币
    propDSpeed: 18,//双倍射速
    ',
  cur       int                DEFAULT 0
  COMMENT '当前的道具最新值',
  variation int                DEFAULT 0
  COMMENT '道具变化量 >0表示增加，小余0表示减少',
  optUid    bigint NOT NULL
  COMMENT '操作者ID，-1表示系统 0表示自己，其他表示用户ID',
  optedUid  bigint NOT NULL
  COMMENT '被操作者用户ID，0表示自己',
  sendType  tinyint            DEFAULT 0
  COMMENT '
        INVALID = 0, --未知类型
        DROP = 1, --掉落
        MERGE = 2, --合成
        USED = 3, --使用
        SEND = 4, --赠送
        RECV = 5, --接收
        LOTTERY = 6, --抽奖
        CDKEY = 7, --CDKey
        WEEKLYCARD = 8, --周卡
        CHONGJIDAJIANGSAI = 9, --冲级大奖赛
        LONGZHUDUOBAO = 10, --龙珠夺宝
        SPLIT = 11, --拆分
        UPGRADEVIP = 12, --升级VIP
        GAMEYULE = 13, --鱼乐游戏
        GROWTH_TASK = 14, --完成成长任务后赠送的龙珠
        MONTHLYCARD = 15, --月卡一次性奖励
        WEEKLYCARD_DAILY = 16, --周卡每日奖励
        MONTHLYCARD_DAILY = 17, --月卡每日奖励
        BUY_GIFT_PKG = 18, --购买礼包
        UPGRADELEVEL = 19, --用户经验等级升级
        VIP_DAILY = 20, --VIP用户每日领取礼物
        ACTIVITY_LEVEL = 21, --冲级活动奖励
        ACTIVITY_DIAMONDS = 22, --钻石消耗活动奖励
        ACTIVITY_SHARING = 23 --分享有礼活动奖励',
  time      datetime           DEFAULT now()
  COMMENT '变化时间'
);
ALTER TABLE user_props_log
  COMMENT = '用户道具变化表';
create index user_props_log_uid_index
  on user_props_log (uid);
create index user_props_log_uid_type
  on user_props_log (type);
