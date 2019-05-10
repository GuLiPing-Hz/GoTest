CREATE TABLE msg
(
    id      bigint PRIMARY KEY NOT NULL comment '消息id 64位
1位保留-31位时间戳（单位秒，136年内不会重复）-8位集群id-4位服务器id-20位sequence',
    type    tinyint            NOT NULL
        COMMENT '消息类型:
        SYSTEM = 0, --系统消息；
        P2P = 1, --点对点消息
        TEAM = 2, --群组消息
        YT_APPLY = 3, --鱼塘申请消息',
    from_id bigint             NOT NULL comment '消息来源ID，可以为uid，或者ytid或者0表示系统',
    to_id   bigint             NOT NULL COMMENT '消息前往ID，可以为uid，或者ytid',
    tm      datetime           NOT NULL,
    msg     text COMMENT '消息内容。'
);
ALTER TABLE msg
    COMMENT = '消息表';
create index msg_tm_index
    on msg (tm);


CREATE TABLE yt
(
    ytid    int PRIMARY KEY AUTO_INCREMENT
        COMMENT '鱼塘id',
    uid     bigint      NOT NULL
        COMMENT '创建者uid',
    name    varchar(50) NOT NULL
        COMMENT '鱼塘名称',
    intro   text COMMENT '鱼塘简介',
    reward  int    DEFAULT 0
        COMMENT '鱼塘每日签到奖励',
    `limit` int    DEFAULT 100
        COMMENT '人数上限',
    tm      datetime COMMENT '创建时间',
    pool    bigint default 0
        comment '鱼塘资金',
    act     bigint default 0
        comment '鱼塘当前活跃度，根据yt_clear_cfg来判断是否需要清零',
    ver     bigint default 0
        COMMENT '鱼塘信息版本。每当鱼塘信息变更了，ver++'
);
ALTER TABLE yt
    COMMENT = '鱼塘表';

CREATE TABLE yt_user
(
    uid     bigint PRIMARY KEY NOT NULL,
    ytid    int                NOT NULL
        COMMENT '鱼塘id 0表示未加入鱼塘',
    tm      datetime           NOT NULL
        COMMENT '加入时间',
    yuhuo   bigint default 0
        comment '累计鱼货，离开鱼塘马上清零',
    checkin int    default 0
        comment '累计签到次数，离开鱼塘马上清零',
    utc     bigint COMMENT '单位秒，离开时间-保护期24小时，此期间无法再申请加入其它鱼塘'
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

create table yt_create_cfg
(
    coin           bigint  default 3000000 not null comment '创建鱼塘需要消耗的金币->划归鱼塘',
    vip            tinyint default 0 comment '创建鱼塘需要的VIP',
    reward         int     default 1000 comment '鱼塘每日签到系统奖励',
    modify_diamond int     default 100 comment '修改鱼塘信息需要花费的钻石'
);
insert into yt_create_cfg(vip) value (0);





