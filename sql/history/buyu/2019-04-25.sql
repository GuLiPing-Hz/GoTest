delete
from fort_log
where fort_id = 0;
alter table fort_log
    add constraint fort_log_pk
        unique (uid, fort_id);

alter table sub_act_cfg
    modify total_cnt int default 0 not null comment '总的数量,-1表示无限供应';
alter table sub_act_cfg
    modify sub_title varchar(128) null comment '子任务标题；仅用于 更新公告';
alter table sub_act_cfg
    modify sub_content varchar(8192) null comment '子任务详细内容；仅用于 更新公告';

INSERT INTO Buyu.activities (id, title, endTime, sortNo, isEnabled)
VALUES (11, '鱼塘排位', '2019-12-31 00:00:00', 11, 0);
INSERT INTO `Buyu`.`activities` (`id`, `title`, `endTime`, `sortNo`, `isEnabled`)
VALUES (12, '鱼货达人', '2019-12-31 00:00:00', 12, 0);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (1, 11, 10, 50000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (1, 11, 3, 10);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (2, 11, 10, 30000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (2, 11, 3, 5);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (3, 11, 10, 10000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (3, 11, 3, 3);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (4, 11, 10, 5000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (4, 11, 3, 1);

INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (1, 12, 10, 10000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (2, 12, 10, 20000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (3, 12, 10, 50000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (4, 12, 10, 200000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (5, 12, 10, 500000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (6, 12, 10, 5000000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (7, 12, 10, 10000000);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`)
VALUES (8, 12, 10, 20000000);

INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 1, 10000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 2, 50000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 3, 100000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 4, 500000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 5, 1000000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 6, 5000000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 7, 10000000, -1, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (12, 8, 50000000, -1, DEFAULT, DEFAULT);

alter table notice
    modify sender bigint null comment '发送者，0:系统（默认）1:管理员 2:其他';
alter table notice
    modify id bigint auto_increment comment '公告编号';
alter table notice
    modify mail_type int null comment '参考trello：awards_type';
alter table fort_log
    modify add_type tinyint(3) default 0 not null comment '参见trello数据库类型字段说明- AWARDS_TYPE';
create index notice_addtime_index
    on notice (addtime);
alter table cdkey_cfg
    add flavor varchar(100) null comment '能使用的渠道,null 表示所有渠道都可以使用,也可以存放uid，表示指定用户使用';

alter table cdkey_cfg
    add reg_tm datetime default '2017-01-01' not null comment '使用的注册最晚时间，低于此时间注册将不能使用兑换码';


create table hbq_dui_log
(
    id      bigint primary key auto_increment,
    tm      datetime not null,
    uid     bigint   not null,
    hbqType int      not null comment '兑换的红包类型'
) comment '红包券兑换日志';

create table diamond_log
(
    id          bigint primary key auto_increment,
    uid         bigint   not null,
    cur         int      not null comment '用户当前最新数量',
    var         int      not null comment '钻石变化量',
    change_type int      null comment '变化类型',
    game_type   int      not null comment '游戏类型：
        LOBBY = 0, -- 大厅
        FISH = 1, -- 捕鱼游戏
        OTHER = 2, -- 其他
        LOTTERY = 3, -- 抽奖
        GAME_YULE = 8, -- 鱼乐游戏
        GAME_SLOT = 9 -- 水浒传',
    room_id     int      not null comment '房间ID or 订单ID',
    tm          datetime not null
);

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
) COMMENT = '消息表';
create index msg_tm_index
    on msg (tm);


CREATE TABLE yt
(
    ytid    int PRIMARY KEY COMMENT '鱼塘id',
    uid     bigint                      NOT NULL
        COMMENT '创建者uid',
    name    varchar(50) charset utf8mb4 NOT NULL
        COMMENT '鱼塘名称',
    intro   varchar(120) charset utf8mb4 COMMENT '鱼塘简介',
    reward  int    DEFAULT 0
        COMMENT '鱼塘每日签到奖励',
    `limit` int    DEFAULT 100
        COMMENT '人数上限',
    tm      datetime COMMENT '创建时间',
    pool    bigint default 0
        comment '鱼塘资金，鱼塘资金的变动不会增加ver版本信息，所以每次客户端请求，pool都会下发新值',
    act     bigint default 0
        comment '鱼塘当前活跃度，根据yt_clear_cfg来判断是否需要清零',
    ver     bigint default 1
        COMMENT '鱼塘信息版本。每当鱼塘信息变更了，ver++'
) COMMENT = '鱼塘表';

drop view if exists view_yt;
create view view_yt as
select a.*, nick_name as nickname, avatar
from yt a
         inner join user b on a.uid = b.uid;

CREATE TABLE yt_user
(
    uid     bigint   NOT NULL,
    ytid    bigint   NOT NULL
        COMMENT '鱼塘id 0表示未加入鱼塘',
    tm      datetime NOT NULL
        COMMENT '加入时间/申请时间',
    yuhuo   bigint  default 0
        comment '在某个鱼塘的累计收取/偷取鱼货；申请加入时携带的鱼货，离开鱼塘马上清0',
    checkin int     default 0
        comment '累计签到次数，离开鱼塘马上清零',
    utc     bigint  default 0 COMMENT '单位秒，离开时间-保护期24小时，此期间无法再申请加入其它鱼塘',
    apply   tinyint default 1 comment '是否申请的标志，1表示正在申请，0表示正式鱼塘用户',

    constraint yt_user_pk
        primary key (uid, ytid)
) COMMENT = '鱼塘用户表';
create index yt_user_tm_index
    on yt_user (tm);
create index yt_user_uid_index
    on yt_user (uid);


drop view if exists view_yt_membcnt;
create view view_yt_membcnt as
select ytid, count(1) as people
from yt_user
where apply = 0
  and ytid > 0
group by ytid;

drop view if exists view_yt_rank_reward;
create view view_yt_rank_reward as
select a.ytid,
       a.uid,
       nick_name as nickname,
       avatar,
       name,
       intro,
       reward,
       `limit`,
       act,
       people
from yt a
         inner join user b on a.uid = b.uid
         inner join view_yt_membcnt c on a.ytid = c.ytid
order by reward desc;


drop view if exists view_yt_rank_act;
create view view_yt_rank_act as
select a.ytid,
       a.uid,
       nick_name as nickname,
       avatar,
       name,
       intro,
       reward,
       `limit`,
       act,
       people
from yt a
         inner join user b on a.uid = b.uid
         inner join view_yt_membcnt c on a.ytid = c.ytid
order by act desc;

create table yt_yuhuo
(
    uid      bigint PRIMARY KEY NOT NULL,
    yuhuocur bigint default 0 comment '今日鱼货',
    yuhuoutc bigint default -1 comment '鱼货状态(单位秒)
-2表示正在初始化，-1表示关闭状态，0表示开启,>0表示倒计时的截止时间'
) COMMENT = '今日鱼货数据，每日6:50分清空表数据';

drop view if exists view_yt_user;
create view view_yt_user as
select a.uid,
       ytid,
       nick_name           as nickname,
       avatar,
       a.tm                as tm,
       ifnull(yuhuocur, 0) as yuhuocur,
       ifnull(yuhuoutc, -2) as yuhuoutc,
       yuhuo,
       utc,
       checkin
from yt_user a
         inner join user b on a.uid = b.uid
         left join yt_yuhuo c on a.uid = c.uid
where ytid > 0
  and apply = 0;

drop view if exists view_coin_log_fee;

drop view if exists view_yt_apply;
create view view_yt_apply as
select a.uid,
       ytid,
       nick_name          as nickname,
       avatar,
       a.tm               as tm,
       floor(-fee / 1000) as yuhuo
from yt_user a
         inner join user b on a.uid = b.uid
         left join (select uid, date(add_time) as tm, sum(fee) as fee
                    from coin_log
                    where change_type in (2, 74)
                      and add_time > date_sub(now(), interval -1 day)
                    group by uid, date(add_time)) c on a.uid = c.uid
where ytid > 0
  and apply = 1;
select *
from view_yt_apply
where ytid = 165272
limit 50;



# 获取某个鱼塘的排名信息
# select num
# from (select a.ytid, (@rowNum := @rowNum + 1) as num
#       FROM view_yt_rank_act as a,
#            (SELECT (@rowNum := 0)) b) c
# where c.ytid = 165271;

# -----------------------------------------------------------------------------------

CREATE TABLE yt_bill_log
(
    ytid int      NOT NULL,
    tm   datetime NOT NULL,
    bill bigint DEFAULT 0
        COMMENT '鱼塘流水',
    CONSTRAINT yt_bill_log_ytid_tm_pk PRIMARY KEY (ytid, tm)
) COMMENT = '鱼塘每日流水表';

CREATE TABLE yt_coin_log
(
    id     bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    tm     datetime           NOT NULL,
    uid    bigint             NOT NULL,
    uuid   varchar(100) comment '签到的设备码,只有签到的时候需要记录，偷取不需要',
    ytid   bigint             NOT NULL,
    reward int                NOT NULL
        COMMENT '奖励信息。或者被偷信息。',
    type   smallint           NOT NULL
        COMMENT '0表示鱼塘签到，1表示偷取鱼货 (鱼货的收取和被偷起始时间都是早上 2:00 am)',
    optuid bigint             NOT NULL
        COMMENT '-1表示系统，0表示自己，如果type是1，表示被偷的uid，'
) COMMENT = '鱼塘次数日志记录';
create index yt_coin_log_tm_index
    on yt_coin_log (tm);

# 查询偷取记录
drop view if exists view_steal_log;
create view view_steal_log as
select a.uid, b.nick_name as name, a.reward, a.tm, a.optuid, c.nick_name as optname
from yt_coin_log a
         inner join user b on a.uid = b.uid
         inner join user c on a.optuid = c.uid;

create table yt_create_cfg
(
    id             tinyint primary key default 1,
    coin           bigint              default 3000000           not null comment '创建鱼塘需要消耗的金币->划归鱼塘',
    vip            tinyint             default 0 comment '创建鱼塘需要的VIP',
    reward         int                 default 1000 comment '鱼塘每日签到系统奖励',
    modify_diamond int                 default 100 comment '修改鱼塘信息需要花费的钻石',
    tm             datetime            default CURRENT_TIMESTAMP not null
        comment '清理活跃起始日期，此后按周期计算到了就再次清空鱼塘活跃度',
    life           smallint(6)         default 7                 not null
        comment '周期，以天为单位'
);
insert into yt_create_cfg(vip, tm) value (0, date(now()));

create table yt_rank_last
(
    ytid bigint primary key not null comment '鱼塘id',
    act  bigint default 0 comment '活跃度',
    num  int comment '排名'
) comment '活跃度清0前的排名值';

drop view if exists view_yt_rank_act_last;
create view view_yt_rank_act_last as
select a.ytid,
       nickname,
       name,
       a.act,
       num
from yt_rank_last a
         inner join view_yt_rank_act b on a.ytid = b.ytid;
# ---------------------------------存储过程 -----------------------------------
# ---------------------------------存储过程 -----------------------------------
# ---------------------------------存储过程 -----------------------------------
# ---------------------------------存储过程 -----------------------------------

-- ----------------------------
-- Procedure structure for `proc_update_yuhuo_by_day` begin
-- 每日6:50分执行
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_update_yuhuo_by_day;
DELIMITER //
CREATE PROCEDURE proc_update_yuhuo_by_day(in vTm datetime)
exec:
BEGIN
    declare tm_yesterday datetime;
    set tm_yesterday = date_sub(vTm, INTERVAL 1 DAY);

    #清空一下。
    delete from yt_yuhuo where true;

    DROP TEMPORARY TABLE IF EXISTS tmp_yt_yhuo;
    CREATE TEMPORARY TABLE tmp_yt_yhuo
    select a.uid, b.ytid, floor(-sum(fee) / 1000) as bill
    from coin_log a
             inner join yt_user b on a.uid = b.uid
    where b.apply = 0
      and a.change_type = 74
      and a.add_time >= tm_yesterday
      and a.add_time < vTm
    group by a.uid, b.ytid
    having bill >= 1000;

    #更新用户的累计鱼货。
    #update yt_user a inner join tmp_yt_yhuo b on a.uid = b.uid set yuhuo=yuhuo + bill;
    #更新鱼塘的活跃度。
    update yt a inner join (select ytid, sum(bill) as bill1 from tmp_yt_yhuo group by ytid) b on a.ytid = b.ytid
    set a.act = a.act + bill1;

    insert into yt_yuhuo select uid, bill, -1 from tmp_yt_yhuo;
    update yt set ver=ver + 1 where true;

    -- 删除临时表
    DROP TEMPORARY TABLE IF EXISTS tmp_yt_yhuo;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_update_yuhuo_by_day` END
-- ----------------------------

-- ----------------------------
-- Procedure structure for `proc_create_yt` begin
-- 创建一个鱼塘
# @return status:
#     0 成功
#     10123, --您已加入一个鱼塘
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_create_yt;
DELIMITER //
CREATE PROCEDURE proc_create_yt(in vUid bigint, in vName text, in vIntro text,
                                in vReward int, in vTm datetime, in vPool bigint)
exec:
BEGIN
    declare vCnt int;
    select count(1) into vCnt from yt_user where uid = vUid and apply = 0 and ytid > 0;
    if vCnt > 0 then
        select 10123 as status;
        leave exec;
    end if;

    #删除所有的申请消息
    delete from yt_user where uid = vUid;
    #加入到鱼塘中
    insert into yt_user(uid, ytid, tm, apply) value (vUid, vUid, vTm, 0);
    #创建一个鱼塘
    insert into yt(ytid, uid, name, intro, reward, tm, pool)
        value (vUid, vUid, vName, vIntro, vReward, vTm, vPool);
    select 0 as status;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_create_yt` END
-- ----------------------------
update yt
set `limit` = 200,
    ver=ver + 1
where true;
-- ----------------------------
-- Procedure structure for `proc_applay_yt` begin
-- 加入一条申请入鱼塘
# @return
#     status:
#         0 提交成功
#         10131 重复申请（服务器统一错误码）
#         10116 提交申请达上限 20个（服务器统一错误码）
#         10136 您申请的鱼塘已经人满了
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_applay_yt;
DELIMITER //
CREATE PROCEDURE proc_applay_yt(in vUid bigint, in vYtid bigint, in vTm datetime)
exec:
BEGIN
    declare v24HoursAgo datetime;
    declare vApplyCnt int;
    declare vMemCnt int;
    declare vYuHuo bigint;

    select count(1) into vApplyCnt from view_yt_apply where uid = vUid and ytid = vYtid;
    if vApplyCnt > 0 then
        select 10131 as status;
        leave exec;
    end if;

    select count(1) into vMemCnt from yt_user where ytid = vYtid and apply = 0;
    if vMemCnt >= 200 then
        select 10136 as status;
        leave exec;
    end if;

    if vApplyCnt >= 20 then
        select 10116 as status;
        leave exec;
    end if;

    select count(1) into vApplyCnt from view_yt_apply where uid = vUid;
    if vApplyCnt >= 20 then
        select 10116 as status;
        leave exec;
    end if;

    set v24HoursAgo = date_sub(vTm, interval 24 hour);
    select -sum(fee) into vYuHuo
    from coin_log
    where uid = vUid
      and change_type in (2, 74) #普通捕鱼，鱼塘捕鱼
      and add_time > v24HoursAgo;

    insert into yt_user(uid, ytid, tm, yuhuo) value (vUid, vYtid, vTm, ifnull(vYuHuo, 0));
    select 0 as status;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_applay_yt` END
-- ----------------------------

-- ----------------------------
-- Procedure structure for `proc_disband_yt` begin
-- 解散鱼塘
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_disband_yt;
DELIMITER //
CREATE PROCEDURE proc_disband_yt(in vYtid bigint, in vUtc bigint)
exec:
BEGIN
    delete from yt_user where ytid = vYtid and apply = 1;
    update yt_user set ytid=0, yuhuo=0, checkin=0, utc=vUtc where ytid = vYtid;
    delete from yt where ytid = vYtid;
END
//
#分隔符还原
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_disband_yt` begin
-- 同意申请加入鱼塘
-- 10132, --操作失败,请求已过时
-- 10136 鱼塘满员
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_accept_apply;
DELIMITER //
CREATE PROCEDURE proc_accept_apply(in vYtid bigint, in vUid bigint, in vTm datetime)
exec:
BEGIN

    declare vCnt int;
    declare vMemCnt int;
    select count(1) into vCnt from view_yt_apply where ytid = vYtid and uid = vUid;

    if vCnt = 0 then
        select 10132 as status;
        leave exec;
    end if;

    select count(1) into vMemCnt from yt_user where ytid = vYtid and apply = 0;
    if vMemCnt >= 200 then
        select 10136 as status;
        leave exec;
    end if;

    #删除该玩家的其他申请
    delete from yt_user where uid = vUid;
    insert into yt_user(uid, ytid, tm, apply) value (vUid, vYtid, vTm, 0);
    update yt set ver=ver + 1 where ytid = vYtid;
    select 0 as status, a.* from view_yt_user a where uid = vUid;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_accept_apply` END
-- ----------------------------
# select count(1) into vCnt from view_yt_apply where ytid = vYtid and uid = vUid;

-- ----------------------------
-- Procedure structure for `proc_update_yuhuo_by_day` begin
-- 签到领取鱼塘奖励
# @return
#     status: 返回状态值
#         10090 签到已经领取
#         10114 请求参数错误，无法获取到对应的鱼塘奖励数据
#         10134 鱼塘资金不足，无法签到
#     reward: 签到奖励
#     pool:鱼塘当前资金
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_yt_checkin;
DELIMITER //
CREATE PROCEDURE proc_yt_checkin(in vUid bigint, in vUUID text, in vYtid bigint, in vNow datetime)
exe:
BEGIN
    declare vToday datetime;
    declare vCnt int;
    declare vReward int;
    declare vSystemReward int;
    declare vPool bigint;
    declare vYtRankEnable int;
    declare vYtRank int;
    declare vActivityReward int;

    set vActivityReward = 0;
    set vToday = date(vNow);

    select count(1) into vCnt
    from yt_coin_log
    where (uuid = vUUID or uid = vUid)
      and tm >= vToday
      and type = 0;

    if vCnt > 0 then
        select 10090 as status;
        leave exe;
    end if;

    select reward, pool into vReward,vPool from yt where ytid = vYtid;

    #判断yuid是否非法
    if vReward is null or vPool is null then
        select 10114 as status;
        leave exe;
    end if;

    select reward into vSystemReward from yt_create_cfg limit 1;

    if vSystemReward is null then
        set vSystemReward = 1000;
    end if;

    #判断当前鱼塘资金是否够发签到奖励
    if vReward - vSystemReward > vPool then
        select 10134 as status;
        leave exe;
    end if;

    select count(1) into vYtRankEnable from activities where id = 12 and isEnabled = 1 and endTime > vNow;
    if vYtRankEnable > 0 then
        select ifnull(num, 0) into vYtRank from yt_rank_last where ytid = vYtid;
        if vYtRank <= 3 then
            select gift_count into vActivityReward
            from sub_act_gift
            where activity_id = 11
              and sub_id = vYtRank
              and gift_type = 10;
        elseif vYtRank <= 10 then
            select gift_count into vActivityReward
            from sub_act_gift
            where activity_id = 11
              and sub_id = 4
              and gift_type = 10;
        end if;
    end if;

    #记录个人签到金额
    insert yt_coin_log(tm, uid, uuid, ytid, reward, type, optuid)
        value (vNow, vUid, vUUID, vYtid, vReward, 0, 0);

    #减少鱼塘库存金币,鱼塘成员累计签到数更新。。。
    set vPool = vPool - (vReward - vSystemReward);
    update yt set pool=vPool, ver=ver + 1 where ytid = vYtid;

    #更新鱼塘成员累计签到。
    update yt_user set checkin=checkin + 1 where uid = vUid;

    #返回签到的金币数量
    select 0 as status, vReward as reward, vPool as pool, vActivityReward as reward2;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_yt_checkin` END
-- ----------------------------

-- ----------------------------
# Procedure structure for `proc_get_yuhuo` begin
# 收取鱼货或者偷取别人的鱼货
# @return
#     status: 状态值
#         0表示开启,
#
#         >0 表示倒计时的截止时间——现在还没到。
#         -1 表示关闭状态
#         -2 表示正在初始化
#         -3 表示鱼货已经被收取
#         -4 一天一个uid只能被同一个uid偷一次
#         -5 一天一个uid最多只能偷别人10次
#     yuhuo:
# ----------------------------
DROP PROCEDURE IF EXISTS proc_get_yuhuo;
DELIMITER //
CREATE PROCEDURE proc_get_yuhuo(in vUid bigint, in vOptedUid bigint,
                                in vYtid bigint, in vTm datetime, in vUtc bigint)
exec:
BEGIN
    declare vYuhuo bigint;
    declare vYuhuoUtc bigint;
    declare vPercent float;
    declare vSteal int;
    declare vToday datetime;
    declare vStealSameCnt int;
    declare vStealOtherCnt int;

    set vToday = date(vTm);
    select yuhuocur, yuhuoutc into vYuhuo,vYuhuoUtc from yt_yuhuo where uid = vOptedUid;

    if vYuhuoUtc is null then
        select -2 as status, 0 as yuhuo;
        leave exec;
    end if;

    if vYuhuoUtc < 0 or vYuhuoUtc > vUtc then #收取或者偷取时间尚未达到
        select vYuhuoUtc as status, 0 as yuhuo;
        leave exec;
    end if;

    if vYuhuo = 0 then
        select -3 as status, 0 as yuhuo;#-3 表示鱼货已经被收取
    end if;


    #如果vUid和vOptedUid表示收取自己的鱼货，反之则是偷取别人的鱼货
    if vUid = vOptedUid then
        update yt_yuhuo set yuhuocur=0 where uid = vOptedUid;
        #更新当前的鱼塘版本信息
        update yt set ver=ver + 1 where ytid = vYtid;
        #加入到累计收取鱼货中(收取)
        update yt_user set yuhuo=yuhuo + vYuhuo where uid = vUid;
        select 0 as status, vYuhuo as yuhuo;
    else
        #一天一个uid只能被同一个uid偷一次
        select count(1) into vStealSameCnt
        from yt_coin_log
        where tm >= vToday
          and uid = vUid
          and optuid = vOptedUid
          and type = 1;
        if vStealSameCnt > 0 then
            select -4 as status, 0 as yuhuo;
            leave exec;
        end if;

        #一天一个uid最多只能偷别人10次
        select count(1) into vStealOtherCnt from yt_coin_log where tm >= vToday and uid = vUid and type = 1;
        if vStealOtherCnt > 10 then
            select -5 as status, 0 as yuhuo;
            leave exec;
        end if;

        select rand() * 0.01 + 0.01 into vPercent;
        set vSteal = floor(vYuhuo * vPercent);

        #记录偷取的金额
        insert into yt_coin_log(tm, uid, ytid, reward, type, optuid)
            value (vTm, vUid, vYtid, vSteal, 1, vOptedUid);
        #加入到累计收取鱼货中(偷取)
        update yt_user set yuhuo=yuhuo + vSteal where uid = vUid;
        #更新当前的鱼货金额
        update yt_yuhuo set yuhuocur=yuhuocur - vSteal where uid = vOptedUid;
        #更新当前的鱼塘版本信息
        update yt set ver=ver + 1 where ytid = vYtid;
        select 0 as status, vSteal as yuhuo;
    end if;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_get_yuhuo` END
-- ----------------------------

-- ----------------------------
-- Procedure structure for `proc_clear_yt_act` begin
-- 清空鱼塘活跃度,清空前把排名放到yt_rank_last。
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_clear_yt_act;
DELIMITER //
CREATE PROCEDURE proc_clear_yt_act()
exec:
BEGIN
    delete from yt_rank_last where true;
    insert into yt_rank_last
    select ytid, act, (@rowNum := @rowNum + 1) as num
    from view_yt_rank_act a,
         (select (@rowNum := 0)) b;
    update yt set act=0 where true;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_clear_yt_act` END
-- ----------------------------


