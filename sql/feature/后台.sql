ALTER TABLE buyu_summary ADD paoCoin1 bigint DEFAULT 0 NULL COMMENT '用户累计购买的炮台花费的金币';

DROP PROCEDURE IF EXISTS buyu_summary;
CREATE DEFINER =`root`@`192.168.100.3` PROCEDURE `buyu_summary`(in tmIn DATETIME)
  BEGIN

    declare v_twoHoursAgo datetime;
    declare v_chanchu bigint;
    declare v_room2flowing bigint;
    declare v_room3flowing bigint;
    declare v_room8flowing bigint;
    declare v_total_coin bigint;
    declare v_dbcoin bigint;
    declare v_getcoin bigint;
    declare v_rechargeCoin bigint;
    declare v_rewardCoin bigint;
    declare v_hbqCoin bigint;
    declare v_paoCoin bigint;
    declare v_paoCoin1 bigint;
    declare v_roomWin2 double;
    declare v_roomWin3 double;
    declare v_roomWin8 double;
    declare v_dragonball2 bigint;
    declare v_dragonball3 bigint;
    declare v_dragonball8 bigint;
    declare v_curBill2 double;
    declare v_curBill3 double;
    declare v_curBill8 double;
    -- 2小时
    /*金币总库存 = 玩家（排除机器人）财富（金币+珍珠+龙珠1,2,3）+ 4个水浒传库存+ 龙珠奖池库存（初级场，中级场，红包场）
    金币奖励 = 系统奖励(不含打鱼获得的奖励)+充值奖励
    预估消耗 = 房间当前流水 * 房间当前抽水平均比例（参见【房间抽水比例】卡片）
    捕鱼消耗(总计) = 房间累积输赢(roomWin)+预估消耗
    英雄传消耗(总计)=水浒传手续费(大、小)

    最终需要展示的值
    金币变化（金币奖励-金币消耗）+系统奖励金币+玩家充值金币+金币消耗(见上面)+玩家财富+渔乐+ 初级龙珠奖池+ 中级龙珠奖池+ 高级龙珠奖池+水浒6项

    按每小时记录捕鱼消耗
    时段捕鱼消耗 = 两个时段的捕鱼消耗相减。*/

    set v_twoHoursAgo = DATE_SUB(tmIn, INTERVAL 1 HOUR);

    -- 用户金币
    select
      ifnull(sum(s.coin), 0),
      ifnull(sum(get_coin), 0)
    into v_total_coin, v_getCoin
    from user_stat s
      left join user u on u.uid = s.uid
    where u.type in (0, 1, 2);

    -- 龙珠+珍珠转换金币+龙珠个人奖池
    select ifnull(sum(a + b + c + d + e), 0)
    into v_dbcoin
    from (select
            sum(pearl * 100)   a,
            sum(db1 * 10000)   b,
            sum(db2 * 100000)  c,
            sum(db3 * 1000000) d,
            sum(get_coin)      e
          from user_stat s
            left join user u on u.uid = s.uid
          where u.type in (0, 1, 2)) a;

    --  龙珠奖池库存（初级场，中级场，红包场）
    -- 房间当前流水 * 房间当前抽水平均比例初级场 10中级场 11红包场 39
    select
      ifnull(dragonball, 0),
      curBill
    into v_dragonball2, v_curBill2
    from room_win_log
    where rid = 2
    order by tm desc
    limit 1;

    select
      ifnull(dragonball, 0),
      curBill
    into v_dragonball3, v_curBill3
    from room_win_log
    where rid = 3
    order by tm desc
    limit 1;

    select
      ifnull(dragonball, 0),
      curBill
    into v_dragonball8, v_curBill8
    from room_win_log
    where rid = 8
    order by tm desc
    limit 1;

    -- 购买炮台金币
    select sum(coin_change)
    into v_paoCoin1
    from coin_log
    where change_type = 34;
    -- 充值金币
    select ifnull(sum(coin_change), 0)
    into v_rechargeCoin
    from coin_log
    where change_type in (7, 8, 9, 10, 18, 46);

    -- 奖励金币
    select ifnull(sum(coin_change), 0)
    into v_rewardCoin
    from coin_log
    where change_type in (0, 1, 6, 11, 12, 17, 19, 20, 26, 29, 37, 38, 39, 40, 41, 45, 47, 49, 50, 52, 53, 55, 62);

    -- 红包券兑换金币
    select ifnull(sum(coin_change), 0)
    into v_hbqCoin
    from coin_log
    where change_type = 60;

    -- 房间累积输赢(roomWin)
    select max(roomWin)
    into v_roomWin2
    from room_win_log
    where rid = 2;
    select max(roomWin)
    into v_roomWin3
    from room_win_log
    where rid = 3;
    select max(roomWin)
    into v_roomWin8
    from room_win_log
    where rid = 8;

    -- 时段流水
    select ifnull(sum(fee), 0)
    into v_room2flowing
    from coin_log
    where room_id = 2 and add_time > v_twoHoursAgo;
    select ifnull(sum(fee), 0)
    into v_room3flowing
    from coin_log
    where room_id = 3 and add_time > v_twoHoursAgo;
    select ifnull(sum(fee), 0)
    into v_room8flowing
    from coin_log
    where room_id = 8 and add_time > v_twoHoursAgo;

    select sum(coin_change)
    into v_paoCoin
    from coin_log
    where change_type = 34 and add_time > v_twoHoursAgo;

    -- 时段产出
    select ifnull(sum(coin_change), 0)
    into v_chanchu
    from coin_log
    where change_type in (0, 1, 6, 11, 12, 17, 19, 20, 26, 29, 37, 38, 39, 40, 41, 45, 47, 49, 50, 52, 53, 55, 62, #金币奖励
                          7, 8, 9, 10, 18, 46, #金币充值
                          60 #红包券兑换金币
    ) and add_time > v_twoHoursAgo;

    insert into buyu_summary (stats_time,
                              chanchu,
                              room2flowing,
                              room3flowing,
                              room8flowing,
                              total_coin,
                              dbcoin,
                              getCoin,
                              recharge_coin,
                              reward_coin,
                              hbqCoin,
                              paoCoin,
                              paoCoin1,
                              roomWin2,
                              curBill2,
                              dragonball2,
                              roomWin3,
                              curBill3,
                              dragonball3,
                              roomWin8,
                              curBill8,
                              dragonball8)
    values (tmIn,
      v_chanchu,
      v_room2flowing * 0.01,
      v_room3flowing * 0.011,
      v_room8flowing * 0.039,
      v_total_coin,
      v_dbcoin,
      v_getCoin,
      v_rechargeCoin,
      v_rewardCoin,
      v_hbqCoin,
      v_paoCoin,
      v_paoCoin1,
      v_roomWin2,
      v_curBill2 * 0.01,
      v_dragonball2,
      v_roomWin3,
      v_curBill3 * 0.011,
      v_dragonball3,
      v_roomWin8,
      v_curBill8 * 0.039,
            v_dragonball8);
  END;


call Buyu.buyu_summary('2019-04-15 16:00:00');


DROP EVENT IF EXISTS buyu_summary;
create definer = root@`192.168.100.3` event buyu_summary
  on schedule
    every '1' HOUR
      starts '2019-04-11 00:00:00'
  on completion preserve
  enable
do
  BEGIN
    CALL buyu_summary(now());
  END;


-- auto-generated definition
drop table if exists online_count;
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
CREATE INDEX online_count_addtime_index
  ON online_count (addtime);

call proc_get_today_flag(167374, '00110011', '2019-04-16', 0,1,0);
select count(1)
      from awards_log
      where awards_type = 45 and uid = 183270 AND tm >= '2019-04-18';

delete from room_win_log;
delete from coin_log;
delete from user_props_log;
delete from online_count;