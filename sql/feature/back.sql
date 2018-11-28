ALTER TABLE day_summary_log_new
  ADD newbillave double DEFAULT 0 NULL
COMMENT '新用户单日平均流水';
ALTER TABLE day_summary_log_new
  ADD totalbillave double DEFAULT 0 NULL
COMMENT '所有用户单日平均流水';

-- ----------------------------
-- #每日数据统计 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS day_summary_log_new;
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
    declare newBillAve double;
    declare totalBillAve double;

    -- 上一天
    #     select date(date_add(sysdate(), interval -1 day))
    #     into v_stats_date;
    set v_stats_date = date(DATE_SUB(NOW(), INTERVAL '1 0:0:0' DAY_SECOND));
    -- 今天

    select date(sysdate())
    into v_stats_date1;

    -- 玩家在线总时长 平均在线时长
    select
      sum(tm),
      avg(tm)
    into v_totaltm, v_avgtm
    from online_log
    where (addtime >= v_stats_date and addtime < v_stats_date1);
    -- 新玩家在线总时长 平均在线时长
    select
      sum(tm),
      avg(tm)
    into v_newtotaltm, v_newavgtm
    from online_log l left join user u on u.uid = l.uid
    where (reg_tm >= v_stats_date and reg_tm < v_stats_date1) and (addtime >= v_stats_date and addtime < v_stats_date1);

    select ifnull(sum(coin_change), 0)
    into v_flowing
    from coin_log
    where change_type = 2 and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0)
    into v_room1
    from coin_log
    where room_id = 1 and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0)
    into v_room2
    from coin_log
    where room_id = 2 and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0)
    into v_room3
    from coin_log
    where room_id = 3 and (add_time >= v_stats_date and add_time < v_stats_date1);

    select ifnull(sum(coin_change), 0)
    into v_room4
    from coin_log
    where room_id = 4 and (add_time >= v_stats_date and add_time < v_stats_date1);

    -- 充值金币数
    select ifnull(sum(coin_change), 0)
    into v_rechargeCoin
    from coin_log
    where change_type in (7, 8, 9, 15) and (add_time >= v_stats_date and add_time < v_stats_date1);
    -- select ifnull(count(distinct uid),0)into v_rechargCount from order_log where (channel=2 and (result = 0 or result = 3) or (channel = 3 and result = 0)) and issandbox =0 and (transtime >= v_stats_date and transtime < v_stats_date1);

    -- select ifnull(sum(money),0)into v_recharge from order_log where (channel=2 and (result = 0 or result = 3) or (channel = 3 and result = 0)) and issandbox =0 and (transtime >= v_stats_date and transtime < v_stats_date1);

    -- 充值金额和充值人数，这个数据是官方的充值，不含微信公众号充值
    select
      count(distinct uid),
      ifnull(sum(money), 0)
    into v_rechargCount, v_recharge
    from order_log
    where (channel = 2 and (result = 0 or result = 3) or (channel = 3 and result = 0)) and issandbox = 0 and
          (addtime >= v_stats_date and addtime < v_stats_date1);

    -- 金币流出剔除
    select ifnull(sum(coin_change), 0)
    into v_out
    from day_coinSummary_log
    where `coin_change` < 0 and `type` = 0 and round in (select distinct (round)
                                                         from `day_coinSummary_log`
                                                         where `type` = 2 and stats_date >= v_stats_date and
                                                               stats_date < v_stats_date1);

    -- 金币流入剔除
    select ifnull(sum(coin_change), 0)
    into v_in
    from day_coinSummary_log
    where `coin_change` > 0 and `type` = 0 and round in (select distinct (round)
                                                         from `day_coinSummary_log`
                                                         where `type` = 2 and stats_date >= v_stats_date and
                                                               stats_date < v_stats_date1);

    -- 奖励.对应的参考注释
    select ifnull(sum(coin_change), 0)
    into v_reward
    from coin_log
    where change_type in (0, 6, 1, 11, 12, 13, 17, 19, 20, 26, 29, 37, 38, 39, 40, 41, 49, 50, 52, 53) and
          add_time >= v_stats_date and add_time < v_stats_date1;

    -- 龙珠流出1
    select ifnull(sum(transfer_count) * 10000, 0)
    into v_dbout1
    from day_longzhu_log
    where transfer_type = 1 and type = 0 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出2
    select ifnull(sum(transfer_count) * 100000, 0)
    into v_dbout2
    from day_longzhu_log
    where transfer_type = 2 and type = 0 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流出3
    select ifnull(sum(transfer_count) * 1000000, 0)
    into v_dbout3
    from day_longzhu_log
    where transfer_type = 3 and type = 0 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入1
    select ifnull(sum(transfer_count) * 10000, 0)
    into v_dbin1
    from day_longzhu_log
    where transfer_type = 1 and type = 2 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入2
    select ifnull(sum(transfer_count) * 100000, 0)
    into v_dbin2
    from day_longzhu_log
    where transfer_type = 2 and type = 2 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 龙珠流入3
    select ifnull(sum(transfer_count) * 1000000, 0)
    into v_dbin3
    from day_longzhu_log
    where transfer_type = 3 and type = 2 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 新手场服务费
    select ifnull(fee, 0)
    into v_taxroom1
    from day_pool_log
    where room_id = 1 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 初级场服务费
    select ifnull(fee, 0)
    into v_taxroom2
    from day_pool_log
    where room_id = 2 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 中级场服务费
    select ifnull(fee, 0)
    into v_taxroom3
    from day_pool_log
    where room_id = 3 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    -- 高级场服务费
    select ifnull(fee, 0)
    into v_taxroom4
    from day_pool_log
    where room_id = 4 and (stats_date >= v_stats_date and stats_date < v_stats_date1);

    #创建临时表，进行查询操作
    DROP TEMPORARY TABLE IF EXISTS tmp_t_calc_newer_bill_ave;
    CREATE TEMPORARY TABLE tmp_t_calc_newer_bill_ave
        select
          uid,
          sum(coin_change) as sum
        from Buyu.coin_log
        where change_type = 2 and date(add_time) = v_stats_date
        group by uid;

    #查询到的结果
    #     CREATE TABLE IF NOT EXISTS day_summary_log_new
    #     insert into Buyu.day_summary_log_new (newbillave,totalbillave) values ()

    #所有用户平均流水
    select ifnull(avg(sum), 0)
    into totalBillAve
    FROM tmp_t_calc_newer_bill_ave;

    #新用户平均流水
    select ifnull(avg(sum), 0)
    into newBillAve
    FROM tmp_t_calc_newer_bill_ave
      inner join Buyu.user
    where tmp_t_calc_newer_bill_ave.uid = Buyu.user.uid
          and date(Buyu.user.reg_tm) = v_stats_date;

    -- 插入数据。其中只展示总服务费
    insert into day_summary_log_new (stats_date, flowing, room1, room2, room3, room4, recharge, rechargeCount
      , rechargeCoin, taxroom1, taxroom2, taxroom3, taxroom4, outcoin, incoin, reward, outdb1, outdb2, outdb3
      , indb1, indb2, indb3, totaltm, avgtm, newtotaltm, newavgtm, newbillave, totalbillave)
    values (v_stats_date, v_flowing, v_room1, v_room2, v_room3, v_room4, v_recharge, v_rechargCount, v_rechargeCoin,
                          v_taxroom1, v_taxroom2, v_taxroom3, v_taxroom4,
                                                  v_out, v_in, v_reward, v_dbout1, v_dbout2, v_dbout3, v_dbin1, v_dbin2,
            v_dbin3, v_totaltm, v_avgtm, v_newtotaltm, v_newavgtm, newBillAve, totalBillAve)

    on duplicate key update flowing = v_flowing, room1 = v_room1, room2 = v_room2, room3 = v_room3, room4 = v_room4,
      recharge                      = v_recharge, rechargeCount = v_rechargCount, rechargeCoin = v_rechargeCoin,
      taxroom1                      = v_taxroom1, taxroom2 = v_taxroom2, taxroom3 = v_taxroom3, taxroom4 = v_taxroom4,
      outcoin                       = v_out, incoin = v_in, reward = v_reward, outdb1 = v_dbout1, outdb2 = v_dbout2,
      outdb3                        = v_dbout3, indb1 = v_dbin1, indb2 = v_dbin2, indb3 = v_dbin3, totaltm = v_totaltm,
      avgtm                         = v_avgtm, newtotaltm = v_newtotaltm, newavgtm = v_newavgtm;

    -- 删除临时表
    DROP TEMPORARY TABLE IF EXISTS tmp_t_calc_newer_bill_ave;

  END
//
DELIMITER ;
-- ----------------------------
-- #每日数据统计 end
-- ----------------------------

call day_summary_log_new();

-- ----------------------------
-- #每日数据统计定时任务 begin
-- ----------------------------
DROP EVENT IF EXISTS e_test;
DELIMITER //
CREATE EVENT e_test
  ON SCHEDULE EVERY 10 second
    STARTS TIMESTAMP '2018-11-28 17:06:00'
  ON COMPLETION PRESERVE
DO
  BEGIN
    select 1;
    CALL day_summary_log_new();
  END
//
DELIMITER ;

#查看event是否开启
show variables like '%event_sche%';
#开启event
set global event_scheduler = 1;
#关闭event
set global event_scheduler = 0;
#关闭事件任务 :
ALTER EVENT e_test
ON COMPLETION PRESERVE
DISABLE;
#开启事件任务 :
ALTER EVENT e_test
ON COMPLETION PRESERVE
ENABLE;
show events;

# SELECT now();
-- ----------------------------
-- #每日数据统计定时任务 end
-- ----------------------------

