-- auto-generated definition
create table bibeilog
(
  round        bigint default '0'                 not null
  comment '局数',
  bibei_round  int default '0'                    not null
  comment '比倍的次数',
  uid          bigint default '0'                 not null
  comment '用户id',
  bibei_pos    int default '0'                    not null
  comment '用户猜的大小和',
  bibei_score  bigint default '0'                 not null
  comment '投入分',
  bibei_result varchar(11) default ''             not null
  comment '比倍结果，两个点数及大小和结果',
  bibei_reward bigint default '0'                 not null
  comment '比倍获得的奖励',
  add_time     datetime default CURRENT_TIMESTAMP not null
  comment '添加记录时间',
  primary key (round, bibei_round)
)
  charset = utf8mb4;

-- auto-generated definition
create table bonuslog
(
  id          bigint unsigned auto_increment
    primary key,
  round       bigint      null
  comment '时间戳对应的显示标志',
  bonus_round int         null
  comment '小玛丽累计次数',
  uid         bigint      null
  comment '玩家的id',
  img_list    varchar(11) null
  comment '中间四张图片信息',
  pos         int         null
  comment '停止位置',
  pos_img     int         null
  comment '停止位置的图片',
  bonus_score bigint      null
  comment '投入的金币值',
  reward      bigint      null
  comment '本局结束后产生的分数',
  add_time    datetime    null
  comment '添加时间'
)
  charset = utf8mb4;

-- auto-generated definition
create table coinlog
(
  id          varchar(255) not null
  comment '订单id'
    primary key,
  gid         bigint(255)  null
  comment '用户id',
  round_store varchar(20)  null
  comment '游戏局数',
  addGameCoin bigint       null
  comment '金币变化值',
  fee         bigint       null
  comment '台费',
  chgType     int          null
  comment '金币变化类型，1.兑换2.slot奖励3.slot分数4.比倍环节5.bonus（小玛丽）环节6总得分',
  appCoin0    bigint       null
  comment '变化前平台币',
  appCoin1    bigint       null
  comment '变化后平台币',
  gameCoin0   bigint       null
  comment '变化前游戏币',
  gameCoin1   bigint       null
  comment '变化后游戏币',
  addTime     datetime     null
  comment '记录添加时间',
  platform    int          null
  comment '来源平台',
  pid         varchar(255) null
  comment '用户在平台的id',
  round_show  varchar(11)  null
  comment '时间戳'
)
  charset = utf8mb4;

create index addtime_idx
  on coinlog (addTime);

create index chgtype_idx
  on coinlog (chgType);

create index round_idx
  on coinlog (round_store);

create index uid_idx
  on coinlog (gid);

  -- auto-generated definition
create table full_rate
(
  id         int(11) unsigned       not null
    primary key,
  full_name  varchar(32) default '' not null,
  rate_value int default '0'        not null,
  pic_9_num  int default '0'        not null
);


-- auto-generated definition
create table line_rate
(
  pic_id         tinyint default '1'    not null
    primary key,
  pic_name       varchar(64) default '' not null,
  three_multiple int default '0'        not null,
  four_multiple  int default '0'        not null,
  five_multiple  int default '0'        not null,
  full_multiple  int default '0'        not null,
  mary_multiple  int default '0'        not null
);


-- auto-generated definition
create table mary_rate
(
  id          int(11) unsigned auto_increment
    primary key,
  pic_id      tinyint default '0'    not null,
  pic_name    varchar(32) default '' not null,
  rate_value1 int default '0'        not null,
  rate_value2 int default '0'        not null,
  rate_value3 int default '0'        not null,
  rate_value5 int default '0'        not null
);

-- auto-generated definition
create table playerlog
(
  round        bigint                             not null
  comment '游戏id'
    primary key,
  uid          bigint default '0'                 not null,
  platform     varchar(11) default '0'            not null,
  show_round   varchar(32) default ''             not null
  comment '游戏局数',
  cell_score   bigint(11) default '0'             not null
  comment '单线分数',
  line_cnt     int default '0'                    not null
  comment '压线条数',
  total_score  bigint default '0'                 not null
  comment '总下分值',
  stock_type   int default '1'                    not null
  comment '1 StockLarge 2 StockNormal',
  reward_type  int default '0'                    not null
  comment '0.Normalpool1.Bigpool2.Largepool',
  slot_imglist varchar(128) default ''            not null
  comment 'slot图形',
  slot_result  varchar(512) default ''            not null
  comment 'slot结果[成线条数，总奖励，[线码，成图，个数，[位置]，奖励]]',
  slot_reward  bigint default '0'                 not null
  comment '总奖励值',
  bonus_count  int default '0'                    not null
  comment '小玛丽的次数',
  bonus_reward bigint default '0'                 not null
  comment '小玛丽的分数',
  bibei_count  int(20) default '0'                not null
  comment '单局游戏中比倍的次数',
  bibei_reward bigint(11) default '0'             not null
  comment '比倍获得的分数',
  user_fee     bigint default '0'                 not null
  comment '玩家支付的手续费',
  total_reward bigint default '0'                 not null
  comment '最终获得的总分值',
  add_time     datetime default CURRENT_TIMESTAMP not null
  comment '新增该记录的时间'
)
  charset = utf8mb4;

create index add_time
  on playerlog (add_time);

create index platform
  on playerlog (platform);

create index uid
  on playerlog (uid);


-- auto-generated definition
create table stock
(
  id           bigint unsigned auto_increment
    primary key,
  stock_large  bigint default '0'                 not null
  comment '大奖池',
  stock_normal bigint default '0'                 not null
  comment '小奖池',
  bibei_large  bigint default '0'                 not null
  comment '大比倍奖池',
  bibei_normal bigint default '0'                 not null
  comment '小比倍奖池',
  fee_large    bigint default '0'                 not null
  comment '大奖池对应的手续费',
  fee_normal   bigint default '0'                 not null
  comment '小奖池对应的手续费',
  mvcnt_large  bigint default '0'                 not null
  comment '当天从Bibei奖池移动到Stock中的次数（大奖池）',
  mvcnt_normal bigint default '0'                 not null
  comment '当天从Bibei奖池移动到Stock中的次数（小奖池）',
  addTime      datetime default CURRENT_TIMESTAMP not null
  comment '添加时间',
  updateTime   datetime default CURRENT_TIMESTAMP not null
  comment '更新时间'
)
  charset = utf8mb4;

create index addTime
  on stock (addTime);

create index updateTime
  on stock (updateTime);

-- auto-generated definition
create table stock_cfg
(
  id           int(11) unsigned auto_increment
    primary key,
  stock_type   int default '0'    not null
  comment '库存类型，StockS、StockSS',
  stock_id     int default '0'    not null
  comment '库存标识',
  stock_silver bigint default '0' not null
  comment '最低库存量',
  stock_rate   int default '0'    not null
  comment '库存率(百分比，97即0.97)',
  pic0         int default '0'    not null
  comment '水浒概率',
  pic1         int default '0'    not null
  comment '忠义概率',
  pic2         int default '0'    not null
  comment '替天概率',
  pic3         int default '0'    not null
  comment '宋江概率',
  pic4         int default '0'    not null
  comment '林冲概率',
  pic5         int default '0'    not null
  comment '鲁智概率',
  pic6         int default '0'    not null
  comment '金刀概率',
  pic7         int default '0'    not null
  comment '银枪概率',
  pic8         int default '0'    not null
  comment '铁斧概率',
  full_value   int default '0'    not null
  comment '全盘入口',
  rand_value   int default '0'    not null
  comment '随机入口',
  zero_value   int default '0'    not null
  comment '空门入口',
  max_award    bigint default '0' not null
  comment '最大奖金'
);

-- auto-generated definition
create table user
(
  uid         bigint unsigned         not null
    primary key,
  platform    varchar(11) default ''  not null
  comment '来源平台，1（默认值，一起秀）2（轻音）',
  nick        varchar(255) default '' not null
  comment '昵称',
  avatar      varchar(512) default '' not null
  comment '头像地址',
  sex         int default '0'         not null
  comment '0保密1男2女',
  coin        bigint default '0'      not null
  comment '游戏币',
  token       varchar(64) default ''  not null
  comment '登录令牌，用于网页登录时玩家信息的校验',
  coin_change bigint default '0'      not null
  comment '当天金币变化'
)
  charset = utf8mb4;

-- auto-generated definition
create table zero_rate
(
  id   int(11) unsigned not null
  comment '空门编号'
    primary key,
  col1 int default '0'  not null
  comment '第一列概率权重',
  col2 int default '0'  not null
  comment '第二列概率权重',
  col3 int default '0'  not null
  comment '第三列概率权重',
  col4 int default '0'  not null
  comment '第四列概率权重',
  col5 int default '0'  not null
  comment '第五列概率权重'
);






