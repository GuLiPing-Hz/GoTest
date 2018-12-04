-- auto-generated definition
create table act_log
(
  uid     bigint(255) not null
  comment '玩家ID',
  mid     int         not null
  comment 'mission id',
  addtime datetime    not null
  comment 'add time',
  act     bigint(255) null
  comment 'mission act',
  act0    bigint(255) null
  comment 'before add mission act',
  act1    bigint(255) null
  comment 'after add mission act'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table active_coin
(
  uid       bigint(32)      not null
  comment 'user id',
  act       int default '0' not null
  comment 'active',
  state     int default '0' null
  comment '0 unreceive,1 receive',
  update_tm datetime        null
  comment 'update time',
  primary key (uid, act)
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table activities
(
  id          int(20) unsigned auto_increment
  comment '活动id'
    primary key,
  title       varchar(30) charset utf8      null
  comment '活动标题，10个汉字以内',
  activityDes varchar(255) charset utf8     null
  comment '活动描述，20个汉字以内',
  sub_id      int                           null
  comment '子活动编号',
  imgUrl      varchar(255) charset utf8     null
  comment '图片地址',
  activityUrl varchar(255) charset utf8     null
  comment '活动地址',
  startTime   datetime                      null
  comment '开始时间',
  endTime     datetime                      null
  comment '结束时间',
  sortNo      int                           null
  comment '排序号',
  createTime  datetime                      null
  comment '创建时间',
  isDeleted   char charset utf8 default 'N' null
  comment '是否删除',
  isEnabled   char charset utf8 default 'Y' null
  comment '是否启动'
)
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table add_coin_log
(
  id        bigint unsigned auto_increment
    primary key,
  uid       bigint   null,
  coin      bigint   null,
  pool_coin bigint   null,
  tm        datetime null
);

create index tm
  on add_coin_log (tm);

-- auto-generated definition
create table agent_addition
(
  type     int default '0' not null
  comment '0:vip 1:代理',
  id       int default '0' not null,
  addition int default '0' not null
  comment '加成',
  primary key (id, type)
)
  comment '推广员加成'
  charset = utf8;

-- auto-generated definition
create table agent_info
(
  Id         int auto_increment
    primary key,
  uid        varchar(255) null
  comment 'id',
  bind_uid   varchar(255) null
  comment '推广员上级id',
  nick_name  varchar(255) null
  comment '昵称',
  phone      varchar(255) null
  comment '电话',
  type       bigint(11)   null
  comment '用户类型',
  isEnabled  int          null
  comment '用户状态',
  updateTime datetime(6)  null,
  createTime datetime(6)  null
)
  comment '推广员信息表'
  charset = utf8;

-- auto-generated definition
create table agent_item
(
  itemNo      int auto_increment
  comment 'ID'
    primary key,
  createTime  timestamp default CURRENT_TIMESTAMP not null
  comment '创建时间',
  isEnabled   varchar(2) default '1'              not null
  comment '是否启用',
  itemName    varchar(50) default ''              not null
  comment '商品名称',
  displayName varchar(50) default ''              not null
  comment '页面展示名称',
  price       varchar(11) default '0'             not null
  comment '商品金额',
  coin        bigint default '0'                  not null
  comment '金币数',
  sortNo      int default '0'                     not null
  comment '排序号'
)
  comment '代理商品列表'
  charset = utf8;

-- auto-generated definition
create table agent_pay_order
(
  orderId       bigint auto_increment
  comment '订单ID'
    primary key,
  createTime    timestamp default CURRENT_TIMESTAMP not null
  comment '创建时间',
  uid           bigint default '0'                  not null
  comment '用户ID',
  orderNo       varchar(50) default ''              not null
  comment '订单编号',
  itemNo        int default '0'                     not null
  comment '商品ID',
  itemName      varchar(50) default ''              not null
  comment '商品名称',
  price         int default '0'                     not null
  comment '商品金额',
  status        tinyint(3) default '0'              not null
  comment '状态：0未支付、1已支付',
  payType       tinyint(3) default '0'              not null
  comment '支付类型：0微信、1支付宝',
  money         int default '0'                     not null
  comment '支付金额',
  payTime       datetime                            null
  comment '支付时间',
  payOrderNo    varchar(50) default ''              not null
  comment '支付订单编号',
  payInfo       text                                null
  comment '支付信息',
  extend        varchar(255) default ''             null
  comment '扩展属性',
  promotionCoin int default '0'                     null
  comment '推广金币',
  vipCoin       int default '0'                     null
  comment 'vip金币'
)
  comment '支付订单'
  charset = utf8;

-- auto-generated definition
create table agent_promotion_reward
(
  uid        bigint(11)                             null
  comment '用户id',
  num        bigint(11)                             null
  comment '自身奖励',
  createTime varchar(10) charset utf8mb4 default '' not null
  comment '创建时间',
  isExamine  varchar(255) default '0'               null
  comment '是否审核0未'
)
  comment '用户推广奖励表(每天算一次)'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table agent_recharge_cfg
(
  fixedAmount int default '0'     null
  comment '固定额度',
  plus        int(10) default '0' null
  comment '下线加成'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table agent_recharge_limit
(
  id            int(11) unsigned auto_increment
  comment '自增序列号'
    primary key,
  uid           bigint   null
  comment '用户id',
  rechargeLimit int      null
  comment '充值限额',
  createTime    datetime null
  comment '时间'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table agent_reward
(
  uid         bigint(11)                             null
  comment '用户id',
  selfReward  bigint(11)                             null
  comment '自身奖励',
  lowerReward bigint(11)                             null
  comment '下级奖励',
  createTime  varchar(10) charset utf8mb4 default '' not null
  comment '创建时间',
  isExamine   varchar(255) default '0'               null
  comment '是否审核0未'
)
  comment '用户奖励表(每天算一次)'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table agent_reward_log
(
  id          bigint auto_increment
    primary key,
  uid         bigint(16) default '0'              not null
  comment '用户id',
  num         int                                 null
  comment '奖励值',
  isGive      varchar(255) default '0'            null
  comment '是否领取',
  createTime  timestamp default CURRENT_TIMESTAMP not null
  comment '奖励添加时间',
  receiveTime timestamp default CURRENT_TIMESTAMP not null
  comment '领取时间',
  type        int                                 null
  comment '奖励类型1：消耗奖励2：推广奖励'
)
  comment '奖励领取记录'
  charset = utf8;

-- auto-generated definition
create table agent_share
(
  id          int(20) unsigned auto_increment
  comment '活动id'
    primary key,
  title       varchar(30) charset utf8            null
  comment '活动主题',
  imgUrl      varchar(255) charset utf8           null
  comment '下载图片',
  exampleUrl  varchar(255) charset utf8mb4        null
  comment '示例图片',
  activityDes varchar(255) charset utf8           null
  comment '文案内容',
  startTime   datetime                            null
  comment '开始时间',
  endTime     datetime                            null
  comment '结束时间',
  createTime  timestamp default CURRENT_TIMESTAMP not null
  comment '创建时间',
  isEnabled   char charset utf8 default 'Y'       null
  comment '是否启动'
)
  comment '分享活动表'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table agent_share_log
(
  id          bigint auto_increment
  comment 'ID'
    primary key,
  createTime  timestamp default CURRENT_TIMESTAMP not null
  comment '提交时间',
  uid         bigint                              null
  comment '用户ID',
  activityId  varchar(255)                        null,
  imgUrl      varchar(255) default ''             not null
  comment '图片',
  status      int default '0'                     not null
  comment '状态0：未审核 1：通过 2：不通过',
  examineTime varchar(255)                        null
  comment '审核时间'
)
  comment '代理分享活动记录'
  charset = utf8;

-- auto-generated definition
create table alms_cfg
(
  id         int(11) unsigned auto_increment
    primary key,
  coin       int default '8000'     not null
  comment '救济金',
  count_down int default '30'       not null
  comment '等待时间，单位秒',
  wares_id   varchar(32) default '' not null
  comment '弹出对应的商品id',
  price      int default '0'        not null
  comment '价格',
  pay_coin   int default '0'        not null
  comment '购买获得的金币数'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;



-- auto-generated definition
create table alms_log
(
  id       int(11) unsigned auto_increment
  comment '自增序列号'
    primary key,
  uid      bigint                    null
  comment '用户id',
  uuid     varchar(128) charset utf8 null
  comment '领取奖励的uuid',
  login_ip varchar(20) charset utf8  null
  comment '领取奖励的ip地址',
  coin     bigint                    null
  comment '领取救济金的金币数量',
  tm       datetime                  null
  comment '领取时间',
  platform varchar(128) charset utf8 null
  comment '设备类型',
  exp      bigint default '0'        not null
  comment '破产的时候的流水值，用于破产保护机制',
  charge   tinyint default '0'       not null
  comment '是否是充值领取破产救济金;0没充值，1充值'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table apppush_log
(
  id        int(11) unsigned auto_increment
    primary key,
  to_uid    varchar(255) null
  comment '接收者用户id',
  dev_type  varchar(255) null
  comment 'andriod或ios',
  push_type varchar(255) null
  comment 'notification',
  send_type tinyint      null
  comment '0.私推1.群发',
  ticker    varchar(255) null
  comment '通知栏提示文字',
  title     varchar(255) null
  comment '通知标题',
  content   varchar(255) null
  comment '通知文字描述',
  curtime   datetime     null
  comment '当前时间',
  checksum  varchar(255) null
  comment '校验值'
);

-- auto-generated definition
create table awards_cfg
(
  aid    int                       null
  comment 'id',
  type   int                       null
  comment '奖励类型',
  `desc` varchar(255) charset utf8 null
  comment '奖励描述',
  coin   int                       null
  comment '奖励金币数量',
  sn     int                       null
  comment '序列号／顺序号',
  switch int(10)                   null
  comment '签到开关'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table awards_log
(
  id       int(11) unsigned auto_increment
  comment '自增序列号'
    primary key,
  uid      bigint                    null
  comment '用户id',
  uuid     varchar(128) charset utf8 null
  comment '领取奖励的uuid',
  login_ip varchar(20) charset utf8  null
  comment '领取奖励的ip地址',
  coin     bigint                    null
  comment '领取奖励的金币数量',
  tm       datetime                  null
  comment '领取时间',
  platform varchar(128) charset utf8 null
  comment '设备类型'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table card_log
(
  id       bigint unsigned auto_increment
    primary key,
  oid      bigint default '0'      not null
  comment '订单id',
  uid      bigint default '0'      not null
  comment '持有人id',
  wares_id varchar(255) default '' not null
  comment '周卡编号，默认8',
  state    tinyint default '0'     not null
  comment '是否有效，0.表示无效1.表示失效',
  start_tm bigint default '0'      not null
  comment '生效时间,utc,单位毫秒',
  end_tm   bigint default '0'      not null
  comment '失效时间,utc,单位毫秒',
  buy_time datetime                null
  comment '购买时间'
);

-- auto-generated definition
create table CDKey_basic
(
  id       int(11) unsigned auto_increment
    primary key,
  CDKey    varchar(255)            null
  comment '生成的CDKey值',
  isused   tinyint                 null
  comment '0.未使用；1.已使用',
  type_id  int                     null
  comment '类型批次标志',
  note     varchar(255) default '' not null
  comment '生成本批次CDKey的用意，必写项',
  state    tinyint                 null
  comment '0.封禁1.启用',
  addTime  datetime                null
  comment '生成时间',
  operator int                     null
  comment '操作人员序号,自定义1.管理员'
);

-- auto-generated definition
create table CDKey_log
(
  id      int(11) unsigned auto_increment
    primary key,
  type_id int          null
  comment '批次号',
  CDKey   varchar(255) null,
  uid     bigint       null
  comment '使用者uid',
  addTime datetime     null
  comment '使用时间'
);

-- auto-generated definition
create table CDKey_TypeDetails
(
  id         int(11) unsigned auto_increment
    primary key,
  type_id    int null
  comment '类型编号',
  gift_id    int null
  comment '礼物类型编号,自定义;0.珍珠；1.db1;2.db2;3.db3;10.金币,11.喇叭,12.龙珠卡,13钻石',
  gift_count int null
  comment '礼物数量'
);

-- auto-generated definition
create table coin_bonus
(
  id         bigint(255) null
  comment '奖励序号',
  coin_bonus bigint(255) null
  comment '金币奖励'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table coin_log
(
  id             int(11) unsigned auto_increment
    primary key,
  uid            bigint(10)         null
  comment '用户id',
  add_time       datetime           null
  comment '添加时间',
  coin_before    bigint             null
  comment '更改前金币数量',
  coin_after     bigint             null
  comment '更改后金币数量',
  coin_change    bigint             null
  comment '更改金币数',
  diamond_before bigint default '0' not null
  comment '更改前钻石数量',
  diamond_after  bigint default '0' not null
  comment '更改后钻石数量',
  diamond_change bigint default '0' not null
  comment '更改钻石数',
  fee            bigint             null
  comment '正常房间的打中鱼获得的金币；PK房间没一局的手续费；',
  change_type    int(10)            null
  comment '0,每日奖励金币;\n1,每日任务\n2,捕鱼游戏中获得的奖励\n3,捕鱼游戏对战房间奖金\n4,捕鱼游戏对战房间准备金、手续费扣除\n5,捕鱼游戏对战房间准备金、手续费退回\n6,后台增加金币\n7,微信公众号充值\n8,爱贝支付充值\n9,苹果支付\n10,首次充值iapppay中的5元时，赠送的金币值。\n11,新注册用户七天奖励\n12,救济金\n13,龙珠夺宝奖励\n14,星启天充值\n15,支付宝充值\n16,抽奖游戏\n17,CDKey兑换\n18,周卡\n19,冲级大奖赛\n20,龙珠夺宝\n21,vip升级\n22,鱼乐游戏\n23,官网充值（官网没有充值，这个是指代理后台充值。）\n24,机器人添加\n25,代理充值奖励\n26,排行奖励\n27,服务费\n28,龙珠卡使用\n29,vip每日奖励\n30,水浒传Slot押注\n31,水浒传Slot和Bonus奖励\n32,水浒传比倍押注\n33,水浒传比倍奖励 34购买炮台消耗的金币 35使用双倍金币消耗钻石记录 36-使用双倍射速消耗钻石记录 37成长任务获得金币钻石数量 38在线奖励 39解锁炮倍，减少钻石，增加金币 40VIP奖励 41LEVEL奖励 42使用锁定道具消耗钻石记录 43使用冻结道具消耗钻石记录 44-使用龙珠卡消耗钻石记录 45-周卡每日奖励 46-月卡奖励 47-月卡每日奖励 48-商城中购买道具 49-冲级活动奖励 50-钻石消耗活动奖励 51-分享有礼活动奖励 52-未知类型的邮件奖励 53系统邮件',
  game_type      int(10)            null
  comment '0大厅，1捕鱼游戏，2其他，3抽奖，4 CDKey兑换，5周卡，6冲级大奖赛，7龙珠夺宝，8鱼乐游戏，9水浒传',
  room_id        int(10)            null
  comment '房间id',
  round          int default '0'    null
  comment '局数，pk房有效，slot有效',
  oid            bigint default '0' null
  comment '订单id，充值造成的金币、钻石变化才有效',
  task_id        int default '0'    not null
  comment 'change_type为37时有效，成长任务日志id'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index add_time
  on coin_log (add_time);

create index change_type
  on coin_log (change_type);

create index oid
  on coin_log (oid);

create index room_id
  on coin_log (room_id);

create index round
  on coin_log (round);

create index uid
  on coin_log (uid);

-- auto-generated definition
create table consumption_log
(
  add_time datetime null,
  uid      int      null,
  consume  int      null,
  room_id  int      null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;


-- auto-generated definition
create table contact
(
  c1 varchar(128) charset utf8 null,
  c2 varchar(128) charset utf8 null,
  c3 varchar(128) charset utf8 null,
  c4 varchar(128) charset utf8 null,
  c5 varchar(128) charset utf8 null,
  c6 varchar(128) charset utf8 null,
  c7 varchar(128) charset utf8 null,
  c8 varchar(128) charset utf8 null,
  c9 varchar(128) charset utf8 null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table day_coin_log
(
  id          int(11) unsigned auto_increment
    primary key,
  uid         bigint(10)          null,
  add_time    datetime            null,
  coin_before bigint(255)         null,
  coin_after  bigint(255)         null,
  coin_change bigint(255)         null,
  fee         bigint(255)         null,
  change_type int(10)             null,
  game_type   int(10)             null,
  room_id     int(10)             null,
  round       int default '0'     null,
  oid         bigint default '0'  null,
  type        tinyint default '0' not null
  comment '用户类型，0.普通用户1.普通推广员2.金牌推广员'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index add_time
  on day_coin_log (add_time);

create index change_type
  on day_coin_log (change_type);

create index oid
  on day_coin_log (oid);

create index room_id
  on day_coin_log (room_id);

create index round
  on day_coin_log (round);

-- auto-generated definition
create table day_coinSummary_log
(
  uid         bigint(10)                          null,
  stats_date  varchar(19) charset utf8 default '' not null
  comment '统计时间',
  coin_before bigint(255)                         null,
  coin_after  bigint(255)                         null,
  coin_change bigint(255)                         null,
  fee         bigint(255)                         null,
  change_type int(10)                             null,
  game_type   int(10)                             null,
  room_id     int(10)                             null,
  round       int default '0'                     null,
  oid         bigint default '0'                  null,
  type        tinyint default '0'                 not null
  comment '用户类型，0.普通用户1.普通推广员2.金牌推广员'
)
  comment '金币每日数据'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table day_longzhu_log
(
  stats_date     datetime    null
  comment '统计时间',
  transfer_count bigint(255) null,
  transfer_type  bigint(255) null,
  type           bigint(255) null
)
  comment '每日龙珠数据'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table day_pool_log
(
  stats_date varchar(19) charset utf8 default '' not null
  comment '统计时间',
  room_id    bigint(10)                          null,
  fee        double                              null
)
  comment '每日税收数据'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table day_summary_log
(
  stats_date    varchar(19) charset utf8 default '' not null
  comment '统计时间',
  flowing       bigint(255) default '0'             not null
  comment '金币库存',
  room1         bigint(255)                         null
  comment '新手场流水',
  room2         bigint(255)                         null
  comment '初级场流水',
  room3         bigint(255)                         null
  comment '中级场流水',
  room4         bigint(255)                         null
  comment '高级场流水',
  outdb         bigint(255)                         null
  comment '转出龙珠',
  indb          bigint(255)                         null
  comment '转入龙珠',
  outcoin       bigint(255)                         null
  comment '转出金币',
  incoin        bigint(255)                         null
  comment '转入金币',
  recharge      bigint(255)                         null
  comment '充值',
  rechargeCount bigint(255)                         null
  comment '充值笔数',
  taxroom1      double                              null
  comment '新手场抽水',
  taxroom2      double                              null
  comment '初级场抽水',
  taxroom3      double                              null
  comment '中级场抽水',
  taxroom4      double                              null
  comment '高级场抽水',
  tax           bigint(255)                         null
  comment '总抽水',
  reward        bigint(255)                         null
  comment '奖励',
  constraint stats_date
  unique (stats_date)
    comment '(null)'
)
  comment '每日数据汇总'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table day_summary_log_new
(
  stats_date    varchar(19) charset utf8 default '' not null
  comment '统计时间',
  flowing       bigint default '0'                  null
  comment '金币库存',
  room1         bigint default '0'                  null
  comment '新手场流水',
  room2         bigint default '0'                  null
  comment '初级场流水',
  room3         bigint default '0'                  null
  comment '中级场流水',
  room4         bigint default '0'                  null
  comment '高级场流水',
  outdb1        bigint default '0'                  null
  comment '转出龙珠1',
  outdb2        bigint default '0'                  null
  comment '转出龙珠2',
  outdb3        bigint default '0'                  null
  comment '转出龙珠3',
  indb1         bigint default '0'                  null
  comment '转入龙珠1',
  indb2         bigint default '0'                  null
  comment '转入龙珠2',
  indb3         bigint default '0'                  null
  comment '转入龙珠3',
  outcoin       bigint default '0'                  null
  comment '转出金币',
  incoin        bigint default '0'                  null
  comment '转入金币',
  recharge      bigint default '0'                  null
  comment '充值',
  rechargeCount bigint default '0'                  null
  comment '充值笔数',
  rechargeCoin  bigint default '0'                  null
  comment '充值金币数',
  taxroom1      double default '0'                  null
  comment '新手场抽水',
  taxroom2      double default '0'                  null
  comment '初级场抽水',
  taxroom3      double default '0'                  null
  comment '中级场抽水',
  taxroom4      double default '0'                  null
  comment '高级场抽水',
  reward        bigint default '0'                  null
  comment '奖励',
  totaltm       double default '0'                  null
  comment '玩家在线时间',
  avgtm         double default '0'                  null
  comment '玩家平均在线时间',
  newtotaltm    double default '0'                  null
  comment '当日新玩家在线时间总计',
  newavgtm      double default '0'                  null
  comment '当日新玩家平均在线时间',
  newbillave    double default '0'                  null
  comment '新用户单日平均流水',
  totalbillave  double default '0'                  null
  comment '所有用户单日平均流水',
  constraint stats_date
  unique (stats_date)
    comment '(null)'
)
  comment '每日数据汇总'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table deposit_log
(
  payid         int(20) auto_increment
  comment '充值序号'
    primary key,
  uid           bigint(10)                           not null
  comment '用户id',
  uname         varchar(255) charset utf8 default '' not null
  comment '用户名',
  pay_money     int(20)                              not null
  comment '充值金额',
  pay_type      int(10)                              not null
  comment '充值渠道类型',
  pay_type_info varchar(255) charset utf8 default '' not null
  comment '充值渠道名称',
  order_id      bigint(255)                          not null
  comment '订单id',
  add_time      datetime                             not null
  comment '充值时间',
  exchange_rate int(20)                              not null
  comment '充值比率 1元兑多少金币',
  coin_in       bigint(255)                          not null
  comment '充值金币数量',
  coin_sucess   int(10)                              not null
  comment '充值金币成功 0成功1失败',
  pay_sucess    int(10)                              not null
  comment '充值rmb成功 0成功1失败',
  back_time     int(10)                              not null
  comment '回调时间',
  coin_before   bigint(255)                          not null
  comment '充值前金币数',
  coin_after    bigint(255)                          not null
  comment '充值后金币数',
  update_flag   int(10)                              not null
  comment '更新成功状态0成功1失败',
  purchase_type int(10)                              not null,
  pay_ip        varchar(255) charset utf8 default '' not null
  comment '充值ip',
  pay_uuid      bigint(255)                          not null
  comment '充值设备码'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table dragon_transfer_log
(
  id             int(11) unsigned auto_increment
    primary key,
  sender_uid     bigint       null
  comment '赠送者',
  recv_uid       bigint       null
  comment '接收者',
  sender_nick    varchar(255) null
  comment '赠送者昵称',
  recv_nick      varchar(255) null
  comment '接受者昵称',
  transfer_type  int          null
  comment '0.珍珠，1.一星龙珠，2.二星龙珠，3.三星龙珠，10.金币，11.喇叭，12.龙珠卡',
  transfer_count int          null
  comment '赠送数量',
  tm             bigint       null
  comment '赠送时间，时间戳',
  addTime        datetime     null
);

-- auto-generated definition
create table dragoncard_use_log
(
  id         int(11) unsigned auto_increment
    primary key,
  uid        bigint   null
  comment '用户uid',
  starttime  bigint   null
  comment '龙珠卡开始时间',
  deadtime   bigint   null
  comment '龙珠卡结束时间',
  addTime    datetime null
  comment '添加时间',
  updateTime datetime null
  comment '更新时间'
)
  charset = utf8;

create index deadtime
  on dragoncard_use_log (deadtime);

create index starttime
  on dragoncard_use_log (starttime);

create index uid
  on dragoncard_use_log (uid);

-- auto-generated definition
create table exp_log
(
  uid     bigint(11) unsigned auto_increment
  comment 'user id'
    primary key,
  mid     int        null
  comment 'mission id',
  addtime datetime   null
  comment 'add time',
  exp     bigint(11) null
  comment 'exp',
  exp0    bigint(11) null
  comment 'before add',
  exp1    bigint(11) null
  comment 'after add'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table fort_cfg
(
  id     int(11) unsigned auto_increment
  comment '炮id'
    primary key,
  pkg_id int(11) unsigned       not null
  comment '炮在背包中的id',
  name   varchar(64) default '' not null
  comment '炮名称',
  type   int default '0'        not null
  comment '1 VIP等级，2 7日赠送 3 首充赠送 4 月卡赠送 5 金币购买',
  value  int default '0'        not null
  comment 'VIP等级，7日赠送，首充，月卡，炮价格'
);

create index gun_pkg_id
  on fort_cfg (pkg_id);

-- auto-generated definition
create table fort_log
(
  id          bigint unsigned auto_increment
  comment '自增id'
    primary key,
  uid         bigint                             not null
  comment '用户id',
  fort_id     int                                not null
  comment '炮台id',
  add_type    tinyint(3) default '0'             not null
  comment '添加类型，0 默认炮台 1 购买 2 成长任务赠送 3 赠送 4 vip赠送 5 周卡赠送 6 月卡赠送 7 购买礼包赠送 8 新手七天奖励 9 邮件赠送 10 抽奖活动',
  add_time    datetime default CURRENT_TIMESTAMP not null
  comment '添加时间',
  expire_time datetime default CURRENT_TIMESTAMP not null
  comment '过期时间',
  expire_utc  bigint default '0'                 not null
  comment '过期时间UTC时间，单位：毫秒'
);

-- auto-generated definition
create table fort_value_cfg
(
  fort_value int default '0' not null
  comment '炮台倍数值'
    primary key,
  diamond    int default '0' not null
  comment '炮台解锁需要的钻石数量',
  coin       int default '0' not null
  comment '炮台解锁赠送的金币数量'
);

-- auto-generated definition
create table gift_cfg
(
  max_fee  bigint(255) null
  comment '税收最多扣除',
  switch   int(10)     null
  comment '赠送开关 1代表开启；0代表关闭',
  fee_rate int         null
  comment 'fee rate',
  vip      int(10)     null
  comment 'vip限制',
  level    int         null
  comment '等级限制'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table gift_log
(
  id            int(11) unsigned auto_increment
  comment '自增id'
    primary key,
  bt            int default '0'    not null
  comment '0珍珠，1一星龙珠，2二星龙珠，3三星龙珠, 10金币， 11喇叭， 12龙珠卡',
  fid           int default '0'    not null
  comment '鱼id，如果是掉落龙珠，即掉落龙珠的鱼的id，如果是掉落珍珠，那么就是掉落珍珠的鱼的id',
  cnt           int default '0'    not null
  comment '掉落、合成、珍珠数，龙珠数；赠送龙珠数',
  uid           bigint default '0' not null
  comment '掉落、合成拥有者id，使用者id',
  ouid          bigint default '0' not null
  comment '赠送者uid，接收者uid',
  tm            datetime           null
  comment '记录添加时间',
  op            int default '0'    not null
  comment '操作类型，1掉落，2合成，3使用，4赠送，5接收',
  pearl0        int default '0'    not null
  comment '珍珠（更新前）',
  db10          int default '0'    not null
  comment '一星龙珠（更新前）',
  db20          int default '0'    not null
  comment '二星龙珠（更新前）',
  db30          int default '0'    not null
  comment '三星龙珠（更新前）',
  lb0           int default '0'    not null
  comment '喇叭（更新前）',
  dc0           int default '0'    not null
  comment '龙珠卡（更新前）',
  frozen0       int default '0'    not null
  comment '冰冻道具（更新前）',
  lock0         int default '0'    not null
  comment '锁定道具（更新前）',
  double_gold0  int default '0'    not null
  comment '双倍金币（更新前）',
  double_speed0 int default '0'    not null
  comment '双倍射速（更新前）',
  pearl1        int default '0'    not null
  comment '珍珠（更新后）',
  db11          int default '0'    not null
  comment '一星龙珠（更新后）',
  db21          int default '0'    not null
  comment '二星龙珠（更新后）',
  db31          int default '0'    not null
  comment '三星龙珠（更新后）',
  lb1           int default '0'    not null
  comment '喇叭（更新后）',
  dc1           int default '0'    not null
  comment '龙珠卡（更新后）',
  frozen1       int default '0'    not null
  comment '冰冻道具（更新后）',
  lock1         int default '0'    not null
  comment '锁定道具（更新后）',
  double_gold1  int default '0'    not null
  comment '双倍金币（更新后）',
  double_speed1 int default '0'    not null
  comment '双倍射速（更新后）'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table gift_pkg_cfg
(
  id         int(11) unsigned auto_increment
    primary key,
  wares_id   varchar(255) default '' not null
  comment '商品ID',
  gift_type  int default '0'         not null
  comment '礼物类型 0:珍珠\n1:一星龙珠\n2:二星龙珠\n3:三星龙珠\n10:金币\n11:喇叭\n12:龙珠卡\n13:钻石\n14:设置VIP\n15:冰冻道具\n16:锁定道具\n17:双倍射速道具\n18:双倍金币道具\n19:自动开炮',
  gift_count int default '0'         not null
  comment '奖励数量',
  once       tinyint default '1'     not null
  comment '发放类型，0.截止时间前每天可领取；1.一次性',
  isvalid    tinyint default '1'     not null
  comment '是否有效，0.无效；1.有效'
);

-- auto-generated definition
create table gift_pkg_log
(
  id       int(11) unsigned auto_increment
    primary key,
  uid      bigint default '0'     not null
  comment '用户id',
  wares_id varchar(32) default '' not null
  comment '商品id',
  room_ids int default '0'        not null
  comment '可以弹出的房间id，按位与的值：3表示房间1 房间2，6表示房间2 房间3',
  add_time datetime               null
  comment '添加时间',
  status   tinyint default '0'    not null
  comment '状态，0未开始，1倒计时开始，2倒计时结束，3已购买',
  room_id  int default '0'        not null
  comment '房间id',
  utc_end  bigint default '0'     not null
  comment '结束时间，utc，单位毫秒'
);

-- auto-generated definition
create table gobang_log
(
  round       bigint(11) auto_increment
  comment '游戏局数'
    primary key,
  winner      bigint(11)               null
  comment '胜者',
  loser       bigint(11)               null
  comment '负者',
  steps       varchar(11) charset utf8 null
  comment '游戏步数',
  beginTime   datetime                 null
  comment '游戏开始时间',
  finishTime  datetime                 null
  comment '游戏结束时间',
  fee         bigint(11)               null
  comment '服务费',
  guidancefee bigint(11)               null
  comment '指导费'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table growth_task_cfg
(
  id         int unsigned auto_increment
    primary key,
  type       int default '0'         not null
  comment '1 击杀n条鱼/2 人物等级到n级/3 炮台解锁至n倍/4 使用锁定道具n次/5 使用冰冻道具n次/6 累计打鱼获得n金币/7 击杀n条黄金鱼/8 击杀n条BOSS鱼/9 使用双倍金币/10 使用双倍射速',
  value      int default '0'         not null
  comment '击杀n条鱼/人物等级到n级/炮台解锁至n倍/使用道具n次/累计打鱼获得n金币/击杀n条赏金鱼/击杀n条BOSS鱼',
  room_ids   int default '15'        not null
  comment '场次按位标识 1初级场 2中级场 4高级场 8VIP场',
  gift_type  int default '0'         not null
  comment '礼物类型 0:珍珠\n1:一星龙珠\n2:二星龙珠\n3:三星龙珠\n10:金币\n11:喇叭\n12:龙珠卡\n13:钻石\n14:设置VIP\n15:冰冻道具\n16:锁定道具\n17:双倍射速道具\n18:双倍金币道具\n20:普通炮台\n21:VIP1赠送 龙息火炮\n22:VIP2赠送 蓝色狂乱\n23:VIP3赠送 邪能弩炮\n24:VIP4赠送 太阳之怒\n25:VIP5赠送 黑洞之眼\n26:VIP6赠送 宇宙湮灭\n27:七日赠送  赤色要塞\n28:首充赠送  金甲战神\n29:月卡赠送  未来之翼\n30:紫晶之辉\n31:深海巨鲨\n32:皇家礼炮\n33:天空之城\n34:凤凰之击\n35:绿色军团',
  gift_count int default '0'         not null
  comment '礼物数量：金币数量/钻石数量/?星龙珠数量/?体验炮台时长/锁定道具数量/冰冻道具数量',
  tip        varchar(128) default '' not null
  comment '任务描述，比如：捕获1条黄金鱼',
  tip1       varchar(128) default '' not null
  comment '任务房间说明，比如：初级场、中级场、高级场可完成',
  is_valid   tinyint default '1'     not null
  comment '是否有效'
);

-- auto-generated definition
create table growth_task_log
(
  id          int(11) unsigned auto_increment
    primary key,
  uid         bigint default '0'                 not null
  comment '用户id',
  task_id     int default '0'                    not null
  comment '成长任务id，growth_task_cfg表中的id对应',
  task_type   int default '1'                    not null
  comment '陈章任务类型，对应growth_task_cfg表中的type',
  value       int default '0'                    not null
  comment '成长任务值，达到growth_task_cfg中fish_value时，成长任务状态设置为1',
  state       tinyint default '0'                not null
  comment '成长任务状态，0 未完成 1已完成 2已领取',
  add_time    datetime default CURRENT_TIMESTAMP not null
  comment '添加时间',
  update_time datetime default CURRENT_TIMESTAMP not null
  comment '更新时间'
);

-- auto-generated definition
create table horselamp
(
  pid      int auto_increment
  comment '跑马灯ID'
    primary key,
  ptitle   varchar(20) charset utf8  null
  comment '跑马灯标题',
  pcontent varchar(255) charset utf8 null
  comment '跑马灯内容',
  paddtime datetime                  null
  comment '发布时间',
  pendtime datetime                  null
  comment '消息结束时间',
  ptm      int(10)                   null
  comment '轮询时间',
  isValid  int(4) default '1'        null
  comment '是否有效'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table invite_cfg
(
  id           bigint auto_increment
  comment '自增字段'
    primary key,
  inviterGift  int      not null
  comment '邀请人奖励金币',
  inviteesGift int      not null
  comment '被邀请人奖励金币',
  time         datetime not null
  comment '指定时间之后注册的用户才能填'
)
  comment '邀请活动配置';

-- auto-generated definition
create table invite_log
(
  uid  bigint                             not null
  comment '被邀请人'
    primary key,
  code bigint                             not null
  comment '邀请码/邀请人ID',
  time datetime default CURRENT_TIMESTAMP null
  comment '邀请时间'
)
  comment '邀请记录';

create index code
  on invite_log (code);

create index uid
  on invite_log (uid);

-- auto-generated definition
create table ip_rule
(
  id     int(11) unsigned auto_increment
    primary key,
  ip     varchar(32) charset utf8 default '' not null,
  islock tinyint default '1'                 not null,
  tm     datetime default CURRENT_TIMESTAMP  not null,
  constraint ip
  unique (ip)
    comment '(null)'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table level_actCfg
(
  sub_id      int(11) unsigned not null
  comment '子任务id'
    primary key,
  activity_id int              null
  comment '活动id',
  level       int              null
  comment '最低等级',
  gift_id     int              null
  comment '礼物编码，0.珍珠；1.一星龙珠;2.二星龙珠;3.三星龙珠;10.金币,11.喇叭,12.龙珠卡',
  gift_cnt    int              null
  comment '礼物数量',
  total_cnt   int              null
  comment '总的数量'
)
  comment '冲级大奖赛配置表';

-- auto-generated definition
create table level_actLog
(
  id       int(11) unsigned auto_increment
    primary key,
  uid      bigint       null,
  level    int          null,
  act_id   int          null,
  sub_id   int          null,
  dev_uuid varchar(255) null
  comment '领取设备',
  addTime  datetime     null
  comment '领取时间'
);

-- auto-generated definition
create table level_gift_cfg
(
  id         int unsigned auto_increment
    primary key,
  level      int                 not null
  comment '经验等级',
  gift_type  int default '0'     not null
  comment '礼物类型 0:珍珠\n1:一星龙珠\n2:二星龙珠\n3:三星龙珠\n10:金币\n11:喇叭\n12:龙珠卡\n13:钻石\n14:设置VIP\n15:冰冻道具\n16:锁定道具\n17:双倍射速道具\n18:双倍金币道具\n19:自动开炮，20~35炮台',
  gift_count int default '0'     not null
  comment '奖励数量，如果奖励类型是炮台，表示天数，天数为-1表示永久',
  once       tinyint default '1' not null
  comment '是否一次性，如果为0表示每天奖励'
);

-- auto-generated definition
create table login_log
(
  login_id   int(11) unsigned auto_increment
  comment '登录序号id'
    primary key,
  uid        bigint                               not null
  comment '用户id',
  io         tinyint default '1'                  not null
  comment '1登录 2登出',
  login_time datetime default CURRENT_TIMESTAMP   not null
  comment '登陆时间',
  dev_desc   varchar(128) charset utf8 default '' not null
  comment '设备描述',
  platform   int(10)                              not null
  comment '设备类型',
  dev_uuid   varchar(128) charset utf8 default '' not null
  comment '登录设备码',
  flavors    varchar(64) charset utf8             null,
  ipaddr     varchar(32) charset utf8 default ''  not null
  comment '登录ip',
  tm         datetime default CURRENT_TIMESTAMP   not null
  comment '登陆时间',
  coin       bigint default '0'                   not null,
  diamond    bigint default '0'                   not null,
  pearl      int default '0'                      not null,
  db1        int default '0'                      not null,
  db2        int default '0'                      not null,
  db3        int default '0'                      not null,
  fort_value int default '0'                      not null,
  qiye       tinyint default '1'                  not null
  comment '是否是企业包，除iOS AppStore下载的为0外，其他的都是1'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index login_time
  on login_log (login_time);

create index uid
  on login_log (uid);

-- auto-generated definition
create table logingame_stat
(
  id          int(11) unsigned auto_increment
  comment '序号'
    primary key,
  uid         bigint(10)                null
  comment '用户id',
  gameid      int(10)                   null
  comment '游戏id',
  last_time   datetime                  null
  comment '最近一次登录游戏时间',
  last_ip     varchar(255) charset utf8 null
  comment '最近一次登录ip',
  login_count int(10)                   null
  comment '进入游戏次数',
  roomid      int                       null
  comment '游戏房间id'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table lucky_cfg
(
  room_id                  int(11) unsigned auto_increment
  comment '房间id'
    primary key,
  add_perday               bigint default '0'      not null
  comment '添加add_perday指定的金币值到房间幸运奖池，一天添加一次',
  add_sum                  bigint default '0'      not null
  comment '接下来需要添加的幸运奖池总量',
  add_total                bigint default '0'      not null
  comment '历史累计添加幸运奖池总额',
  pool_floor               bigint default '10000'  not null
  comment '当当前幸运奖池低于floor值时，添加add_perday指定的金币值到房间幸运奖池',
  user_max_coin            bigint default '300000' not null
  comment '用户金币值大约此值，不能进入幸运玩家',
  limit_coin               bigint default '200000' not null
  comment '在当前房间最多可以获得的幸运奖励',
  limit_level              int default '20'        not null
  comment '房间中等级小于此值，才能成为幸运玩家',
  multiple                 int default '3'         not null,
  lucky_time_min           int default '2'         not null,
  lucky_time_max           int default '5'         not null,
  lucky_duration_min       int default '4'         not null,
  lucky_duration_max       int default '10'        null,
  rate_min                 int default '500'       not null
  comment '幸运概率最小值',
  rate_max                 int default '1000'      not null
  comment '幸运概率最大值',
  super_multiple           int default '3'         not null,
  super_lucky_time_min     int default '2'         not null,
  super_lucky_time_max     int default '5'         not null,
  super_lucky_duration_min int default '5'         not null,
  super_lucky_duration_max int default '10'        not null,
  super_rate_min           int default '1000'      not null
  comment '超级幸运概率最小值',
  super_rate_max           int default '2000'      not null
  comment '超级幸运概率最大值'
);

-- auto-generated definition
create table lucky_log
(
  id      int(11) unsigned auto_increment
    primary key,
  uid     bigint   null,
  room_id int      null,
  desk_id int      null,
  tm      datetime null,
  rewards int      null
);

-- auto-generated definition
create table lzdb_log
(
  id              int(11) unsigned auto_increment
    primary key,
  uid             bigint   null,
  personal_points bigint   null
  comment '用户自己的积分',
  all_points      bigint   null
  comment '所有玩家的积分',
  ac_id           int      null
  comment '活动id',
  sub_id          int      null
  comment '子任务id',
  addTime         datetime null
  comment '增加日志时间'
);

-- auto-generated definition
create table lzdb_points_cfg
(
  id      int(11) unsigned auto_increment
    primary key,
  p_type  int    null
  comment '得分类型，0珍珠1一星龙珠2二星龙珠3三星龙珠',
  p_count int    null
  comment '所需得分',
  r_point bigint null
  comment '奖励积分值'
);

-- auto-generated definition
create table lzdb_reward_cfg
(
  id        int(11) unsigned auto_increment
    primary key,
  r_points  bigint  null
  comment '所需积分',
  r_type    int     null
  comment '奖励类型',
  r_cnt     int     null
  comment '奖励数量',
  p_type    tinyint null
  comment '积分类型，1.个人2.所有玩家',
  relate_id int     null
  comment '对应的id'
);

-- auto-generated definition
create table mail_reward
(
  id          int(11) unsigned auto_increment
    primary key,
  mail_giftid varchar(255) null
  comment '礼物编号',
  reward_id   int          null
  comment '礼物类型编号,自定义;0.珍珠；1.db1;2.db2;3.db3;10.金币,11.喇叭,12.龙珠卡',
  reward_cnt  int          null
  comment '对应礼物数量',
  isvalid     tinyint      null
  comment '该礼物是否有效,0.无效1.有效',
  addTime     datetime     null
  comment '编辑时间',
  operator    int          null
  comment '操作者'
);

-- auto-generated definition
create table mergepercent_cfg
(
  cnt     int null,
  percent int null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table mission
(
  uid   bigint(255) not null
  comment 'user id',
  mid   int         not null
  comment 'mission id  每日任务id',
  value int         null
  comment 'mission value',
  state int         null
  comment 'incomplete complete received',
  primary key (uid, mid)
)
  engine = MyISAM
  collate = utf8_unicode_ci;

-- auto-generated definition
create table newer7_cfg
(
  day  int null,
  coin int null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table newer7_log
(
  id       int(11) unsigned auto_increment
  comment '自增id'
    primary key,
  uid      bigint                    null
  comment '用户id',
  uuid     varchar(128) charset utf8 null
  comment 'uuid',
  login_ip varchar(20) charset utf8  null
  comment '登录ip',
  days     int                       null
  comment '第几天领取',
  coin     bigint                    null
  comment '获得金币值',
  tm       datetime                  null
  comment '领取时间',
  platform int                       null
  comment '平台类型'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index tm
  on newer7_log (tm);

create index uid
  on newer7_log (uid);

-- auto-generated definition
create table notice
(
  id          int(11) unsigned auto_increment
  comment '公告编号'
    primary key,
  title       varchar(20)        null
  comment '公告标题',
  content     varchar(255)       null
  comment '公告内容',
  sender      tinyint            null
  comment '发送者，0:系统（默认）1:管理员 2:其他',
  receiver    bigint             null
  comment '接收者，0:全部 1:vip用户 2:普通用户 uid.某个用户',
  isTop       int default '0'    not null
  comment '是否置顶，0:不置顶 1:置顶',
  addtime     datetime           null
  comment '公告发布时间',
  isValid     int(4) default '1' not null
  comment '表示该公告是否有效,0.无效1.有效',
  showOrder   int default '1'    null
  comment '公告显示位置',
  ntype       int default '0'    not null
  comment '消息类型，0.查阅型消息1.邮件型消息',
  mail_giftid varchar(255)       null
  comment '如果是邮件型的消息带有礼物id',
  mail_type   int                null
  comment '0.系统发放奖励1.VIP升级奖励2.礼物赠送3.周卡4.排行奖励'
)
  engine = MyISAM;

-- auto-generated definition
create table online_count
(
  id      int(11) unsigned auto_increment
    primary key,
  tp      int      null
  comment '统计类型，1.在线人数2.游戏人数',
  p_cnt   bigint   null,
  tms     int      null
  comment '统计频率1.每分钟5.每五分钟',
  addtime datetime null
);

-- auto-generated definition
create table online_log
(
  id      bigint unsigned auto_increment
  comment '自增id'
    primary key,
  uid     bigint   null
  comment '用户id',
  tm      int      null
  comment '在线时长（分钟）',
  addtime datetime null
  comment '记录添加时间'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index addtime
  on online_log (addtime);

create index uid
  on online_log (uid);

-- auto-generated definition
create table online_reward_cfg
(
  id          int(11) unsigned auto_increment
  comment '对应online_reward_log中的reward_id'
    primary key,
  tm          int default '0'                    not null
  comment '需要的时长',
  reward      int default '0'                    not null
  comment '可以领取的奖励金币',
  valid       tinyint default '1'                not null
  comment '是否有效 0 无效，1有效',
  update_time datetime default CURRENT_TIMESTAMP not null
  comment '最后更新时间'
);

-- auto-generated definition
create table online_reward_log
(
  id          int(11) unsigned auto_increment
    primary key,
  uid         bigint default '0'                 not null
  comment '用户id',
  reward_id   int default '0'                    not null
  comment 'online_reward_cfg中的id',
  tm          int default '0'                    not null
  comment '该任务需要时间',
  reward      int default '0'                    not null
  comment '奖励金币值',
  accum_tm    int default '0'                    not null
  comment '累计时间',
  state       int default '0'                    not null
  comment '状态，0初始状态 1可领取 2 已领取',
  update_time datetime default CURRENT_TIMESTAMP not null
  comment '最后更新时间',
  insert_time datetime default CURRENT_TIMESTAMP not null
  comment '数据的插入时间'
);

create index insert_time
  on online_reward_log (insert_time);

create index reward_id
  on online_reward_log (reward_id);

create index uid
  on online_reward_log (uid);

-- auto-generated definition
create table onlinetime
(
  uid     bigint(11) null
  comment '用户id',
  gameid  int        null
  comment '游戏id 1.捕鱼2.五子棋',
  logintm datetime   null
  comment '进入游戏时间',
  leavetm datetime   null
  comment '退出游戏时间'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table order_log
(
  oid       bigint auto_increment
  comment '订单号'
    primary key,
  transid   varchar(64) charset utf8             null
  comment '计费支付平台的交易流水号',
  tradeno   varchar(64) default ''               not null
  comment '捕鱼平台生产的唯一订单号',
  platform  int default '0'                      null
  comment '平台id，0安卓爱贝支付，1ios爱贝支付，2apple支付',
  channel   int default '0'                      null
  comment '支付渠道，1苹果，2支付宝，3微信',
  uid       bigint default '0'                   null
  comment '用户id',
  waresid   varchar(64) charset utf8 default '0' null
  comment '商品id（1-6apple支付，7-12爱贝支付）',
  money     float(10, 2) default '0.00'          null
  comment '订单金额（单位：元）',
  coin      bigint default '0'                   null
  comment '订单金币值',
  diamond   bigint default '0'                   not null
  comment '购买商品获得的钻石值',
  result    int default '0'                      null
  comment 'json返回值，0成功，-1初始值，其他失败;微信支付1成功，其他失败',
  issandbox tinyint default '0'                  not null
  comment '是否是沙箱测试账号充值',
  isget     tinyint default '0'                  null
  comment '客户端是否已经获取过（用于客户端提交友盟）',
  addtime   datetime                             null
  comment '订单生成时间',
  transtime datetime                             null
  comment '订单完成时间'
)
  collate = utf8mb4_unicode_ci;

create index addtime
  on order_log (addtime);

create index addtime_idx
  on order_log (addtime)
  comment '(null)';

create index channel
  on order_log (channel);

create index isget
  on order_log (isget);

create index result
  on order_log (result);

create index uid_idx
  on order_log (uid)
  comment '(null)';

-- auto-generated definition
create table pay_exp_log
(
  id      bigint auto_increment
    primary key,
  uid     bigint                             not null
  comment '用户id',
  oid     bigint                             not null
  comment '订单号',
  paycoin int                                not null
  comment '充值对应的金币数值',
  exp     bigint                             null
  comment '用户充值的时候对应的经验值',
  time    datetime default CURRENT_TIMESTAMP null
  comment '写入时间'
)
  comment '充值的时候对应的用户经验值';

create index uid
  on pay_exp_log (uid);

-- auto-generated definition
create table pk_game_log
(
  id        int(11) unsigned auto_increment
    primary key,
  round     bigint       null
  comment '比赛编号',
  uid       bigint       null
  comment '房主id',
  game_type int          null
  comment '0.对战，1.挑战',
  gold      bigint       null
  comment '>0.胜利,<0.失败',
  nick      varchar(255) null,
  addTime   bigint       null
  comment '时间戳'
);

-- auto-generated definition
create table pool_changelog
(
  id      bigint(11) unsigned auto_increment
    primary key,
  ip      varchar(16) charset utf8 null
  comment '改动ip',
  roomid  int                      null
  comment '1.初级2.中级3.高级',
  reward  bigint                   null
  comment '改动价值，+为增加，-为减少',
  addTime datetime                 null
  comment '改动时间'
)
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table pool_log
(
  id         int(11) unsigned auto_increment
    primary key,
  room_id    int                not null
  comment '房间id',
  tm         datetime           not null
  comment '添加时间',
  win        double default '0' not null
  comment '房间总额',
  pool       double default '0' not null
  comment '房间小奖池',
  fee        double default '0' not null
  comment '手续费总额',
  dragonball double default '0' not null
  comment '龙珠奖池总额',
  lucky      double default '0' not null
  comment '幸运奖池总额'
)
  charset = utf8;

create index room_id
  on pool_log (room_id);

create index tm
  on pool_log (tm);

-- auto-generated definition
create table portal_activity
(
  id         int auto_increment
  comment '活动ID'
    primary key,
  createTime timestamp default CURRENT_TIMESTAMP  not null
  comment '创建时间',
  isDeleted  varchar(1) charset utf8 default 'N'  not null
  comment '是否删除',
  isEnabled  varchar(1) charset utf8 default 'Y'  not null
  comment '是否启用',
  name       varchar(50) charset utf8 default ''  not null
  comment '活动名称',
  image      varchar(300) charset utf8 default '' not null
  comment '活动图片',
  url        varchar(300) charset utf8            null
  comment '活动链接',
  startTime  datetime                             null
  comment '开始时间',
  endTime    datetime                             null
  comment '结束时间',
  sortNo     int default '0'                      not null
  comment '排序号'
)
  comment '门户活动'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table portal_download
(
  id         int auto_increment
  comment '下载ID'
    primary key,
  createTime timestamp default CURRENT_TIMESTAMP  not null
  comment '创建时间',
  isDeleted  varchar(1) charset utf8 default 'N'  not null
  comment '是否删除',
  isEnabled  varchar(1) charset utf8 default 'Y'  not null
  comment '是否启用',
  name       varchar(50) charset utf8 default ''  not null
  comment '下载名称',
  image      varchar(300) charset utf8 default '' not null
  comment '下载图片',
  url        varchar(300) charset utf8            null
  comment '下载链接',
  startTime  datetime                             null
  comment '开始时间',
  endTime    datetime                             null
  comment '结束时间',
  sortNo     int default '0'                      not null
  comment '排序号',
  type       tinyint(3) default '0'               not null
  comment '类型：0公众号、1官网'
)
  comment '门户下载'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table portal_login
(
  id         bigint auto_increment
  comment 'ID'
    primary key,
  createTime timestamp default CURRENT_TIMESTAMP not null
  comment '创建时间',
  uid        bigint default '0'                  not null
  comment '用户ID',
  token      varchar(50) charset utf8 default '' not null
  comment '令牌',
  systemName varchar(50) charset utf8            null
  comment '系统名称',
  ip         varchar(50) charset utf8            null
  comment 'IP'
)
  comment '门户登录'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table portal_pay_order
(
  orderId    bigint auto_increment
  comment '订单ID'
    primary key,
  createTime timestamp default CURRENT_TIMESTAMP not null
  comment '创建时间',
  userId     bigint default '0'                  not null
  comment '用户ID',
  orderNo    varchar(50) charset utf8 default '' not null
  comment '订单编号',
  price      float(10, 2) default '0.00'         not null
  comment '价格',
  status     tinyint(1) default '0'              not null
  comment '状态：0默认、1支付成功',
  payType    tinyint(2) default '0'              not null
  comment '支付类型',
  clientIP   varchar(15) charset utf8 default '' not null
  comment '客户端IP',
  money      float(10, 2) default '0.00'         not null
  comment '金额',
  payTime    timestamp                           null
  comment '支付时间',
  payOrderNo varchar(50) charset utf8            null
  comment '支付订单编号',
  payInfo    text charset utf8                   null
  comment '支付信息',
  waresId    varchar(255)                        null
  comment '商品id',
  complete   tinyint default '0'                 not null
  comment '是否已经完成0:未完成 1:已完成',
  constraint orderNo
  unique (orderNo)
    comment '(null)'
)
  comment '门户支付订单'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table portal_sms
(
  id         bigint auto_increment
  comment 'ID'
    primary key,
  createTime timestamp default CURRENT_TIMESTAMP  not null
  comment '创建时间',
  type       tinyint(3) default '0'               not null
  comment '类型：0手机绑定',
  phone      varchar(11) charset utf8 default ''  not null
  comment '手机',
  content    varchar(300) charset utf8 default '' not null
  comment '内容'
)
  comment '门户短信'
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table props_cfg
(
  id            int(11) unsigned auto_increment
    primary key,
  type          int default '0' not null
  comment '道具类型id，12:龙珠卡\n15:冰冻道具\n16:锁定道具\n17:双倍射速道具\n18:双倍金币道具',
  coin          int default '0' not null
  comment '道具价格，金币',
  diamond       int default '0' not null
  comment '道具价格，钻石',
  duration      int default '0' not null
  comment '道具时长',
  times_per_day int default '0' not null
  comment '每天购买次数，0表示不限次数'
);

-- auto-generated definition
create table rank_rewards_cfg
(
  num     int(11) unsigned auto_increment
  comment '排名'
    primary key,
  rewards int default '0' not null
  comment '奖励金币'
)
  charset = utf8;

-- auto-generated definition
create table region
(
  id         bigint unsigned        not null
  comment 'ip地址对应的整形值'
    primary key,
  ip         varchar(32) default '' not null,
  country    varchar(64) default '' not null,
  area       varchar(64) default '' not null,
  region     varchar(64) default '' not null,
  city       varchar(64) default '' not null,
  county     varchar(64) default '' not null,
  isp        varchar(64) default '' not null,
  country_id varchar(32) default '' not null,
  area_id    varchar(32) default '' not null,
  region_id  varchar(32) default '' not null,
  city_id    varchar(32) default '' not null,
  county_id  varchar(64) default '' not null,
  isp_id     varchar(32) default '' not null
);

-- auto-generated definition
create table reward_rank
(
  id        int(11) unsigned auto_increment
    primary key,
  ordershow int          null
  comment '排序',
  uid       bigint       null,
  nick_name varchar(255) null,
  showcoin  bigint       null
  comment '显示金币值',
  avatar    varchar(255) null
);

-- auto-generated definition
create table room_lucky_log
(
  id       int(11) unsigned auto_increment
    primary key,
  room_id  int      null
  comment '房间id',
  add_coin bigint   null
  comment '添加幸运奖池金币值',
  tm       datetime null
  comment '添加幸运奖池时间'
);

-- auto-generated definition
create table stats_coin
(
  stats_time    varchar(19) charset utf8 default '' not null
  comment '统计时间',
  total_coin    bigint default '0'                  not null
  comment '金币库存',
  pearl         bigint                              null
  comment '珍珠库存',
  db1           bigint                              null
  comment '一星龙珠',
  db2           bigint                              null
  comment '二星龙珠',
  db3           bigint                              null
  comment '三星龙珠',
  dbcoin        varchar(255)                        null
  comment '龙珠转化金币',
  yulecoin      bigint                              null
  comment '鱼乐金币库存',
  stock_large   bigint                              null
  comment '大奖池',
  stock_normal  bigint                              null
  comment '小奖池',
  bibei_large   bigint                              null
  comment '大比倍奖池',
  bibei_normal  bigint                              null
  comment '小比倍奖池',
  fee_large     bigint                              null
  comment '大奖池对应的手续费',
  fee_normal    bigint                              null
  comment '小奖池对应的手续费',
  buyufee       double                              null
  comment '捕鱼手续费',
  recharge_coin bigint                              null
  comment '充值金币',
  reward_coin   bigint                              null
  comment '奖励金币',
  win1          double                              null,
  pool1         double                              null,
  fee1          double                              null,
  dragonball1   double                              null,
  win2          double                              null,
  pool2         double                              null,
  fee2          double                              null,
  dragonball2   double                              null,
  win3          double                              null,
  pool3         double                              null,
  fee3          double                              null,
  dragonball3   double                              null,
  win4          double                              null,
  pool4         double                              null,
  fee4          double                              null,
  dragonball4   double                              null,
  constraint stats_time
  unique (stats_time)
    comment '(null)'
)
  comment '统计流水'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table stats_data_hour
(
  statsDate   varchar(255) charset utf8 default '' not null
  comment '统计日期',
  statsHour   varchar(255) default ''              not null,
  win1        bigint default '0'                   not null
  comment '新手场房间总额',
  win2        bigint default '0'                   not null
  comment '初级场房间总额',
  win3        bigint default '0'                   not null
  comment '中级场房间总额',
  win4        bigint default '0'                   not null
  comment '高级场房间总额',
  dragonball1 bigint default '0'                   not null
  comment '新手场龙珠奖池总额',
  dragonball2 bigint default '0'                   not null
  comment '初级场龙珠奖池总额',
  dragonball3 bigint default '0'                   not null
  comment '中级场龙珠奖池总额',
  dragonball4 bigint default '0'                   not null
  comment '高级场龙珠奖池总额'
)
  comment '每小时龙珠奖池统计数据'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table stats_remain
(
  dru                 bigint(10)             null
  comment '新增用户',
  stat_time           varchar(19) default '' not null
  comment '统计时间',
  add_time            varchar(19) default '' not null
  comment '添加时间',
  second_day          varchar(255)           null
  comment '第二日留存',
  third_day           varchar(255)           null
  comment '第3日留存',
  fourth_day          varchar(255)           null
  comment '第4日留存',
  seventh_day         varchar(255)           null
  comment '第7日留存',
  eighth_day          varchar(255)           null
  comment '第8日留存',
  fourteen_day        varchar(255)           null
  comment '第14日留存',
  fifteenth_day       varchar(255)           null
  comment '第15日留存',
  thirtieth_day       varchar(255)           null
  comment '第30日留存',
  thirtieth_first_day varchar(255)           null
  comment '第31日留存',
  constraint stats_date
  unique (stat_time)
    comment '(null)'
)
  comment '用户留存表'
  charset = utf8;

-- auto-generated definition
create table stats_summary
(
  stats_date         varchar(10) charset utf8 default '' not null
  comment '统计日期',
  reg_user_count     int default '0'                     not null
  comment '注册人数',
  login_user_count   int default '0'                     not null
  comment '登录人数',
  max_online         int default '0'                     not null
  comment '最高在线',
  avg_online         int default '0'                     not null
  comment '平均在线',
  deposit_user_count int default '0'                     not null
  comment '充值人数',
  depost_count       int default '0'                     not null
  comment '充值次数',
  pay_money          int default '0'                     not null
  comment '充值金额',
  total_coin         bigint default '0'                  not null
  comment '金币库存',
  total_exp          bigint default '0'                  not null
  comment '今日流水',
  constraint stats_date
  unique (stats_date)
    comment '(null)'
)
  comment '统计汇总'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table stats_user_coin
(
  stats_date datetime           null
  comment '统计日期',
  uid        bigint(10)         not null
  comment '用户id',
  coin       bigint default '0' not null
)
  comment '统计用户流水'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table sub_act_cfg
(
  activity_id int default '0'          not null
  comment '活动id',
  sub_id      int(11) unsigned         not null
  comment '子任务id',
  level       int default '0'          not null
  comment '最低等级',
  total_cnt   int default '0'          not null
  comment '总的数量',
  sub_title   varchar(128) default ''  not null
  comment '子任务标题；仅用于 更新公告',
  sub_content varchar(8192) default '' not null
  comment '子任务详细内容；仅用于 更新公告',
  primary key (sub_id, activity_id)
)
  comment '冲级大奖赛配置表';

-- auto-generated definition
create table sub_act_gift
(
  sub_id      int default '0' not null
  comment '子任务id',
  activity_id int default '0' not null
  comment '活动id',
  gift_type   int default '0' not null
  comment '礼物编码，0.珍珠；1.一星龙珠;2.二星龙珠;3.三星龙珠;10.金币,11.喇叭,12.龙珠卡',
  gift_count  int default '0' not null
  comment '礼物数量',
  primary key (sub_id, activity_id, gift_type)
);

-- auto-generated definition
create table sub_act_log
(
  id       int(11) unsigned auto_increment
    primary key,
  uid      bigint       null,
  level    int          null,
  act_id   int          null,
  sub_id   int          null,
  dev_uuid varchar(255) null
  comment '领取设备',
  addTime  datetime     null
  comment '领取时间'
);

-- auto-generated definition
create table tb_bonus
(
  id    int(11) unsigned auto_increment
    primary key,
  bonus int default '0' null,
  fee   int default '0' null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table transfer_cfg
(
  id            int(11) unsigned auto_increment
    primary key,
  limit_type    tinyint null
  comment '限制类型0.level,1.vip,2.money,3.level和vip组合...',
  vip           int     null
  comment 'vip限制',
  level         int     null
  comment '等级限制',
  money         bigint  null
  comment '充值限制',
  maxcoin_level bigint  null
  comment '针对level金币限制',
  maxcoin_vip   bigint  null
  comment '针对VIP金币限制',
  maxcoin_money bigint  null
  comment '针对充值金币限制'
);

-- auto-generated definition
create table user
(
  uid          bigint(10)                            not null
  comment '用户id'
    primary key,
  upwd         varchar(255)                          null
  comment '用户密码',
  uname        varchar(64)                           null
  comment '用户名即账号',
  nick_name    varchar(255)                          null
  comment '用户昵称',
  avatar       varchar(255) charset utf8 default '0' null
  comment '用户头像url',
  sex          int(10)                               null
  comment '用户性别',
  isrobot      int                                   null
  comment '1机器人 0用户',
  ufreeze      int(10)                               null
  comment '用户封停0正常1封停',
  plat_src     int(10)                               null
  comment '用户来源平台',
  reg_tm       datetime                              null
  comment '注册时间',
  login_tm     datetime                              null
  comment '最近登录时间',
  logout_tm    datetime                              null
  comment '最近登出时间',
  reg_ip       varchar(255) charset utf8             null
  comment '注册ip',
  login_ip     varchar(255) charset utf8             null
  comment '登录ip',
  dev_desc     varchar(255) charset utf8             null
  comment '设备描述',
  platform     varchar(255) charset utf8             null
  comment '设备类型 3安卓 4,5苹果 0其它',
  dev_uuid     varchar(255) charset utf8             null
  comment '设备uuid',
  login_uuid   varchar(255)                          null,
  flavors      varchar(255) charset utf8             null
  comment '渠道商',
  ID_number    varchar(20) charset utf8              null
  comment '身份证号码',
  ID_valid     int(10)                               null
  comment '身份验证是否通过',
  phone        varchar(11) charset utf8              null
  comment '手机号码',
  phone_valid  int(10)                               null
  comment '手机号码验证通过',
  wx_openid    varchar(255) charset utf8             null
  comment '微信openid',
  find_pwd_que varchar(255) charset utf8             null
  comment '密保问题',
  find_pwd_anw varchar(255) charset utf8             null
  comment '密保答案',
  real_name    varchar(32) charset utf8              null
  comment '真实姓名',
  token        varchar(255) charset utf8             null
  comment '登录token',
  weapon       int default '0'                       null
  comment '炮类型',
  ver          varchar(16) charset utf8              null
  comment '版本',
  qiye         tinyint default '1'                   not null
  comment '是否是企业包，除iOS AppStore下载的为0外，其他的都是1',
  type         tinyint default '0'                   null
  comment '用户类型，0.普通用户1.普通推广员2.金牌推广员3.机器人',
  bind_uid     bigint                                null
  comment '普推对应金推uid',
  bind_date    datetime                              null
  comment '绑定时间',
  isnewReg     tinyint                               null
  comment '是否新注册用户，1是0不是',
  wx_unionid   varchar(255)                          null
  comment '微信公共id',
  device_token varchar(255)                          null
  comment '友盟推送用',
  freeze       tinyint default '0'                   not null
  comment '是否冻结账号',
  reg_tm_utc   bigint default '0'                    not null
  comment '注册时间，utc时间'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

create index login_tm
  on user (login_tm);

create index reg_tm
  on user (reg_tm);

-- auto-generated definition
create table user_exp
(
  uid         bigint unsigned                     not null
  comment '用户id'
    primary key,
  exp         double default '0'                  not null
  comment '经验值',
  update_time timestamp default CURRENT_TIMESTAMP not null
  on update CURRENT_TIMESTAMP,
  exp2        bigint                              null
  comment '用户今日以前的输赢值'
);

-- auto-generated definition
create table user_stat
(
  uid            bigint unsigned             not null
  comment '用户id'
    primary key,
  nick_name      varchar(255)                null
  comment '用户昵称',
  level          int(10) default '0'         null
  comment '用户等级',
  vip            int(10) default '0'         null
  comment 'vip等级',
  log_cnt        int(20) default '0'         null
  comment '用户登录次数',
  online_time    bigint default '1'          null
  comment '用户在线总时长',
  game_cnt       int(20) default '0'         null
  comment '用户游戏次数',
  game_time      bigint default '0'          null
  comment '用户游戏时长',
  money          double default '0'          null
  comment '用户充值（rmb）',
  coin           bigint default '0'          null
  comment '用户当前金币',
  agent_coin     bigint default '0'          not null
  comment '代理充值获得的金币',
  get_coin       bigint default '0'          null
  comment '用户获得金币总额',
  pay_coin       bigint default '0'          null
  comment '用户充值获得金币',
  award_coin     bigint default '0'          null
  comment '用户系统奖励金币',
  giftin_coin    bigint default '0'          null
  comment '礼物接收金币',
  giftout_coin   bigint default '0'          null
  comment '礼物赠送金币',
  pearl          int default '0'             null
  comment '珍珠数量',
  db1            int default '0'             null
  comment '一星龙珠数量',
  db2            int default '0'             null
  comment '二星龙珠数量',
  db3            int default '0'             null
  comment '三星龙珠数量',
  lb             int default '0'             null
  comment '喇叭数量',
  user_win       bigint default '0'          null
  comment '用户输赢金币',
  winstat        int(20) default '0'         null
  comment '输赢状态 0平，1赢，2输',
  exp            bigint default '0'          null
  comment '经验值',
  act            int default '0'             null
  comment '活跃度',
  dc             int default '0'             null
  comment '龙珠卡数量',
  froze          int default '0'             not null
  comment '冰冻道具',
  `lock`         int default '0'             not null
  comment '锁定道具',
  double_gold    int default '0'             not null
  comment '双倍金币',
  double_speed   int default '0'             not null
  comment '双倍射速',
  auto_fire      bigint default '0'          not null
  comment '自动开炮功能截止utc时间，单位毫秒',
  lb_points      bigint default '0'          null
  comment '购买喇叭获取的积分值',
  lz_points      bigint default '0'          null
  comment '捕获龙珠获得的积分值',
  lucky_flag     tinyint default '0'         not null
  comment '是否超级幸运玩家',
  lucky_room1    bigint default '0'          not null
  comment '新手场获得的幸运奖励金币总数',
  lucky_room2    bigint default '0'          not null
  comment '初级场获得的幸运奖励金币总数',
  add_rate       float(10, 2) default '0.00' not null
  comment '增加与的死亡概率',
  multi_rate     float(10, 2) default '1.00' not null
  comment '与死亡概率系数',
  coin_limit     bigint default '0'          not null
  comment '个人总值上/下限',
  transfer_coin  bigint default '0'          not null
  comment '当天金币转移值',
  isgetvipreward int default '2'             not null
  comment '是否获取vip奖励，0.不能获取，1.未获取，2.已获取',
  first_pay      int default '0'             not null
  comment '是否第一次充值 1、5元 2、10元。。。',
  diamond        bigint default '0'          not null
  comment '用户钻石数量',
  pay_diamond    bigint default '0'          not null
  comment '用户充值获得钻石',
  award_diamond  bigint default '0'          not null
  comment '用户系统奖励钻石',
  diamond_used   bigint default '0'          not null
  comment '用户已使用的钻石总数',
  fort_value     int default '100'           not null
  comment '当前解锁的炮倍值',
  daily_reward   int default '0'             not null
  comment '当前解锁的炮倍值',
  daily_giftpkg  int default '0'             not null
  comment '今日是否显示充值礼包，按位标识不同礼包是否显示，1：16号礼包；2：17号礼包；。。。'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table userlevel_cfg
(
  level int(11) unsigned default '0' not null
  comment '等级'
    primary key,
  exp   bigint(11)                   null
  comment '达到该等级所需的经验值'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table uservip_cfg
(
  vip      int unsigned default '0' not null
  comment '对应vip等级'
    primary key,
  pay_coin bigint                   null
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table uuid_rule
(
  id     int(11) unsigned auto_increment
    primary key,
  uuid   varchar(255) charset utf8 null,
  islock int(10)                   null,
  constraint uuid
  unique (uuid)
    comment '(null)'
)
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table vip_gift_cfg
(
  id         int unsigned auto_increment
    primary key,
  vip        int                 not null
  comment 'vip等级',
  gift_type  int default '0'     not null
  comment '奖励类型：0.珍珠；1.db1;2.db2;3.db3;10.金币,11.喇叭,12.龙珠卡,20~35炮台',
  gift_count int default '0'     not null
  comment '奖励数量',
  once       tinyint default '0' not null
  comment '是否一次性，如果为0表示每天奖励'
);

-- auto-generated definition
create table wares_cfg
(
  wares_id             varchar(32) default '' not null
  comment '商品id'
    primary key,
  wares_type           tinyint default '1'    not null
  comment '商品类型，1购买金币，2购买钻石，4周卡/月卡/新手/炮倍礼包，5限购礼包',
  first_pay            int default '0'        not null
  comment '首充标志，user_stat表中first_pay字段为按位与的一个值',
  mail_giftid          varchar(16) default '' not null
  comment '对应mail_reward表中的mail_giftid字段',
  price                int default '0'        not null
  comment '商品价格，单位元',
  discount             int default '0'        not null
  comment '折扣价格',
  coin                 int default '0'        not null
  comment '购买商品获得多少金币',
  first_reward_coin    int default '0'        not null
  comment '首次购买商品奖励多少金币',
  diamond              int default '0'        not null
  comment '购买商品获得多少钻石',
  first_reward_diamond int default '0'        not null
  comment '首次购买商品奖励多少钻石',
  room_ids             int default '0'        not null
  comment '在哪些房间有效，按位与',
  countdown            int default '0'        not null
  comment '显示限购礼包的倒计时时间'
);

-- auto-generated definition
create table web_config
(
  id        int(11) unsigned auto_increment
    primary key,
  xchg_rate int null
  comment '兑换比例，1元对应？金币'
)
  collate = utf8mb4_unicode_ci;

-- auto-generated definition
create table weekly_card_cfg
(
  id         int(11) unsigned auto_increment
    primary key,
  card_id    varchar(255)        null
  comment '周卡商品id',
  reward_id  int                 null
  comment '奖励编号，0.珍珠；1.一星龙珠;2.二星龙珠;3.三星龙珠;10.金币,11.喇叭,12.龙珠卡',
  reward_cnt int                 null
  comment '奖励数量',
  send_type  tinyint default '1' not null
  comment '发放类型，0.直接发放；1.邮件发放',
  isvalid    tinyint default '1' not null
  comment '是否有效，0.无效；1.有效'
);

-- auto-generated definition
create table weekly_card_log
(
  id       bigint unsigned auto_increment
    primary key,
  oid      bigint              null
  comment '订单id',
  uid      bigint              null
  comment '持有人id',
  card_id  varchar(255)        null
  comment '周卡编号，默认8',
  state    tinyint default '0' not null
  comment '是否有效，0.表示无效1.表示失效',
  start_tm bigint              null
  comment '生效时间，utc',
  end_tm   bigint              null
  comment '失效时间,utc',
  addTime  datetime            null
  comment '购买时间'
);

-- auto-generated definition
create table wg_buylb_cfg
(
  id     int(11) unsigned auto_increment
    primary key,
  price  bigint null
  comment '价格',
  points bigint null
  comment '可得积分'
);

-- auto-generated definition
create table wg_buylb_log
(
  id                  int(11) unsigned auto_increment
    primary key,
  oid                 varchar(255) default '' not null
  comment '交易编码',
  uid                 bigint                  null,
  usercoin_before     bigint                  null
  comment '购买前用户金币',
  usercoin_after      bigint                  null
  comment '购买后用户金币',
  user_lbpoint_before bigint                  null,
  user_lbpoints_after bigint                  null,
  cost_coin           bigint                  null
  comment '本次消耗金币值',
  cnt                 bigint                  null
  comment '喇叭数量',
  lbpoints            bigint                  null
  comment '获得积分',
  code                int                     null
  comment '订单状态，0成功，其他值参考enum的ret值',
  addTime             datetime                null,
  updateTime          datetime                null
);

-- auto-generated definition
create table wg_cost_cfg
(
  id       int(11) unsigned auto_increment
    primary key,
  cost     bigint  null
  comment '花费积分值',
  cnt      int     null
  comment '游戏次数',
  gametype tinyint null
  comment '1表示1次10表示10次'
);

-- auto-generated definition
create table wg_game_log
(
  id      int(11) unsigned auto_increment
    primary key,
  round   bigint       null
  comment '游戏局数',
  uid     bigint       null
  comment '用户的记录',
  records varchar(255) null
  comment '游戏记录描述',
  cnt     int          null
  comment '游戏次数',
  addTime datetime     null
);

-- auto-generated definition
create table wg_records
(
  id        bigint unsigned auto_increment
    primary key,
  round     bigint   not null
  comment '局数序号',
  uid       bigint   null,
  cost      bigint   null
  comment '本次消耗积分值',
  game_type int      null
  comment '次数,1表示每局1次，10表示每局10次',
  reward_id int      null
  comment '礼物编号，具体信息查看wg_rewardsInfo',
  addTime   datetime null
  comment '添加时间'
);

-- auto-generated definition
create table wg_rewardsInfo
(
  id                       int(11) unsigned auto_increment
    primary key,
  reward_name              varchar(64) null
  comment '奖品名称',
  reward_cnt               int         null
  comment '奖品数量',
  reward_pos               int         null
  comment '奖品所在位置',
  reward_type              int         null
  comment '奖品类型',
  reward_probability       int         null
  comment '奖品被抽中概率',
  reward_probability_fasce int         null
);

-- auto-generated definition
create table wg_usr_cfg
(
  id    int(11) unsigned auto_increment
    primary key,
  vip   int     null
  comment 'vip等级',
  count int     null
  comment '次数',
  tp    tinyint null
);

-- auto-generated definition
create table ylc_cfg
(
  id         bigint auto_increment
  comment '自增字段'
    primary key,
  gameId     int null
  comment '游戏ID 1捕鱼 2英雄传 3渔乐',
  vipLimit   int null
  comment 'vip等级限制',
  levelLimit int null
  comment '等级限制'
)
  comment '游戏进入的配置信息';

-- auto-generated definition
create table yule_betlog
(
  id      int(11) unsigned auto_increment
    primary key,
  type    tinyint default '0' not null
  comment '3 机器人 0普通用户',
  round   bigint              null,
  uid     bigint              null,
  betcoin bigint              null,
  pos     int                 null,
  addTime varchar(255)        null
);

-- auto-generated definition
create table yule_gamelog
(
  round        bigint unsigned auto_increment
  comment '游戏期数'
    primary key,
  playercount  int          null
  comment '本期参与游戏人数',
  totalbet     bigint       null
  comment '本期总下分值',
  totalbetInfo varchar(255) null
  comment '本期下分明细',
  totalreward  bigint       null
  comment '本期发放奖励总额',
  changeCoin   bigint       null
  comment '本期变化值',
  totalsum     bigint       null
  comment '累计值',
  pos          int          null
  comment '本期开奖位：0.大金龙1.金龟2.金鲨3.灯笼鱼4.海草鱼5.小丑鱼6.炸弹',
  ongameplayer int          null
  comment '当前在游戏中的玩家统计',
  fee          bigint       null,
  addtime      datetime     null
);

-- auto-generated definition
create table yule_playerlog
(
  round       bigint unsigned     not null
  comment '游戏期数',
  uid         bigint              not null,
  type        tinyint default '0' not null
  comment '3 机器人 0普通用户',
  nick        varchar(255)        null,
  betInfo     varchar(255)        null
  comment '每个位置的下分明细',
  rewardpos   int                 null
  comment '本局获奖位置',
  betsum      bigint              null
  comment '本局下分总值',
  winsum      bigint              null
  comment '本局获奖总值',
  changeCoin  bigint              null
  comment '本期金币变化值',
  totalchange bigint              null
  comment '当前累计变化值',
  fee         bigint              null,
  addTime     datetime            null,
  primary key (round, uid)
);

create index addTime
  on yule_playerlog (addTime);

create index winsum
  on yule_playerlog (winsum);

-- auto-generated definition
create table yule_resultlog
(
  round   bigint unsigned not null,
  uid     bigint          null,
  winsum  bigint          null
  comment '上期获奖总值',
  pos     int             null
  comment '上期开奖位置',
  addTime datetime        null
)
  comment '存放用户中途退出时的结果数据';

-- auto-generated definition
create table yule_summary
(
  stats_date varchar(19) charset utf8 default '' not null
  comment '统计时间'
    primary key,
  coin       varchar(255)                        null
)
  comment '每日鱼乐数据'
  engine = MyISAM
  collate = utf8mb4_unicode_ci;

