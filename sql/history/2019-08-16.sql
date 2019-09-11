# 优化每日任务，防止出现插入重复的数据

create table online_reward_log2
(
    uid       bigint  not null comment '用户ID',
    reward_id tinyint not null comment 'online_reward_cfg中的id',
    state     tinyint  default 0 comment '当前任务状态',
    tm        datetime default now() comment '数据更新时间',
    constraint online_reward_log2_pk
        primary key (uid, reward_id)
)
    comment '每日任务优化存储数据';


drop procedure if exists proc_reset_by_day;
create procedure proc_reset_by_day(IN tm datetime, IN tm_utc bigint, IN vip_limit int)
BEGIN
    #Routine body goes here...
    #     DECLARE sumTotal INT; #声明变量sumTotal
    #每日0点清空在线日志
    call online_log_summary(tm); #数据统计。
    DELETE FROM online_log2 where true;
    DELETE from online_reward_log2 where true;
    #清理表
    #更新每日捕获龙的记录
    INSERT INTO catch_dragon_log (uid, dragon_cnt, tm)
    SELECT uid,
           value,
           tm
    FROM mission
    WHERE value > 0
      AND mid = 4;
    #用户每日状态信息清空
    UPDATE user_stat
    SET lz_points     = 0,
        transfer_coin = 0
    where true;

    #更新VIP每日奖励领取状态
    update user_stat
    set isgetvipreward = 1
    WHERE vip >= vip_limit;
    #     update user_stat
    #     set isgetvipreward = 0
    #     WHERE vip < vip_limit;

    #更新任务状态
    UPDATE mission
    set value = 0,
        state = 0
    where true;

    #更新活跃度金币
    UPDATE active_coin
    SET state     = 0,
        update_tm = tm
    where true;

    #把炮台过期的炮置为默认炮台
    UPDATE user
        inner join fort_log on user.uid = fort_log.uid and weapon = fort_id
    SET weapon = 0
    WHERE expire_utc > 0
      AND expire_utc < tm_utc;

    #更新周月卡记录，把过期的置为无效状态。
    UPDATE card_log
    SET state = 0
    WHERE end_tm < tm_utc
      and state = 1;
END;

INSERT INTO `Buyu`.`activities` (`id`, `title`, `endTime`, `sortNo`, `isEnabled`)
VALUES (13, '找回账号', '2019-10-01 00:00:00', 13, 1);
INSERT INTO `Buyu`.`activities` (`id`, `title`, `endTime`, `sortNo`, `isEnabled`)
VALUES (14, '游戏推广', '2019-10-01 00:00:00', 14, 1);

alter table user
    modify type tinyint default 0 null comment '用户类型：
0.普通用户
1.金推推广员(不可登录游戏，可登后台)
2.普通推广员(可登录游戏，不可登后台)
3.渔乐机器人，
4.捕鱼机器人，
5.后台推广（待审核,不可登游戏，不可登后台）
6.银推推广员(不可登录游戏，可登后台)
7~10.推广员预留(不可登录游戏，可登后台)
';

alter table invite_log
    add status tinyint default 2 null comment '推广成员的奖励是否已经领取，默认已领取。';

# 更新邀请奖励
update invite_cfg
set inviterGift=50000
where id = 1;

-- ----------------------------
-- Procedure structure for `proc_add` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_get_my_promotion;
CREATE PROCEDURE proc_get_my_promotion(in vUid bigint)
BEGIN
    declare cnt int;
    declare cnt2 int;
    #查询总共推广了多少用户
    select count(1) into cnt from invite_log where code = vUid;
    #查询当前还有多少推广用户已经升到5级，但是还没有领取的推广奖励
    select count(1) into cnt2 from invite_log where code = vUid and status = 1;
    select cnt, cnt2;
END;
-- ----------------------------
-- Procedure structure for `proc_get_my_promotion` END
-- ----------------------------
# call proc_get_my_promotion(165272);

alter table yt
    add tgy bigint default 0 null comment '鱼塘推广员。';

-- ----------------------------
-- Procedure structure for `proc_create_yt` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_create_yt;
CREATE PROCEDURE proc_create_yt(in vUid bigint, in vName text, in vIntro text,
                                in vReward int, in vTgy bigint, in vTm datetime, in vPool bigint)
exec:
BEGIN
    # 创建一个鱼塘
    # @return status:
    #     0 成功
    #     10123, --您已加入一个鱼塘

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
    insert into yt(ytid, uid, name, intro, reward, tm, pool, tgy)
        value (vUid, vUid, vName, vIntro, vReward, vTm, vPool, vTgy);
    select 0 as status;
END;
-- ----------------------------
-- Procedure structure for `proc_create_yt` END
-- ----------------------------

drop view if exists view_yt_rank_act_last;
create view view_yt_rank_act_last as
select a.ytid,
       nickname,
       name,
       a.act,
       num
from yt_rank_last a
         inner join view_yt_rank_act b on a.ytid = b.ytid
where a.act > 0;

UPDATE `Buyu`.`yt_create_cfg` t
SET t.`reward` = 0
WHERE t.`id` = 1;

alter table user
    alter column week_card_utc set default 0;
alter table user
    alter column ID_number set default '';

alter table invite_log
    modify status tinyint default 2 null comment '推广成员的奖励是否已经领取，默认已领取。0未完成，1完成待领取，2已领取';

alter table wares_cfg
    add `desc` varchar(100) charset utf8mb4 null comment '传入到支付平台的商品描述';

alter table wares_cfg
    drop column mail_giftid;

delete
from wares_cfg
where true;
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('10', 1, 8, 98, 0, 980000, 980000, 0, 0, 0, 0, 588000, '98元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('101', 2, 65536, 6, 0, 0, 0, 60, 60, 0, 0, 48, '6元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('102', 2, 131072, 12, 0, 0, 0, 120, 120, 0, 0, 96, '12元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('103', 2, 262144, 30, 0, 0, 0, 300, 300, 0, 0, 240, '30元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('104', 2, 524288, 98, 0, 0, 0, 980, 980, 0, 0, 784, '98元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('105', 2, 1048576, 198, 0, 0, 0, 1980, 1980, 0, 0, 1584, '198元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('106', 2, 2097152, 648, 0, 0, 0, 6480, 6480, 0, 0, 5184, '648元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('11', 1, 16, 198, 0, 1980000, 1980000, 0, 0, 0, 0, 1188000, '198元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('12', 1, 32, 648, 0, 6480000, 6480000, 0, 0, 0, 0, 3888000, '648元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('200', 1, 0, 50, 0, 900000, 0, 0, 0, 0, 0, 0, '推广员50元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('201', 1, 0, 100, 0, 1800000, 0, 0, 0, 0, 0, 0, '推广员100元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('202', 1, 0, 500, 0, 9000000, 0, 0, 0, 0, 0, 0, '推广员500元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('203', 1, 0, 1000, 0, 18000000, 0, 0, 0, 0, 0, 0, '推广员1000元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('204', 1, 0, 5000, 0, 90000000, 0, 0, 0, 0, 0, 0, '推广员5000元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('7', 1, 1, 6, 0, 60000, 60000, 0, 0, 0, 0, 36000, '6元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('8', 1, 2, 12, 0, 120000, 120000, 0, 0, 0, 0, 72000, '12元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('9', 1, 4, 30, 0, 300000, 300000, 0, 0, 0, 0, 180000, '30元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms1', 6, 0, 6, 0, 120000, 0, 0, 0, 0, 0, 0, '破产礼包1');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms12', 6, 0, 30, 0, 540000, 0, 0, 0, 0, 0, 0, '破产礼包3');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms120', 6, 0, 120, 0, 2160000, 0, 0, 0, 0, 0, 0, '破产礼包6');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms24', 6, 0, 50, 0, 900000, 0, 0, 0, 0, 0, 0, '破产礼包4');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms6', 6, 0, 10, 0, 180000, 0, 0, 0, 0, 0, 0, '破产礼包2');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('alms60', 6, 0, 80, 0, 1440000, 0, 0, 0, 0, 0, 0, '破产礼包5');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.11', 1, 0, 30, 0, 300000, 0, 0, 0, 0, 0, 0, '苹果30元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.22', 1, 0, 50, 0, 500000, 0, 0, 0, 0, 0, 0, '苹果50元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.33', 1, 0, 88, 0, 880000, 0, 0, 0, 0, 0, 0, '苹果88元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.44', 1, 0, 128, 0, 1280000, 0, 0, 0, 0, 0, 0, '苹果128元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.55', 1, 0, 158, 0, 1580000, 0, 0, 0, 0, 0, 0, '苹果158元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.66', 1, 0, 198, 0, 1980000, 0, 0, 0, 0, 0, 0, '苹果198元金币');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.1', 2, 0, 30, 0, 0, 0, 30, 0, 0, 0, 0, '苹果30元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.2', 2, 0, 50, 0, 0, 0, 50, 0, 0, 0, 0, '苹果50元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.3', 2, 0, 88, 0, 0, 0, 88, 0, 0, 0, 0, '苹果88元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.4', 2, 0, 128, 0, 0, 0, 128, 0, 0, 0, 0, '苹果128元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.5', 2, 0, 158, 0, 0, 0, 158, 0, 0, 0, 0, '苹果158元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.diamond.6', 2, 0, 198, 0, 0, 0, 198, 0, 0, 0, 0, '苹果188元钻石');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.pkgfort', 8, 536870912, 198, 0, 3600000, 0, 0, 0, 0, 0, 0, '炮倍礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.sevenday', 4, 0, 12, 0, 20000, 0, 10, 0, 0, 0, 0, '周卡');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('lailai.fish.thirtyday', 4, 1073741824, 50, 35, 20000, 0, 0, 0, 0, 0, 0, '月卡');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg128', 5, 0, 128, 0, 2360000, 0, 880, 0, 8, 0, 0, '128元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg18', 5, 0, 18, 0, 320000, 0, 200, 0, 1, 600, 0, '18元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg198', 5, 0, 198, 0, 3600000, 0, 1400, 0, 8, 0, 0, '198元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg30', 5, 0, 30, 0, 540000, 0, 350, 0, 3, 600, 0, '30元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg68', 5, 0, 68, 0, 1260000, 0, 800, 0, 2, 600, 0, '68元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkg98', 5, 0, 98, 0, 1800000, 0, 600, 0, 6, 600, 0, '98元限购礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('pkgnewbie', 7, 268435456, 5, 0, 100000, 0, 50, 0, 0, 0, 0, '新手礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx10', 3, 0, -10, 0, 0, 0, 0, 0, 0, 0, 0, '10元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx100', 3, 0, -100, 0, 0, 0, 0, 0, 0, 0, 0, '100元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx2', 3, 0, -2, 0, 0, 0, 0, 0, 0, 0, 0, '2元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx20', 3, 0, -20, 0, 0, 0, 0, 0, 0, 0, 0, '20元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx5', 3, 0, -5, 0, 0, 0, 0, 0, 0, 0, 0, '5元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('tx50', 3, 0, -50, 0, 0, 0, 0, 0, 0, 0, 0, '50元红包兑换');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip10', 10, 0, 50, 0, 1100000, 0, 0, 0, 0, 0, 0, 'VIP10特权礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip5', 10, 0, 6, 0, 120000, 0, 0, 0, 0, 0, 0, 'VIP5特权礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip6', 10, 0, 12, 0, 240000, 0, 0, 0, 0, 0, 0, 'VIP6特权礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip7', 10, 0, 30, 0, 600000, 0, 0, 0, 0, 0, 0, 'VIP7特权礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip8', 10, 0, 30, 0, 630000, 0, 0, 0, 0, 0, 0, 'VIP8特权礼包');
INSERT INTO Buyu.wares_cfg (wares_id, wares_type, first_pay, price, discount, coin, first_reward_coin, diamond,
                            first_reward_diamond, room_ids, countdown, ext_value, info)
VALUES ('vip9', 10, 0, 50, 0, 1050000, 0, 0, 0, 0, 0, 0, 'VIP9特权礼包');

alter table coin_log
    drop column diamond_before;
alter table coin_log
    drop column diamond_after;
alter table coin_log
    drop column diamond_change;

alter table pay_log
    add reid bigint default 0 null comment '返利ID，默认0，如果绑定了推荐关系，则是推荐人，如果在鱼塘里面，则是鱼塘推广员';

alter table pay_log
    add retype tinyint default 0 null comment '默认0,0无类型，1推广用户充值，2鱼塘充值,3出售金币';

alter table pay_log
    add reid2 bigint default 0 null comment '返利金推ID，默认0';

-- ----------------------------
-- Procedure structure for `proc_insert_pay` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_insert_pay;
CREATE PROCEDURE proc_insert_pay(in vUid bigint, in vTradeNo text, in vChannel tinyint, in vWaresId text,
                                 vMoney float, vTm datetime)
BEGIN
    declare vReUid bigint;
    declare vReUid2 bigint;
    declare vReType tinyint;
    declare vYTid bigint;

#     首先检查是否有推荐人
    select code into vReUid from invite_log where uid = vUid;
    if vReUid is null then
#         其次检查是否加入鱼塘
        select ytid into vYTid from yt_user where uid = vUid and apply = 0;
        if vYTid is not null then
#             如果加入鱼塘了，当前是否有推广员
            select tgy into vReUid from yt where ytid = vYTid;
            if vReUid is not null then
                set vReType = 2;
            end if;
        end if;
    else
        set vReType = 1;
    end if;

    if vReUid is null then
        set vReUid = 0;
        set vReType = 0;
        set vReUid2 = 0;
    else
        select bindUid into vReUid2 from agent_user_info where uid = vReUid;
    end if;

    #     select vReUid;
    insert into pay_log(tradeno, channel, uid, waresid, money, addtime, transtime, reid, retype, reid2)
        value (vTradeNo, vChannel, vUid, vWaresId, vMoney, vTm, vTm, vReUid, vReType, vReUid2);
    select oid as insert_id from pay_log where tradeno = vTradeNo;
END;
-- ----------------------------
-- Procedure structure for `proc_insert_pay` END
-- ----------------------------

# call proc_get_my_promotion(165272);
call proc_insert_pay(188939, '', 1, '', 1, now());

select uname
from user
where uname like 'WX\\_%'
order by uname desc
limit 1 offset 0;
select count(1)
from user
where uname like 'WX\\_%';

create table user_cfg
(
    id   int default 1 null,
    wxid int default 1 null
);

create unique index user_cfg_id_uindex
    on user_cfg (id);

alter table user_cfg
    add constraint user_cfg_pk
        primary key (id);
INSERT INTO `Buyu`.`user_cfg` (`id`, `wxid`)
VALUES (DEFAULT, 100000);

DROP PROCEDURE IF EXISTS proc_get_wxid;
-- ----------------------------
-- Procedure structure for `proc_get_wxid` BEGIN
-- ----------------------------
CREATE PROCEDURE proc_get_wxid()
BEGIN
    update user_cfg set wxid=wxid + 1;
    select wxid from user_cfg where id = 1;
END;
-- ----------------------------
-- Procedure structure for `proc_get_wxid` END
-- ----------------------------

create index pay_log_tradeno_index
    on pay_log (tradeno);


drop table add_coin_log;


# SQL性能优化
create index agent_award_log_uid_index
    on agent_award_log (uid);
create index agent_award_cfg_uid_index
    on agent_award_cfg (uid);

create index yt_user_ytid_index
    on yt_user (ytid);
create index yt_tgy_index
    on yt (tgy);
create index yt_ytid_index
    on yt (ytid);


