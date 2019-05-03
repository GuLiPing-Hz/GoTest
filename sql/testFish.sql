SELECT *
FROM Buyu.user_stat
WHERE uid = 177861;
SELECT *
FROM Buyu.user
WHERE uid = 177861;
SELECT *
FROM Buyu.user_stat
WHERE uid = 177854;

SELECT *
FROM Buyu.yule_playerlog
WHERE uid = 177775;
SELECT SUM(changeCoin)
FROM Buyu.yule_playerlog
WHERE uid = 177775;
SELECT *
FROM user_stat
WHERE uid = 170652;
SELECT *
FROM user_stat
WHERE uid = 170662;
SELECT *
FROM lzdb_log
WHERE uid = 170652;
SELECT *
FROM lzdb_log
WHERE uid = 170662;

# UPDATE lb_points()
UPDATE user_stat
SET lb_points = ?
WHERE uid = ?;
UPDATE user_stat
SET lb_points = 10
WHERE uid = 100027;

#查询用户充值成功的订单
select *
from order_log
where (channel = 2 and (result = 0 or result = 3))
      and issandbox = 0 and addtime >= '2018-10-18 00:00:00' and addtime < '2018-11-06';

#查询用户领取破产记录的日志
SELECT CURDATE();
# count(1) as cnt
select count(1) as cnt
from alms_log
where uid = 170652 AND tm >= (curdate());
select count(1) as cnt
from alms_log
where uid = 170284 AND tm >= (curdate());
select *
from alms_log
where uid = 170652 AND DATE(tm) = (curdate());

SELECT *
FROM user
WHERE uid = 170652;
SELECT *
FROM user_stat
WHERE uid = 170652;
SELECT *
FROM user
WHERE uid = 177849;


SELECT *
FROM ip_rule
WHERE id > 9 OR tm > ADDDATE(NOW(), INTERVAL -2 MINUTE);
select *
from mail_reward
where id = 21338;

# call proc_select_collapse_count(165338,'445CE2DC-86E1-DA19-76BE-A493EB78EA78');
insert into alms_log (uid, uuid, login_ip, coin, tm, platform, exp, charge)
values (177863, '43BA4F0B-776F-F184-6AFE-FB5C6095EA90', '183.156.125.192'
  , 8000, '2018-11-30 12:53:20', '0', 107850, 0);

UPDATE mission
SET state = 1
WHERE uid = 177863 AND mid = 2 AND state = 0 AND value >= 60;
insert into notice (title, content, sender, receiver, addtime, isValid, showOrder, ntype, mail_type, mail_giftid)
values ('新手七天奖励', '新手七天大礼，奖励已送达，请签收。', 0, 177863, '2018-12-04 11:26:20', 1, 2, 1, 5, '7');
select
  reward_id,
  reward_cnt,
  operator
from mail_reward
where mail_giftid = 7 and isvalid = 1;

SELECT totalsum
FROM yule_gamelog
ORDER BY round DESC
limit 1;
update user_stat
set coin = 0
where uid in (
  select uid
  from user
  where type = 3
);


update user_stat
set coin = 18000
where uid = 165331;
delete from room_win_log;

UPDATE user_stat AS a, card_log AS b
SET a.daily_reward = a.daily_reward + 2
WHERE a.uid = b.uid AND b.end_tm > 0 AND b.wares_id = 'lailai.fish.thirtyday';

ALTER TABLE user_props_log
  MODIFY sendType smallint NOT NULL
  COMMENT '参见Trello。数据库字段说明';
ALTER TABLE coin_log
  MODIFY game_type tinyint DEFAULT 0
  COMMENT 'LOBBY = 0, -- 大厅
        FISH = 1, -- 捕鱼游戏
        OTHER = 2, -- 其他
        LOTTERY = 3, -- 抽奖
        GAME_YULE = 8, -- 鱼乐游戏
        GAME_SLOT = 9 -- 水浒传';

update dragoncard_use_log
set starttime = 0, deadtime = 0, updateTime = '2018-12-20 14:56:21'
where uid = 165617;
update user_stat
set coin = 2000
where uid between 166668 and 167250;
delete from user_exp
where uid between 166668 and 167250;
update user_stat
set multi_rate = 1;

select count(1)
from gift_pkg_log;
SELECT *
FROM growth_task_log
WHERE uid = 166847
ORDER BY task_id;

SELECT *
FROM portal_pay_order
WHERE oid = 144;
SELECT *
FROM portal_pay_order
WHERE orderId = 144;

select *
from user
where token = 'f162fe500f22f01415039cc50f218e10';

call xianwan_user(167348, "MARKET_xw");
call xianwan_user2(167348, "MARKET_xw");

select
  variation,
  count(variation)
from user_props_log
  inner join user on user.uid = user_props_log.uid
where flavors = 'MARKET_sp'
group by variation;

INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 0, 0, '2018-11-16 03:44:13', '2018-11-16 02:32:24', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 8, 7, '2018-11-16 03:45:03', '2018-11-16 02:33:15', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 10, 2, '2018-11-16 04:02:19', '2018-11-16 02:50:31', 1542398539000);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 1, 4, '2018-11-18 02:54:03', '2018-11-18 01:42:08', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 2, 4, '2018-11-19 18:15:30', '2018-11-19 18:15:32', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 9, 6, '2018-12-05 10:08:09', '2018-12-05 10:07:11', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 14, 2, '2018-12-05 12:19:18', '2018-12-05 12:18:20', 1544069958000);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 3, 4, '2018-12-17 19:38:47', '2018-12-17 19:37:02', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 4, 4, '2018-12-24 13:27:47', '2018-12-24 13:25:36', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 5, 4, '2018-12-30 12:34:59', '2018-12-30 12:32:26', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 13, 40, '2019-02-05 02:29:04', '2019-02-05 02:24:08', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 12, 62, '2019-02-05 21:04:54', '2019-02-05 20:59:58', 1559739894000);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 6, 40, '2019-02-17 02:53:37', '2019-02-17 02:53:28', -1);
INSERT INTO Buyu.fort_log (uid, fort_id, add_type, add_time, expire_time, expire_utc)
VALUES (165276, 0, 0, '2019-02-28 00:48:26', '2019-02-28 00:47:35', -1);

delete from room_win_log;
select days
from newer7_log
where uid = 197016 and DATE(tm) = DATE(DATE_SUB(NOW(), INTERVAL '1 0:0:0' DAY_SECOND));
CREATE INDEX user_stat_level_index
  ON user_stat (level DESC);
ALTER TABLE user
  MODIFY isrobot tinyint DEFAULT 0
  COMMENT '1机器人 0用户';

DROP INDEX user_stat_level_index
ON user_stat;
CREATE INDEX user_stat_level_index
  ON user_stat (level DESC);
CREATE INDEX user_stat_online_time_index
  ON user_stat (online_time DESC);

update user_stat
set level        = 0, vip = 0, log_cnt = 4, online_time = 51, game_cnt = 0, game_time = 0, money = 0, coin = 33300,
  agent_coin     = 0, pay_coin = 0, award_coin = 17800, giftin_coin = 0, giftout_coin = 0, pearl = 0, db1 = 0, db2 = 0,
  db3            = 0, lb = 0, hbq = 30,
  hb             = 0, exp = 4700, act = 0, dc = 0, froze = 0, `lock` = 0, double_gold = 0, double_speed = 0,
  auto_fire      = 0, lb_points = 0, lz_points = 0, lucky_flag = 0,
  fuzi           = 1, fuzi_total = 1, guide_flag = 1, multi_rate = 1, auto_card = 1, transfer_coin = 0,
  isgetvipreward = 0, first_pay = 0, diamond = 0
  , pay_diamond  = 0, award_diamond = 0, diamond_used = 0, fort_value = 50000, daily_reward = 0, daily_giftpkg = 0
where uid = 167371;
# Prepare error=Error 1054: Unknown column 'level' in 'field list'


insert into coin_log (uid, coin_before, coin_after, coin_change, add_time, fee, diamond_before, diamond_after, diamond_change, change_type, game_type, room_id, round, oid)
values (166533, 0, 180000, 180000, '2019-03-15 01:42:21', 0, 1239, 1239, 0, 24, 2, 0, 0, 0),
  (166762, 40700, 31900, 9400, '2019-03-15 01:42:31', 600, 1281, 1281, 0, 2, 1, 8, 0, 0),
  (166704, 180800, 175000, 11400, '2019-03-15 01:42:31', 5600, 1883, 1883, 0, 2, 1, 8, 0, 0),
  (166798, 192200, 191100, 1100, '2019-03-15 01:42:31', 0, 1458, 1458, 0, 2, 1, 8, 0, 0),
  (166967, 180000, 178300, 2700, '2019-03-15 01:42:31', 1000, 1704, 1704, 0, 2, 1, 2, 0, 0);

select *
from user
where uid = 167367;
select *
from user
where uid = 165668;
SELECT *
FROM portal_pay_order
WHERE orderId = 1147484200;
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484193, '2019-01-14 17:04:51', 188787, '201901140908226803773165', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 17:04:52', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484200, '2019-01-14 19:12:40', 188787, '201901141116128223258722', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:42', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484196, '2019-01-14 19:12:28', 188787, '201901141115597143766994', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:29', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484197, '2019-01-14 19:12:31', 188787, '201901141116031936825686', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:33', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484198, '2019-01-14 19:12:34', 188787, '201901141116052872730613', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:35', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484191, '2019-01-14 16:53:25', 188787, '201901140856576571003182', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 16:53:27', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484194, '2019-01-14 19:12:23', 188787, '201901141115547354131214', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:24', 'tx2.0', '兑换红包成功', 'tx2', 1);
INSERT INTO Buyu.portal_pay_order (orderId, createTime, userId, orderNo, price, status, payType, clientIP, money, payTime, payOrderNo, payInfo, waresId, complete)
VALUES (1147484195, '2019-01-14 19:12:26', 188787, '201901141115572740362686', -2, 1, 4, '115.210.65.179', 2,
                    '2019-01-14 19:12:27', 'tx2.0', '兑换红包成功', 'tx2', 1);


SELECT
  act,
  state
FROM active_coin
WHERE uid = 167367
ORDER BY act ASC;
delete from growth_task_log;
SELECT
  act,
  state
FROM active_coin
WHERE uid = 167367
ORDER BY act ASC;
select sum(money) as s
from pay_log
where uid = 167367 and result = 0 and channel in (2, 3);

insert into pay_log (tradeno, channel, uid, waresid, money, result, addtime, transtime)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);


select count(1) as cnt
from cdkey_log2
where id = 'Mv0ycH8lAiro6D5' and uid = 167367;
SELECT *
FROM growth_task_log
WHERE uid = 167367
ORDER BY task_id ASC;
delete from growth_task_log;

explain select
  l.uid,
  max(case ut
      when 0
        then total
      else 0 end) 上,
  max(case ut
      when 2
        then total
      else 0 end) 下,
  s.money,
  sum(variation)
from day_user_longzhu l
  left join user_stat s on s.uid = l.uid
  left join user_props_log p
    on p.uid = l.uid
where p.type = 40
      and p.sendType = 2
group by l.uid;

delete from yule_gamelog;

delete from coin_log;
select sum(fee) from coin_log;

