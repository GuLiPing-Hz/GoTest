create table z_huizong_qudao_hour
(
    tm             datetime         not null,
    flavor         varchar(50) comment '渠道来源',
    in_cnt         int    default 0 null comment '今日当前时间的登录人数',
    new_cnt        int    default 0 null comment '距离上一次统计时间段的注册人数，默认一小时',
    check_cnt      int    default 0 null comment '距离上一次统计时间段的签到次数，默认一小时',

    money          bigint default 0 null comment '距离上一次统计时间段的充值，默认一小时',
    hb             bigint default 0 null comment '距离上一次统计时间段的红包提现，默认一小时',
    hbq_hbc        bigint default 0 null comment '距离上一次统计时间段的红包场红包，默认一小时',
    hbq_lzc        bigint default 0 null comment '距离上一次统计时间段的龙珠场场红包，默认一小时',
    hbq_lzdb       bigint default 0 null comment '距离上一次统计时间段的龙珠夺宝红包，默认一小时',
    hbq_other      bigint default 0 null comment '距离上一次统计时间段的集福红包,每日任务红包，默认一小时',

    coin_pay       bigint default 0 comment '距离上一次统计时间充值增加金币',
    coin_day       bigint default 0 comment '距离上一次统计时间签到增加金币',
    coin_vip       bigint default 0 comment '距离上一次统计时间vip奖励增加金币',
    coin_card      bigint default 0 comment '距离上一次统计时间周月卡每日奖励增加金币',
    coin_rank      bigint default 0 comment '距离上一次统计时间排行榜奖励增加金币',
    coin_alms      bigint default 0 comment '距离上一次统计时间免费破产增加金币',
    coin_task      bigint default 0 comment '距离上一次统计时间成长任务增加金币',
    coin_mission   bigint default 0 comment '距离上一次统计时间在线任务,每日任务增加金币',
    coin_mail      bigint default 0 comment '距离上一次统计时间邮件奖励增加金币',
    coin_key       bigint default 0 comment '距离上一次统计时间CDKey增加金币',
    coin_invite    bigint default 0 comment '距离上一次统计时间邀请奖励增加金币',
    coin_lzc       bigint default 0 comment '距离上一次统计时间龙珠场抽宝箱增加金币',
    coin_yt        bigint default 0 comment '距离上一次统计时间鱼货提取/偷取增加金币',
    coin_fanli     bigint default 0 comment '距离上一次统计时间返利增加金币(含充值，售卖)',
    coin_other     bigint default 0 comment '距离上一次统计时间活动奖励76，解锁炮值(目前无需解锁)增加金币',

    bill_room2     bigint default 0 comment '距离上一次统计时间初级场当前流水',
    bill_room3     bigint default 0 comment '距离上一次统计时间中级场当前流水',
    bill_room8     bigint default 0 comment '距离上一次统计时间红包场当前流水',

    coin_fee_room2 bigint default 0 comment '距离上一次统计时间初级场预估消耗金币',
    coin_fee_room3 bigint default 0 comment '距离上一次统计时间中级场预估消耗金币',
    coin_fee_room8 bigint default 0 comment '距离上一次统计时间红包场预估消耗金币',
    coin_fee_fort  bigint default 0 comment '距离上一次统计时间购买炮台消耗金币',

    primary key (tm, flavor)
);
create index z_huizong_qudao_hour_tm_index
    on z_huizong_qudao_hour (tm);

-- ----------------------------
-- Procedure structure for `proc_huizong_by_platform` BEGIN
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_huizong_by_flavor;
CREATE PROCEDURE proc_huizong_by_flavor(in vFlavor text)
exec:
BEGIN

    #按渠道单位时间分类汇总
    declare vTmToday datetime;
    declare vTmBegin datetime;
    declare vTmEnd datetime;
    declare vTaskCoin bigint;
    declare vTaskBall bigint;
    declare vOtherCoin bigint;
    declare vOtherBall bigint;
    declare vLZCCoin bigint;
    declare vLZCBall1 bigint;
    declare vLZCBall2 bigint;
    declare vLZCBall3 bigint;

    declare vRoomBill2 bigint;
    declare vRoomBill3 bigint;
    declare vRoomBill8 bigint;
    declare vRoomBillFee2 bigint;
    declare vRoomBillFee3 bigint;
    declare vRoomBillFee8 bigint;

    set vTmToday = date(now());
    set vTmEnd = date_add(date(now()), interval hour(now()) hour);
    select tm into vTmBegin from z_huizong_qudao_hour where flavor = vFlavor order by tm desc limit 1;
    if vTmBegin is null then
        set vTmBegin = '2017-01-01';
    end if;

#成长任务奖励增加金币
    select sum(s) into vTaskCoin
    from (select uid, sum(coin_change) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type = 37
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#成长任务奖励增加3星龙珠
    select sum(s) into vTaskBall
    from (select uid, sum(variation) as s
          from user_value_props_log
          where time >= vTmBegin
            and time < vTmEnd
            and sendType = 37
            and type = 3
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;

#龙珠场抽宝箱增加金币
    select sum(s) into vLZCCoin
    from (select uid, sum(coin_change) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type = 65
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#龙珠场抽宝箱增加一星龙珠,
    select sum(s) into vLZCBall1
    from (select uid, sum(variation) as s
          from user_value_props_log
          where time >= vTmBegin
            and time < vTmEnd
            and sendType = 65
            and type = 1
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#龙珠场抽宝箱增加二星龙珠,
    select sum(s) into vLZCBall2
    from (select uid, sum(variation) as s
          from user_value_props_log
          where time >= vTmBegin
            and time < vTmEnd
            and sendType = 65
            and type = 2
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#龙珠场抽宝箱增加三星龙珠,
    select sum(s) into vLZCBall3
    from (select uid, sum(variation) as s
          from user_value_props_log
          where time >= vTmBegin
            and time < vTmEnd
            and sendType = 65
            and type = 3
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;

#活动奖励增加金币
    select sum(s) into vOtherCoin
    from (select uid, sum(coin_change) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type in (39, 76)
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#活动奖励增加3星龙珠
    select sum(s) into vOtherBall
    from (select uid, sum(variation) as s
          from user_value_props_log
          where time >= vTmBegin
            and time < vTmEnd
            and sendType = 62
            and type = 3
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;

#红包场流水
    select ifnull(sum(s), 0) into vRoomBill8
    from (select uid, sum(fee) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type in (2, 74)
            and room_id = 2
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#初级场流水
    select ifnull(sum(s), 0) into vRoomBill2
    from (select uid, sum(fee) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type in (2, 74)
            and room_id = 3
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;
#中级场流水
    select ifnull(sum(s), 0) into vRoomBill3
    from (select uid, sum(fee) as s
          from coin_log
          where add_time >= vTmBegin
            and add_time < vTmEnd
            and change_type in (2, 74)
            and room_id = 8
          group by uid) a
             inner join user on a.uid = user.uid
    where flavors = vFlavor;

    set vRoomBillFee2 = vRoomBill2 * 0.012;
    set vRoomBillFee3 = vRoomBill3 * 0.012;
    set vRoomBillFee8 = vRoomBill8 * 0.039;

    insert into z_huizong_qudao_hour(tm, flavor, in_cnt, new_cnt, check_cnt, money, hb, hbq_hbc, hbq_lzc,
                                     hbq_lzdb,
                                     hbq_other,
                                     coin_pay, coin_day, coin_vip, coin_card, coin_rank,
                                     coin_alms, coin_task, coin_mission, coin_mail, coin_key, coin_invite,
                                     coin_lzc, coin_yt, coin_fanli, coin_other,
                                     bill_room2, bill_room3, bill_room8,
                                     coin_fee_room2, coin_fee_room3, coin_fee_room8, coin_fee_fort)
        value (vTmEnd, vFlavor,
               ifnull((#今日登录人数 - 今日累计值
                          select count(distinct login_log.uid)
                          from login_log
                                   inner join user on login_log.uid = user.uid
                          where tm >= vTmToday
                            and io = 1
                            and user.flavors = vFlavor), 0),
               ifnull((#新注册玩家
                          select count(1)
                          from user
                          where reg_tm >= vTmBegin
                            and reg_tm < vTmEnd
                            and flavors = vFlavor), 0),
               ifnull((#签到次数
                          select count(1)
                          from awards_log
                                   inner join user on awards_log.uid = user.uid
                          where tm >= vTmBegin
                            and tm < vTmEnd
                            and awards_type = 0
                            and flavors = vFlavor), 0),
               ifnull((#充值金额
                          select sum(s)
                          from (select uid, sum(money) as s
                                from pay_log
                                where addtime >= vTmBegin
                                  and addtime < vTmEnd
                                  and result = 0
                                  and channel in (1, 2, 3, 6)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#红包金额 单位元
                          select sum(s)
                          from (select uid, sum(money) as s
                                from pay_log
                                where addtime >= vTmBegin
                                  and addtime < vTmEnd
                                  and result = 0
                                  and channel in (4, 5, 8)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#红包场红包券
                          select sum(s)
                          from (select uid, sum(variation) as s
                                from user_value_props_log
                                where time >= vTmBegin
                                  and time < vTmEnd
                                  and type = 40
                                  and sendType = 2
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#龙珠场红包券
                          select sum(s)
                          from (select uid, sum(variation) as s
                                from user_value_props_log
                                where time >= vTmBegin
                                  and time < vTmEnd
                                  and type = 40
                                  and sendType = 65
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#龙珠夺宝红包券
                          select sum(s)
                          from (select uid, sum(variation) as s
                                from user_value_props_log
                                where time >= vTmBegin
                                  and time < vTmEnd
                                  and type = 40
                                  and sendType = 66
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#集福活动,每日任务或者其他活动获得的红包券
                          select sum(s)
                          from (select uid, sum(variation) as s
                                from user_value_props_log
                                where time >= vTmBegin
                                  and time < vTmEnd
                                  and type = 40
                                  and sendType in (1, 62)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor), 0),
               ifnull((#充值增加金币(包含购买周月卡的一次性奖励金币)
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type in (7, 8, 14, 15, 23, 18, 46)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#签到增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 0
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#VIP升级+每日奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type in (29, 40)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#周月卡签到奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type in (45, 47)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#排行榜奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 26
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#破产奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 12
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#成长任务奖励增加金币+龙珠折算金币
                          ifnull(vTaskCoin, 0) + 1000000 * ifnull(vTaskBall, 0)
                          ), 0),
               ifnull((#在线/每日任务奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type in (1, 38)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#邮件奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 52
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#CDKey奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 17
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#邀请奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 55
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#龙珠场抽宝箱奖励增加金币
                              ifnull(vLZCCoin, 0) + 10000 * ifnull(vLZCBall1, 0) +
                              100000 * ifnull(vLZCBall1, 0) +
                              1000000 * ifnull(vLZCBall3, 0)
                          ), 0),
               ifnull((#鱼塘鱼货奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type in (69, 70)
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#金币返利奖励增加金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 77
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0),
               ifnull((#活动/升炮奖励增加金币+龙珠折算金币
                          ifnull(vOtherCoin, 0) + 1000000 * ifnull(vOtherBall, 0)
                          ), 0),
               vRoomBill2, vRoomBill3, vRoomBill8, vRoomBillFee2, vRoomBillFee3, vRoomBillFee8,
               ifnull((#购买炮台消耗金币
                          select sum(s)
                          from (select uid, sum(coin_change) as s
                                from coin_log
                                where add_time >= vTmBegin
                                  and add_time < vTmEnd
                                  and change_type = 34
                                group by uid) a
                                   inner join user on a.uid = user.uid
                          where flavors = vFlavor
                      ), 0)
        );
END;
-- ----------------------------
-- Procedure structure for `proc_huizong_by_flavor` END
-- ----------------------------


call proc_huizong_by_flavor('AOfficial');
call proc_huizong_by_flavor('MARKET_xw1');

DROP EVENT IF EXISTS event_huizong_by_hour;
CREATE EVENT event_huizong_by_hour
    ON SCHEDULE
        EVERY 24 hour
            STARTS TIMESTAMP '2019-01-01 18:30:00'
    ON COMPLETION PRESERVE
    COMMENT '每小时执行一次数据汇总'
    DO
    BEGIN
        call proc_huizong_by_flavor('MARKET_xw1');
        call proc_huizong_by_flavor('MARKET_xw2');
        call proc_huizong_by_flavor('MARKET_xw3');
        call proc_huizong_by_flavor('MARKET_xw4');
        call proc_huizong_by_flavor('MARKET_xw5');
        call proc_huizong_by_flavor('MARKET_xw6');
        call proc_huizong_by_flavor('MARKET_xw');
        call proc_huizong_by_flavor('MARKET_sp');
        call proc_huizong_by_flavor('AOfficial');
    END;






