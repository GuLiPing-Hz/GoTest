DELIMITER //
create procedure agent_promotion_reward()
BEGIN
    declare v_createTime varchar(10);
    declare v_uid int(11);
    declare v_num int(11);


    declare done int default false;
    declare my_cursor cursor for (select uid
                                  from user
                                  where type in (1, 2));
    declare continue handler for not found set done = true;

    select date_format(DATE_ADD(sysdate(), INTERVAL -1 DAY), '%Y-%m-%d') into v_createTime;

    open my_cursor;
    myLoop:
        loop
            fetch my_cursor
                into v_uid;
            if done
            then
                leave myLoop;
            end if;

            select count(1) into v_num
            from user
            where bind_uid = v_uid
              and DATEDIFF(bind_date, NOW()) = -1;

            insert into agent_promotion_reward (uid, num, createTime, isExamine)
            values (v_uid, v_num, v_createTime, 0);


        end loop myLoop;
    close my_cursor;


END//
DELIMITER ;

DELIMITER //
create procedure agent_recharge_reward()
BEGIN
    declare v_createTime varchar(10);
    declare v_uid int(11);
    declare v_selfReward int(11);
    declare v_lowerReward int(11);


    declare done int default false;
    declare my_cursor cursor for (select uid
                                  from user
                                  where type in (1, 2));
    declare continue handler for not found set done = true;

    select date_format(DATE_ADD(sysdate(), INTERVAL -1 DAY), '%Y-%m-%d') into v_createTime;

    open my_cursor;
    myLoop:
        loop
            fetch my_cursor
                into v_uid;
            if done
            then
                leave myLoop;
            end if;

            select ifnull(sum(coin_change), 0) into v_selfReward
            from coin_log l
                     left join user u on u.uid = l.uid
            where change_type = 3
              and u.bind_date < l.add_time
              and DATEDIFF(add_time, NOW()) = 0
              and u.uid = v_uid;

            select ifnull(sum(coin_change), 0) into v_lowerReward
            from coin_log l
                     left join user u on u.uid = l.uid
            where change_type = 3
              and u.bind_date < l.add_time
              and DATEDIFF(add_time, NOW()) = 0
              and u.bind_uid = v_uid;


            insert into agent_reward (uid, selfReward, lowerReward, createTime, isExamine)
            values (v_uid, v_selfReward, v_lowerReward, v_createTime, 0);


        end loop myLoop;
    close my_cursor;


END//
DELIMITER ;

DELIMITER //
create procedure day_coin()
BEGIN
    insert into day_coinSummary_log (uid, stats_date, coin_before, coin_after, coin_change, fee, change_type, game_type,
                                     room_id, round, oid, type)
        (
            -- 金币流入流出记录
            select a.uid,
                   a.add_time,
                   a.coin_before,
                   a.coin_after,
                   a.coin_change,
                   a.fee,
                   a.change_type,
                   a.game_type,
                   a.room_id,
                   a.round,
                   a.oid,
                   b.type
            from coin_log a,
                 user b
            where a.round > 0
              and a.change_type = 3
              and a.uid = b.uid
              and add_time >= ADDDATE(CURDATE(), INTERVAL -1 DAY)
              AND add_time < CURDATE()
        );
END//
DELIMITER ;

DELIMITER //
create procedure day_longzhu_log()
BEGIN
    insert into day_longzhu_log (stats_date, transfer_count, transfer_type, type)
        (
            select addTime,
                   transfer_count,
                   transfer_type,
                   type
            from (
                     -- 统计接收的记录。剔除代理之间的赠送。sender_uid 对应的用户type 不等于 recv_uid 对应的用户type。用户类型type = 2的记录为接收记录
                     select addTime,
                            sum(transfer_count) transfer_count,
                            transfer_type,
                            type,
                            sender_nick,
                            recv_nick
                     from (
                              select u.type,
                                     sender_uid,
                                     recv_uid,
                                     transfer_type,
                                     transfer_count,
                                     addTime,
                                     sender_nick,
                                     recv_nick
                              from dragon_transfer_log l
                                       left join user u
                                                 on u.uid = l.recv_uid
                                       left join user u2
                                                 on u2.uid = l.sender_uid
                              where u.type = 2
                                and u2.type != 2
                                and DATEDIFF(addTime, NOW()) = -1) a
                     group by transfer_type, addTime, type, sender_nick, recv_nick
                     union
                     -- 统计赠送的记录。用户类型 type = 0 的记录为赠送记录
                     select addTime,
                            sum(transfer_count) transfer_count,
                            transfer_type,
                            type,
                            sender_nick,
                            recv_nick
                     from (
                              select u.type,
                                     sender_uid,
                                     recv_uid,
                                     transfer_type,
                                     transfer_count,
                                     addTime,
                                     sender_nick,
                                     recv_nick
                              from dragon_transfer_log l
                                       left join user u
                                                 on u.uid = l.recv_uid
                                       left join user s
                                                 on s.uid = l.sender_uid
                              where u.type = 0
                                and s.type = 2
                                and DATEDIFF(addTime, NOW()) = -1) a
                     group by transfer_type, addTime, type, sender_nick, recv_nick) a
        );

END//
DELIMITER ;

DELIMITER //
create procedure day_pool_log()
BEGIN

    -- 每日税收
    insert into day_pool_log (stats_date, room_id, fee)


        (select date_sub(curdate(), interval 1 day) day,
                a.room_id,
                (a.fee - b.fee)                     fee
         from (select room_id,
                      fee
               from pool_log
               where to_days(tm) = to_days(now())
               limit 5) a,
              (select room_id,
                      fee
               from pool_log
               where datediff(tm, now()) = -1
               limit 5) b
         where a.room_id = b.room_id);

    -- 	(select b.tm,a.room_id,(a.fee-b.fee)fee from
    -- 	(select tm,room_id,fee from pool_log where to_days(tm) = to_days(now()) order by tm asc limit 4)a,
    -- 	(select tm,room_id,fee from pool_log where DATEDIFF(tm,NOW())=-1 order by tm asc limit 4)b
    -- 	where a.room_id = b.room_id);


END//
DELIMITER ;

DELIMITER //
create procedure day_summary_log()
BEGIN


    -- 运营数据分析，每日数据统计。现在改成new了
    declare v_stats_date varchar(10);

    declare v_stats_date1 varchar(10);

    declare v_flowing bigint;


    declare v_room1 bigint;


    declare v_room2 bigint;


    declare v_room3 bigint;


    declare v_room4 bigint;


    declare v_recharge bigint;


    declare v_rechargCount bigint;
    declare v_rechargeCoin bigint;


    declare v_taxroom1 double;


    declare v_taxroom2 double;


    declare v_taxroom3 double;


    declare v_taxroom4 double;


    declare v_tax double;

    declare v_out bigint;


    declare v_in bigint;


    declare v_reward bigint;


    declare v_dbout1 bigint;


    declare v_dbout2 bigint;


    declare v_dbout3 bigint;


    declare v_dbout bigint;


    declare v_dbin1 bigint;


    declare v_dbin2 bigint;


    declare v_dbin3 bigint;


    declare v_dbin bigint;


    declare v_fee double;

    -- 上一天


    select date(date_add(sysdate(), interval -1 day)) into v_stats_date;

    -- 今天

    select date(sysdate()) into v_stats_date1;


    select ifnull(sum(coin_change), 0) into v_flowing
    from coin_log
    where change_type = 2
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_room1
    from coin_log
    where room_id = 1
      and (add_time >= v_stats_date and add_time < v_stats_date1);


    select ifnull(sum(coin_change), 0) into v_room2
    from coin_log
    where room_id = 2
      and (add_time >= v_stats_date and add_time < v_stats_date1);


    select ifnull(sum(coin_change), 0) into v_room3
    from coin_log
    where room_id = 3
      and (add_time >= v_stats_date and add_time < v_stats_date1);


    select ifnull(sum(coin_change), 0) into v_room4
    from coin_log
    where room_id = 4
      and (add_time >= v_stats_date and add_time < v_stats_date1);


    select ifnull(count(distinct uid), 0) into v_rechargCount
    from order_log
    where result in (0, 1)
      and issandbox = 0
      and (transtime >= v_stats_date and transtime < v_stats_date1);


    select ifnull(sum(coin_change), 0) into v_rechargeCoin
    from coin_log
    where change_type in (7, 8, 9, 15)
      and (transtime >= v_stats_date and transtime < v_stats_date1);

    select ifnull(sum(money), 0) into v_recharge
    from order_log
    where result in (0, 1)
      and issandbox = 0
      and (transtime >= v_stats_date and transtime < v_stats_date1);

    -- 金币流出

    SELECT ifnull(SUM(coin_change), 0) into v_out
    FROM day_coinSummary_log
    WHERE `coin_change` < 0
      AND `type` = 0
      AND round IN (SELECT DISTINCT (round)
                    FROM `day_coinSummary_log`
                    WHERE `type` = 2
                      and stats_date >= v_stats_date
                      and stats_date < v_stats_date1);

    -- 金币流入

    SELECT ifnull(SUM(coin_change), 0) into v_in
    FROM day_coinSummary_log
    WHERE `coin_change` > 0
      AND `type` = 0
      AND round IN (SELECT DISTINCT (round)
                    FROM `day_coinSummary_log`
                    WHERE `type` = 2
                      and stats_date >= v_stats_date
                      and stats_date < v_stats_date1);

    -- 奖励
    select ifnull(sum(coin_change), 0) into v_reward
    from coin_log
    where change_type in (0, 1, 11, 12, 13, 16, 17, 19, 20)
      and add_time >= v_stats_date
      and add_time < v_stats_date1;

    -- 龙珠流出1

    select ifnull(sum(transfer_count) * 10000, 0) into v_dbout1
    from day_longzhu_log
    where transfer_type = 1
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出2

    select ifnull(sum(transfer_count) * 100000, 0) into v_dbout2
    from day_longzhu_log
    where transfer_type = 2
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出3

    select ifnull(sum(transfer_count) * 1000000, 0) into v_dbout3
    from day_longzhu_log
    where transfer_type = 3
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入1

    select ifnull(sum(transfer_count) * 10000, 0) into v_dbin1
    from day_longzhu_log
    where transfer_type = 1
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入2

    select ifnull(sum(transfer_count) * 100000, 0) into v_dbin2
    from day_longzhu_log
    where transfer_type = 2
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入3

    select ifnull(sum(transfer_count) * 1000000, 0) into v_dbin3
    from day_longzhu_log
    where transfer_type = 3
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 新手场服务费

    select ifnull(fee, 0) into v_taxroom1
    from day_pool_log
    where room_id = 1
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 初级场服务费

    select ifnull(fee, 0) into v_taxroom2
    from day_pool_log
    where room_id = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);
    -- 中级场服务费

    select ifnull(fee, 0) into v_taxroom3
    from day_pool_log
    where room_id = 3
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 高级场服务费

    select ifnull(fee, 0) into v_taxroom4
    from day_pool_log
    where room_id = 4
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- tax 总服务费 = v_taxroom1+v_taxroom2+v_taxroom3+v_taxroom4

    -- outcoin 金币流出

    -- incoin 金币流入

    -- outdb 龙珠流出 = v_dbout1+v_dbout2+v_dbout3

    -- indb 龙珠流入 = v_dbin1+v_dbin2+v_dbin3

    insert into day_summary_log (stats_date, flowing, room1, room2, room3, room4, recharge, rechargeCount, rechargeCoin,
                                 taxroom1, taxroom2, taxroom3, taxroom4, tax, outcoin, incoin, reward, outdb, indb)


    values (v_stats_date, v_flowing, v_room1, v_room2, v_room3, v_room4, v_recharge, v_rechargCount, v_rechargeCoin,
            v_taxroom1, v_taxroom2, v_taxroom3, v_taxroom4,
            (v_taxroom1 + v_taxroom2 + v_taxroom3 + v_taxroom4)
               , v_out, v_in, v_reward, (v_dbout1 + v_dbout2 + v_dbout3), (v_dbin1 + v_dbin2 + v_dbin3))


    on duplicate key update flowing       = v_flowing,
                            room1         = v_room1,
                            room2         = v_room2,
                            room3         = v_room3,
                            room4         = v_room4,
                            recharge      = v_recharge,
                            rechargeCount = v_rechargCount,
                            rechargeCoin  = v_rechargeCoin,


                            taxroom1      = v_taxroom1,
                            taxroom2      = v_taxroom2,
                            taxroom3      = v_taxroom3,
                            taxroom4      = v_taxroom4,
                            tax           = v_tax,


                            outcoin       = v_out,
                            incoin        = v_in,
                            reward        = v_reward,
                            outdb         = v_dbout,
                            indb          = v_dbin;


END//
DELIMITER ;

DELIMITER //
create procedure day_summary_log_new()
BEGIN


    -- 运营数据分析，每日数据统计新增玩家在线时间总计 玩家平均在线时间  当日新玩家在线时间总计 当日新玩家平均在线时间
    declare v_stats_date varchar(10);
    declare v_stats_date1 varchar(10);
    declare v_flowing bigint;
    declare v_room1 bigint;
    declare v_room2 bigint;
    declare v_room3 bigint;
    declare v_room4 bigint;
    declare v_recharge bigint;
    declare v_rechargCount bigint;
    declare v_rechargeCoin bigint;
    declare v_taxroom1 double;
    declare v_taxroom2 double;
    declare v_taxroom3 double;
    declare v_taxroom4 double;
    declare v_out bigint;
    declare v_in bigint;
    declare v_reward bigint;
    declare v_dbout1 bigint;
    declare v_dbout2 bigint;
    declare v_dbout3 bigint;
    declare v_dbin1 bigint;
    declare v_dbin2 bigint;
    declare v_dbin3 bigint;
    declare v_fee double;
    declare v_totaltm bigint;
    declare v_avgtm bigint;
    declare v_newtotaltm bigint;
    declare v_newavgtm bigint;
    declare yestoday datetime;
    declare v_newBillAve double;
    declare v_totalBillAve double;

    -- 上一天
    #     select date(date_add(sysdate(), interval -1 day))
    #     into v_stats_date;
    set v_stats_date = date(DATE_SUB(NOW(), INTERVAL '1 0:0:0' DAY_SECOND));
    -- 今天

    select date(sysdate()) into v_stats_date1;

    -- 玩家在线总时长 平均在线时长
    select sum(tm),
           avg(tm)
           into v_totaltm, v_avgtm
    from online_log
    where (addtime >= v_stats_date and addtime < v_stats_date1);
    -- 新玩家在线总时长 平均在线时长
    select sum(tm),
           avg(tm)
           into v_newtotaltm, v_newavgtm
    from online_log l
             left join user u on u.uid = l.uid
    where (reg_tm >= v_stats_date and reg_tm < v_stats_date1)
      and (addtime >= v_stats_date and addtime < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_flowing
    from coin_log
    where change_type = 2
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_room1
    from coin_log
    where room_id = 1
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_room2
    from coin_log
    where room_id = 2
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_room3
    from coin_log
    where room_id = 3
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0) into v_room4
    from coin_log
    where room_id = 4
      and (add_time >= v_stats_date and add_time < v_stats_date1);

    -- 充值金币数
    select ifnull(sum(coin_change), 0) into v_rechargeCoin
    from coin_log
    where change_type in (7, 8, 9, 15)
      and (add_time >= v_stats_date and add_time < v_stats_date1);
    -- select ifnull(count(distinct uid),0)into v_rechargCount from order_log where (channel=2 and (result = 0 or result = 3) or (channel = 3 and result = 0)) and issandbox =0 and (transtime >= v_stats_date and transtime < v_stats_date1);

    -- select ifnull(sum(money),0)into v_recharge from order_log where (channel=2 and (result = 0 or result = 3) or (channel = 3 and result = 0)) and issandbox =0 and (transtime >= v_stats_date and transtime < v_stats_date1);

    -- 充值金额和充值人数，这个数据是官方的充值，不含微信公众号充值
    select count(distinct uid),
           ifnull(sum(money), 0)
           into v_rechargCount, v_recharge
    from order_log
    where (channel = 2 and (result = 0 or result = 3) or (channel = 3 and result = 0))
      and issandbox = 0
      and (addtime >= v_stats_date and addtime < v_stats_date1);

    -- 金币流出剔除
    select ifnull(sum(coin_change), 0) into v_out
    from day_coinSummary_log
    where `coin_change` < 0
      and `type` = 0
      and round in (select distinct (round)
                    from `day_coinSummary_log`
                    where `type` = 2
                      and stats_date >= v_stats_date
                      and stats_date < v_stats_date1);

    -- 金币流入剔除
    select ifnull(sum(coin_change), 0) into v_in
    from day_coinSummary_log
    where `coin_change` > 0
      and `type` = 0
      and round in (select distinct (round)
                    from `day_coinSummary_log`
                    where `type` = 2
                      and stats_date >= v_stats_date
                      and stats_date < v_stats_date1);

    -- 奖励.对应的参考注释
    select ifnull(sum(coin_change), 0) into v_reward
    from coin_log
    where change_type in (0, 6, 1, 11, 12, 13, 17, 19, 20, 26, 29, 37, 38, 39, 40, 41, 49, 50, 52, 53)
      and add_time >= v_stats_date
      and add_time < v_stats_date1;

    -- 龙珠流出1
    select ifnull(sum(transfer_count) * 10000, 0) into v_dbout1
    from day_longzhu_log
    where transfer_type = 1
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出2
    select ifnull(sum(transfer_count) * 100000, 0) into v_dbout2
    from day_longzhu_log
    where transfer_type = 2
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出3
    select ifnull(sum(transfer_count) * 1000000, 0) into v_dbout3
    from day_longzhu_log
    where transfer_type = 3
      and type = 0
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入1
    select ifnull(sum(transfer_count) * 10000, 0) into v_dbin1
    from day_longzhu_log
    where transfer_type = 1
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入2
    select ifnull(sum(transfer_count) * 100000, 0) into v_dbin2
    from day_longzhu_log
    where transfer_type = 2
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入3
    select ifnull(sum(transfer_count) * 1000000, 0) into v_dbin3
    from day_longzhu_log
    where transfer_type = 3
      and type = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 新手场服务费
    select ifnull(fee, 0) into v_taxroom1
    from day_pool_log
    where room_id = 1
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 初级场服务费
    select ifnull(fee, 0) into v_taxroom2
    from day_pool_log
    where room_id = 2
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 中级场服务费
    select ifnull(fee, 0) into v_taxroom3
    from day_pool_log
    where room_id = 3
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 高级场服务费
    select ifnull(fee, 0) into v_taxroom4
    from day_pool_log
    where room_id = 4
      and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    #创建临时表，进行查询操作
    DROP TEMPORARY TABLE IF EXISTS tmp_t_calc_newer_bill_ave;
    CREATE TEMPORARY TABLE tmp_t_calc_newer_bill_ave
    select uid,
           sum(coin_change) as sum
    from Buyu.coin_log
    where change_type = 2
      and date(add_time) = v_stats_date
    group by uid;

    #查询到的结果
    #     CREATE TABLE IF NOT EXISTS day_summary_log_new
    #     insert into Buyu.day_summary_log_new (newbillave,totalbillave) values ()

    #所有用户平均流水
    select ifnull(avg(sum), 0) into v_totalBillAve
    FROM tmp_t_calc_newer_bill_ave;

    #新用户平均流水
    select ifnull(avg(sum), 0) into v_newBillAve
    FROM tmp_t_calc_newer_bill_ave
             inner join Buyu.user
    where tmp_t_calc_newer_bill_ave.uid = Buyu.user.uid
      and date(Buyu.user.reg_tm) = v_stats_date;

    -- 插入数据。其中只展示总服务费
    insert into day_summary_log_new ( stats_date, flowing, room1, room2, room3, room4, recharge, rechargeCount
                                    , rechargeCoin, taxroom1, taxroom2, taxroom3, taxroom4, outcoin, incoin, reward
                                    , outdb1, outdb2, outdb3
                                    , indb1, indb2, indb3, totaltm, avgtm, newtotaltm, newavgtm, newbillave
                                    , totalbillave)
    values (v_stats_date, v_flowing, v_room1, v_room2, v_room3, v_room4, v_recharge, v_rechargCount, v_rechargeCoin,
            v_taxroom1, v_taxroom2, v_taxroom3, v_taxroom4,
            v_out, v_in, v_reward, v_dbout1, v_dbout2, v_dbout3, v_dbin1, v_dbin2,
            v_dbin3, v_totaltm, v_avgtm, v_newtotaltm, v_newavgtm, v_newBillAve, v_totalBillAve)

    on duplicate key update flowing       = v_flowing,
                            room1         = v_room1,
                            room2         = v_room2,
                            room3         = v_room3,
                            room4         = v_room4,
                            recharge      = v_recharge,
                            rechargeCount = v_rechargCount,
                            rechargeCoin  = v_rechargeCoin,
                            taxroom1      = v_taxroom1,
                            taxroom2      = v_taxroom2,
                            taxroom3      = v_taxroom3,
                            taxroom4      = v_taxroom4,
                            outcoin       = v_out,
                            incoin        = v_in,
                            reward        = v_reward,
                            outdb1        = v_dbout1,
                            outdb2        = v_dbout2,
                            outdb3        = v_dbout3,
                            indb1         = v_dbin1,
                            indb2         = v_dbin2,
                            indb3         = v_dbin3,
                            totaltm       = v_totaltm,
                            avgtm         = v_avgtm,
                            newtotaltm    = v_newtotaltm,
                            newavgtm      = v_newavgtm,
                            newbillave    = v_newBillAve,
                            totalbillave  = v_totalBillAve;

    -- 删除临时表
    DROP TEMPORARY TABLE IF EXISTS tmp_t_calc_newer_bill_ave;

END//
DELIMITER ;

DELIMITER //
create procedure day_yule_summary()
BEGIN


    -- 鱼乐每日金币值
    declare v_stats_date varchar(10);

    declare v_stats_date1 varchar(10);

    declare v_coin bigint;

    select date(date_add(sysdate(), interval -1 day)) into v_stats_date;

    select date(sysdate()) into v_stats_date1;

    select sum(changeCoin) into v_coin
    from yule_gamelog
    where (addtime >= v_stats_date and addtime < v_stats_date1);

    insert into yule_summary (stats_date, coin)
    values (v_stats_date, v_coin)
    on duplicate key update stats_date = v_stats_date;

END//
DELIMITER ;

DELIMITER //
create procedure everyday_analysis()
BEGIN


    -- 首页龙珠奖池变化曲线.暂时未开启。系统中使用的是实时数据


    declare v_statsDate nvarchar(19);


    declare v_statsHour nvarchar(19);


    declare v_win1 double;


    declare v_win2 double;


    declare v_win3 double;

    -- declare v_dragonball1 double;


    -- declare v_dragonball2 double;


    -- declare v_dragonball3 double;


    select sysdate() into v_statsDate;


    select hour(sysdate()) into v_statsHour;


    select win into v_win1
    from pool_log
    where room_id = 1
    order by tm desc
    limit 1;


    select win into v_win2
    from pool_log
    where room_id = 2
    order by tm desc
    limit 1;


    select win into v_win3
    from pool_log
    where room_id = 3
    order by tm desc
    limit 1;

    -- 	select win into v_win4 from pool_log where room_id = 4 order by tm desc limit 1;


    -- select dragonball into v_dragonball1 from pool_log where room_id = 1 order by tm desc limit 1;


    -- 	select dragonball into v_dragonball2 from pool_log where room_id = 2 order by tm desc limit 1;


    -- 	select dragonball into v_dragonball3 from pool_log where room_id = 3 order by tm desc limit 1;


    -- 	select dragonball into v_dragonball4 from pool_log where room_id = 4 order by tm desc limit 1;


    if not exists(select 1
                  from stats_data_hour
                  where statsDate = v_statsDate
                    and statsHour = v_statsHour)
    then


        insert into stats_data_hour (statsDate, statsHour, win1, win2, win3)


        values (v_statsDate, v_statsHour, v_win1, v_win2, v_win3);


    else


        update stats_data_hour
        set statsDate = v_statsDate,
            statsHour = v_statsHour,
            win1      = v_win1,
            win2      = v_win2,
            win3      = v_win3


        where statsDate = v_statsDate
          and statsHour = v_statsHour;


    end if;


END//
DELIMITER ;

DELIMITER //
create procedure proc_select_collapse_count(IN uidIn bigint, IN uuidIn text)
BEGIN
    #Routine body goes here...
    declare cnt int;
    declare cnt1 int; #声明变量c
    declare cnt2 int; #声明变量c
    declare expMax bigint;
    declare chargeLast int;

    #查询uid领取破产记录
    select count(1),
           max(exp)
           into cnt1, expMax
    from Buyu.alms_log
    where uid = uidIn
      and date(tm) = current_date();

    select exp,
           charge
           into expMax, chargeLast
    from Buyu.alms_log
    where uid = uidIn
      and date(tm) = current_date()
    order by tm desc
    limit 1;

    #查询uuid领取破产记录
    select count(1) into cnt2
    from Buyu.alms_log
    where uuid = uuidIn
      and date(tm) = current_date();

    #判断哪个值最大
    if cnt1 > cnt2
    then
        set cnt = cnt1;
    else
        set cnt = cnt2;
    end if;

    if isnull(expMax)
    then
        set expMax = 0;
    end if;

    if isnull(chargeLast)
    then
        set chargeLast = 0;
    end if;

    #把最后的结果返回
    select cnt,
           expMax     as exp,
           chargeLast as charge;
END//
DELIMITER ;

DELIMITER //
create procedure proc_select_payexp_last(IN uidIn bigint)
BEGIN
    #Routine body goes here...
    select paycoin,
           exp
    from Buyu.pay_exp_log
    where uid = uidIn
      and date(time) = curdate()
    order by time desc
    limit 1;
END//
DELIMITER ;

DELIMITER //
create procedure proc_update_user_exp(IN uidIn bigint, IN expIn bigint)
BEGIN
    #Routine body goes here...
    declare today datetime;

    declare lastExp bigint;
    declare lastUpdate datetime;
    declare lastExp2 bigint;

    set today = now();
    select exp,
           exp2,
           update_time
           into lastExp, lastExp2, lastUpdate
    from Buyu.user_exp
    where uid = uidIn;

    if lastExp2 is null
    then
        set lastExp2 = lastExp;
    end if;

    if isnull(lastUpdate)
    then
        insert into Buyu.user_exp (uid, exp) values (uidIn, expIn);
        set lastExp2 = 0;
    elseif date(today) = date(lastUpdate)
    then
        update Buyu.user_exp
        set exp = expIn
        where uid = uidIn;
    else
        update Buyu.user_exp
        set exp  = expIn,
            exp2 = lastExp
        where uid = uidIn;
        set lastExp2 = lastExp;
    end if;

    select lastExp2 as exp2;
END//
DELIMITER ;

DELIMITER //
create procedure stat_remain_player()
BEGIN


    declare today DATE DEFAULT curdate();


    declare yesterday DATE DEFAULT date_sub(today, interval 1 day);


    declare days_ago_2 DATE DEFAULT date_sub(today, interval 2 day);


    declare days_ago_3 DATE DEFAULT date_sub(today, interval 3 day);


    declare days_ago_4 DATE DEFAULT date_sub(today, interval 4 day);


    declare days_ago_6 DATE DEFAULT date_sub(today, interval 6 day);


    declare days_ago_7 DATE DEFAULT date_sub(today, interval 7 day);
    declare days_ago_8 DATE DEFAULT date_sub(today, interval 8 day);
    declare days_ago_13 DATE DEFAULT date_sub(today, interval 13 DAY);
    declare days_ago_14 DATE DEFAULT date_sub(today, interval 14 DAY);
    declare days_ago_15 DATE DEFAULT date_sub(today, interval 15 DAY);
    declare days_ago_29 DATE DEFAULT date_sub(today, interval 29 DAY);
    declare days_ago_30 DATE DEFAULT date_sub(today, interval 30 DAY);
    declare days_ago_31 DATE DEFAULT date_sub(today, interval 31 DAY);

    -- 统计昨天注册人数


    insert into stats_remain (dru, stat_time, add_time)
    select count(uid),
           date_sub(curdate(), interval 1 day),
           now()
    from user
    where reg_tm between date_sub(curdate(), interval 1
                                  day) and curdate();

    -- 1日留存
    update stats_remain
    set second_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u on u.uid = l.uid and (u.reg_tm between days_ago_2 and yesterday) and
                                                     (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_2 and yesterday)
                   )
    )
    where stat_time = days_ago_2;

    -- 2日留存


    update stats_remain
    set third_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u on u.uid = l.uid and (u.reg_tm between days_ago_3 and days_ago_2) and
                                                     (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_3 and days_ago_2)
                   )
    )
    where stat_time = days_ago_3;

    -- 3日留存

    update stats_remain
    set fourth_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u on u.uid = l.uid and (u.reg_tm between days_ago_4 and days_ago_3) and
                                                     (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_4 and days_ago_3)
                   )
    )
    where stat_time = days_ago_4;

    -- 6日留存


    update stats_remain
    set seventh_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u on u.uid = l.uid and (u.reg_tm between days_ago_7 and days_ago_6) and
                                                     (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_7 and days_ago_6)
                   )
    )
    where stat_time = days_ago_7;

    -- 7日留存

    update stats_remain
    set eighth_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u on u.uid = l.uid and (u.reg_tm between days_ago_8 and days_ago_7) and
                                                     (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_8 and days_ago_7)
                   )
    )
    where stat_time = days_ago_8;

    -- 13日留存
    UPDATE stats_remain
    SET fourteen_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u
                                           on u.uid = l.uid and (u.reg_tm between days_ago_14 and days_ago_13) and
                                              (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_14 and days_ago_13)
                   )
    )
    where stat_time = days_ago_14;

    -- 14日留存
    UPDATE stats_remain
    SET fifteenth_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u
                                           on u.uid = l.uid and (u.reg_tm between days_ago_15 and days_ago_14) and
                                              (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_15 and days_ago_14)
                   )
    )
    where stat_time = days_ago_15;

    -- 29日留存
    UPDATE stats_remain
    SET thirtieth_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u
                                           on u.uid = l.uid and (u.reg_tm between days_ago_30 and days_ago_29) and
                                              (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_30 and days_ago_29)
                   )
    )
    where stat_time = days_ago_30;

    -- 30日留存
    UPDATE stats_remain
    SET thirtieth_first_day = (
        select (
                       (select count(distinct u.uid)
                        from login_log l
                                 left join user u
                                           on u.uid = l.uid and (u.reg_tm between days_ago_31 and days_ago_30) and
                                              (login_time between yesterday and today))
                       /
                       (select count(uid)
                        from user
                        where reg_tm between days_ago_31 and days_ago_30)
                   )
    )
    where stat_time = days_ago_31;

end;

DELIMITER //
create procedure stats_coin()
BEGIN


    declare v_stat_time varchar(19);


    declare v_total_coin bigint;


    declare v_pearl1 bigint;


    declare v_db1 bigint;


    declare v_db2 bigint;


    declare v_db3 bigint;


    declare v_dbcoin bigint;

    declare v_yulecoin bigint;

    declare v_stock_large bigint;

    declare v_stock_normal bigint;

    declare v_bibei_large bigint;

    declare v_bibei_normal bigint;

    declare v_fee_large bigint;

    declare v_fee_normal bigint;

    declare v_buyufee double;
    declare v_rechargeCoin bigint;
    declare v_rewardCoin bigint;
    declare v_win1 double;
    declare v_win2 double;
    declare v_win3 double;
    declare v_win4 double;
    declare v_pool1 double;
    declare v_pool2 double;
    declare v_pool3 double;
    declare v_pool4 double;
    declare v_fee1 double;
    declare v_fee2 double;
    declare v_fee3 double;
    declare v_fee4 double;
    declare v_dragonball1 double;
    declare v_dragonball2 double;
    declare v_dragonball3 double;
    declare v_dragonball4 double;

    select date_format(sysdate(), '%Y-%m-%d %H:%i:%s') into v_stat_time;


    select ifnull(sum(s.coin), 0) into v_total_coin
    from user_stat s
             left join user u on u.uid = s.uid
    where u.type != 3;

    select ifnull(sum(s.pearl), 0) into v_pearl1
    from user_stat s
             left join user u on u.uid = s.uid
    where u.type != 3;

    select ifnull(sum(s.db1), 0) into v_db1
    from user_stat s
             left join user u on u.uid = s.uid
    where u.type != 3;

    select ifnull(sum(s.db2), 0) into v_db2
    from user_stat s
             left join user u on u.uid = s.uid
    where u.type != 3;

    select ifnull(sum(s.db3), 0) into v_db3
    from user_stat s
             left join user u on u.uid = s.uid
    where u.type != 3;

    select ifnull(sum(a + b + c + d), 0) into v_dbcoin
    from (select sum(pearl * 100)   a,
                 sum(db1 * 10000)   b,
                 sum(db2 * 100000)  c,
                 sum(db3 * 1000000) d
          from user_stat s
                   left join user u on u.uid = s.uid
          where u.type != 3) a;

    -- 鱼乐库存
    select (sum(betsum) - sum(winsum)) into v_yulecoin
    from yule_playerlog
    where type = 0;

    -- 水浒传手续费
    select stock_large,
           stock_normal,
           bibei_large,
           bibei_normal,
           fee_large,
           fee_normal
           into v_stock_large, v_stock_normal, v_bibei_large, v_bibei_normal, v_fee_large, v_fee_normal
    from slot.stock;

    -- buyu手续费
    select ifnull(sum(fee), 0) into v_buyufee
    from (select fee
          from pool_log
          order by tm desc
          limit 5) a;
    -- 充值金币
    select ifnull(sum(coin_change), 0) into v_rechargeCoin
    from coin_log
    where change_type in (7, 8, 9, 15);
    -- 奖励金币
    select ifnull(sum(coin_change), 0) into v_rewardCoin
    from coin_log
    where change_type in (0, 6, 1, 11, 12, 13, 17, 19, 20, 26, 29, 37, 38, 39, 40, 41, 49, 50, 52, 53);

    -- 渔场的库存
    select win,
           pool,
           fee,
           dragonball
           into v_win1, v_pool1, v_fee1, v_dragonball1
    from (select *
          from pool_log
          order by tm desc
          limit 5) a
    where room_id = 1;
    select win,
           pool,
           fee,
           dragonball
           into v_win2, v_pool2, v_fee2, v_dragonball2
    from (select *
          from pool_log
          order by tm desc
          limit 5) a
    where room_id = 2;
    select win,
           pool,
           fee,
           dragonball
           into v_win3, v_pool3, v_fee3, v_dragonball3
    from (select *
          from pool_log
          order by tm desc
          limit 5) a
    where room_id = 3;
    select win,
           pool,
           fee,
           dragonball
           into v_win4, v_pool4, v_fee4, v_dragonball4
    from (select *
          from pool_log
          order by tm desc
          limit 5) a
    where room_id = 4;

    insert into stats_coin (stats_time, total_coin, pearl, db1, db2, db3, dbcoin, yulecoin, stock_large,
                            stock_normal, bibei_large, bibei_normal, fee_large, fee_normal, buyufee, recharge_coin,
                            reward_coin, win1, pool1, fee1, dragonball1, win2, pool2, fee2, dragonball2, win3, pool3,
                            fee3, dragonball3, win4, pool4, fee4, dragonball4)

    values (v_stat_time, v_total_coin, v_pearl1, v_db1, v_db2, v_db3, v_dbcoin, v_yulecoin, v_stock_large,
            v_stock_normal, v_bibei_large, v_bibei_normal, v_fee_large, v_fee_normal, v_buyufee,
            v_rechargeCoin, v_rewardCoin, v_win1, v_pool1, v_fee1,
            v_dragonball1, v_win2, v_pool2, v_fee2, v_dragonball2, v_win3,
            v_pool3, v_fee3, v_dragonball3, v_win4, v_pool4,
            v_fee4, v_dragonball4)


    on duplicate key update total_coin    = v_total_coin,
                            pearl         = v_pearl1,
                            db1           = v_db1,
                            db2           = v_db2,
                            db3           = v_db3,
                            dbcoin        = v_dbcoin,
                            yulecoin      = v_yulecoin,
                            stock_large   = v_stock_large,
                            stock_normal  = v_stock_normal,
                            bibei_large   = v_bibei_large,
                            bibei_normal  = v_bibei_normal,

                            fee_large     = v_fee_large,
                            fee_normal    = v_fee_normal,
                            buyufee       = v_buyufee,
                            recharge_coin = v_rechargeCoin,
                            reward_coin   = v_rewardCoin,
                            pool1         = v_pool1,
                            pool2         = v_pool2,
                            pool3         = v_pool3,
                            pool4         = v_pool4,
                            win1          = v_win1,
                            pool1         = v_pool1,
                            fee1          = v_fee1,
                            dragonball1   = v_dragonball1,
                            win2          = v_win2,
                            pool2         = v_pool2,
                            fee2          = v_fee2,
                            dragonball2   = v_dragonball2,
                            win3          = v_win3,
                            pool3         = v_pool3,
                            fee3          = v_fee3,
                            dragonball3   = v_dragonball3,
                            win4          = v_win4,
                            pool4         = v_pool4,
                            fee4          = v_fee4,
                            dragonball4   = v_dragonball4;


END//
DELIMITER ;

DELIMITER //
create procedure stats_summary()
BEGIN

    declare v_stats_date varchar(10);
    declare v_stats_date1 varchar(10);
    declare v_reg_user_count int;
    declare v_login_user_count int;
    declare v_max_online int;
    declare v_avg_online int;
    declare v_deposit_user_count int;
    declare v_depost_count int;
    declare v_pay_money int;
    declare v_total_coin bigint;
    declare v_total_exp bigint;

    -- 上一天
    select date(date_add(sysdate(), interval -1 day)) into v_stats_date;
    -- 今天
    select date(sysdate()) into v_stats_date1;

    select count(1) into v_reg_user_count
    from user
    where reg_tm >= v_stats_date
      and reg_tm < v_stats_date1;

    select count(distinct uid) into v_login_user_count
    from login_log
    where login_time >= v_stats_date
      and login_time < v_stats_date1;

    select ifnull(max(p_cnt), 0),
           ifnull(avg(p_cnt), 0)
           into v_max_online, v_avg_online
    from online_count
    where tms = 5
      and tp = 1
      and (addtime >= v_stats_date and addtime < v_stats_date1);

    -- 充值金额和充值人数，充值笔数，其中不含微信公众号
    select count(distinct uid),
           count(1),
           ifnull(sum(money), 0)
           into v_deposit_user_count, v_depost_count, v_pay_money
    from order_log
    where ((channel = 2 and (result = 0 or result = 3)) or (channel = 3 and result = 0))
      and issandbox = 0
      and addtime >= v_stats_date
      and addtime < v_stats_date1;

    select ifnull(sum(coin), 0),
           ifnull(sum(exp), 0)
           into v_total_coin, v_total_exp
    from user_stat;

    -- select ifnull(max(tm), 0), ifnull(avg(tm), 0) into v_max_online, v_avg_online from online_log where addtime >= v_stats_date and addtime < v_stats_date1;

    -- 今日流水 = 总流水 - 昨日流水.在16日凌晨执行下，计算出总流水。17日凌晨计算下，得出总流水。两者减 16流水-15流水= 16日流水
    set v_total_exp = v_total_exp - (select ifnull(total_exp, 0) total_exp
                                     from stats_summary
                                     where stats_date = date(date_add(sysdate(), interval -2 day)));

    -- select ifnull(sum(exp), 0) into v_total_exp from user_stat;
    --
    select ifnull(sum(exp), 0) into v_total_exp
    from user_stat
    where date(update_time) = v_stats_date;

    insert into stats_summary (stats_date, reg_user_count, login_user_count, max_online, avg_online, deposit_user_count,
                               depost_count, pay_money, total_coin, total_exp)
    values (v_stats_date, v_reg_user_count, v_login_user_count, v_max_online, v_avg_online, v_deposit_user_count,
            v_depost_count, v_pay_money, v_total_coin, v_total_exp)
    on duplicate key update reg_user_count     = v_reg_user_count,
                            login_user_count   = v_login_user_count,
                            max_online         = v_max_online,
                            avg_online         = v_avg_online,
                            deposit_user_count = v_deposit_user_count,
                            depost_count       = v_depost_count,
                            pay_money          = v_pay_money,
                            total_coin         = v_total_coin,
                            total_exp          = v_total_exp;

    -- 用户流水
    -- insert into stats_user_coin(stats_date, uid, coin) select now(), uid, sum(coin_change) from coin_log where date(add_time) = v_stats_date group by uid;

END//
DELIMITER ;

DELIMITER //
create procedure stats_user_coin()
BEGIN
    -- 用户流水
    insert into stats_user_coin (stats_date, uid, coin)
    select now(),
           uid,
           sum(coin_change)
    from coin_log
    group by uid;
END//
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_reset_by_day` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_reset_by_day;
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# DEFINER指定权限的存储过程
# CREATE DEFINER =`root`@`localhost` PROCEDURE `proc_reset_by_day`(IN a int, IN b int, OUT sum int)
CREATE PROCEDURE proc_reset_by_day(in tm DATETIME, in tm_utc bigint, in vip_limit int)
BEGIN
    #Routine body goes here...
    #     DECLARE sumTotal INT; #声明变量sumTotal
    #每日0点清空在线日志
    call online_log_summary(tm); #数据统计。
    DELETE FROM online_log2 where true;
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
        transfer_coin = 0 where true;

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
        state = 0 where true;

    #更新活跃度金币
    UPDATE active_coin
    SET state     = 0,
        update_tm = tm where true;

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
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_reset_by_day` END
-- ----------------------------
# call proc_reset_by_day('2019-03-14',1552521600000,4);

-- ----------------------------
-- Procedure structure for `proc_get_today_flag` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_get_today_flag;
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# DEFINER指定权限的存储过程
# CREATE DEFINER =`root`@`localhost` PROCEDURE `proc_reset_by_day`(IN a int, IN b int, OUT sum int)
CREATE PROCEDURE proc_get_today_flag(in uidIn bigint, in uuidIn text, in tmIn DATETIME, in check_week int,
                                     in check_month int, in check_vip int)
BEGIN

    declare day_uuid, day_uid, day, week, month, vip int;
    set week = 1, month = 1, vip = 1;

    #检查每日奖励领取 -uuid
    select count(1) into day_uuid
    from newer7_log
    where uuid = uuidIn
      AND tm >= tmIn;
    #检查每日奖励领取 -uid
    select count(1) into day_uid
    from newer7_log
    where uid = uidIn
      AND tm >= tmIn;

    if day_uuid > day_uid
    then
        set day = day_uuid;
    else
        set day = day_uid;
    end if;

    #检查周卡奖励领取
    if check_week = 1
    then
        select count(1) into week
        from awards_log
        where awards_type = 45
          and uid = uidIn
          AND tm >= tmIn;
    end if;

    #检查月卡奖励领取
    if check_month = 1
    then
        select count(1) into month
        from awards_log
        where awards_type = 47
          and uid = uidIn
          AND tm >= tmIn;
    end if;

    #vip检查每日奖励领取
    if check_vip = 1
    then
        select count(1) into vip
        from awards_log
        where awards_type = 29
          and uid = uidIn
          AND tm >= tmIn;
    end if;

    select day,
           week,
           month,
           vip;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_get_today_flag` END
-- ----------------------------
call proc_get_today_flag(167371, 'ABBA1969-3E3A-84D9-39CA-A685A9476255', '2019-03-13', 1, 1, 0);