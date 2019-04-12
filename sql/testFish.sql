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
  (166967, 180000, 178300, 2700, '2019-03-15 01:42:31', 1000, 1704, 1704, 0, 2, 1, 2, 0, 0),
  (166673, 724500, 723900, 3200, '2019-03-15 01:42:31', 2600, 1162, 1162, 0, 2, 1, 8, 0, 0),
  (167136, 915500, 913400, 13600, '2019-03-15 01:42:31', 11500, 1132, 1132, 0, 2, 1, 8, 0, 0),
  (167069, 491600, 490900, 1300, '2019-03-15 01:42:31', 600, 1549, 1549, 0, 2, 1, 8, 0, 0),
  (166731, 231100, 219600, 16100, '2019-03-15 01:42:31', 4600, 1417, 1417, 0, 2, 1, 8, 0, 0),
  (167041, 145100, 149900, 3000, '2019-03-15 01:42:31', 7800, 1196, 1197, 1, 2, 1, 8, 0, 0),
  (167085, 14300, 6500, 9000, '2019-03-15 01:42:31', 1200, 1412, 1412, 0, 2, 1, 8, 0, 0),
  (166854, 300, 180300, 180000, '2019-03-15 01:42:38', 0, 1557, 1557, 0, 24, 2, 0, 0, 0),
  (167005, 700, 180700, 180000, '2019-03-15 01:42:41', 0, 1272, 1272, 0, 24, 2, 0, 0, 0),
  (166917, 134900, 122300, 12600, '2019-03-15 01:42:42', 0, 1316, 1316, 0, 2, 1, 8, 0, 0),
  (167148, 681100, 660000, 52400, '2019-03-15 01:42:42', 31300, 1156, 1156, 0, 2, 1, 8, 0, 0),
  (166694, 180700, 172000, 12100, '2019-03-15 01:42:42', 3400, 1521, 1521, 0, 2, 1, 8, 0, 0),
  (167070, 139400, 146100, 17200, '2019-03-15 01:42:42', 23900, 1130, 1130, 0, 2, 1, 8, 0, 0),
  (167080, 449100, 442900, 7000, '2019-03-15 01:42:42', 800, 1218, 1218, 0, 2, 1, 8, 0, 0),
  (166711, 64800, 65000, 50800, '2019-03-15 01:42:42', 51000, 1044, 1044, 0, 2, 1, 8, 0, 0),
  (166847, 180100, 167900, 31300, '2019-03-15 01:42:42', 19100, 1301, 1301, 0, 2, 1, 8, 0, 0),
  (166671, 161000, 205700, 16000, '2019-03-15 01:42:42', 60700, 1525, 1525, 0, 2, 1, 8, 0, 0),
  (166859, 172500, 173500, 15600, '2019-03-15 01:42:42', 16600, 1442, 1442, 0, 2, 1, 8, 0, 0),
  (166845, 243100, 248400, 32500, '2019-03-15 01:42:42', 37800, 1431, 1431, 0, 2, 1, 8, 0, 0),
  (166392, 217600, 204700, 54300, '2019-03-15 01:42:44', 41400, 1259, 1259, 0, 2, 1, 2, 0, 0),
  (166647, 181300, 179200, 8000, '2019-03-15 01:42:44', 5900, 1165, 1165, 0, 2, 1, 2, 0, 0),
  (166366, 980300, 955900, 52200, '2019-03-15 01:42:44', 27800, 1057, 1057, 0, 2, 1, 2, 0, 0),
  (166397, 316500, 315500, 24600, '2019-03-15 01:42:44', 23600, 1048, 1048, 0, 2, 1, 2, 0, 0),
  (166613, 79500, 76400, 5900, '2019-03-15 01:42:44', 2800, 1700, 1700, 0, 2, 1, 2, 0, 0),
  (166576, 212100, 194900, 24000, '2019-03-15 01:42:44', 6800, 1454, 1454, 0, 2, 1, 2, 0, 0),
  (166489, 223900, 236700, 21000, '2019-03-15 01:42:44', 33800, 1186, 1186, 0, 2, 1, 2, 0, 0),
  (166587, 288900, 254100, 50800, '2019-03-15 01:42:44', 16000, 1180, 1180, 0, 2, 1, 2, 0, 0),
  (166365, 975800, 929700, 75600, '2019-03-15 01:42:44', 29500, 900, 900, 0, 2, 1, 2, 0, 0),
  (166518, 105500, 89800, 25300, '2019-03-15 01:42:44', 9600, 1255, 1255, 0, 2, 1, 2, 0, 0),
  (166502, 240800, 256600, 22600, '2019-03-15 01:42:44', 38400, 1126, 1126, 0, 2, 1, 2, 0, 0),
  (166604, 141600, 132200, 38500, '2019-03-15 01:42:44', 29100, 1099, 1099, 0, 2, 1, 2, 0, 0),
  (166460, 113900, 107400, 16100, '2019-03-15 01:42:44', 9600, 1171, 1171, 0, 2, 1, 2, 0, 0),
  (166551, 299200, 263800, 45200, '2019-03-15 01:42:44', 9800, 1043, 1043, 0, 2, 1, 2, 0, 0),
  (166445, 263500, 266100, 63900, '2019-03-15 01:42:44', 66500, 1292, 1292, 0, 2, 1, 2, 0, 0),
  (166493, 184000, 153200, 43200, '2019-03-15 01:42:44', 12400, 1378, 1378, 0, 2, 1, 2, 0, 0),
  (166491, 547500, 555100, 9300, '2019-03-15 01:42:44', 16900, 1474, 1474, 0, 2, 1, 2, 0, 0),
  (166642, 108300, 77700, 42300, '2019-03-15 01:42:44', 11700, 1341, 1341, 0, 2, 1, 2, 0, 0),
  (166607, 170400, 160200, 24600, '2019-03-15 01:42:44', 14400, 1462, 1462, 0, 2, 1, 2, 0, 0),
  (166432, 209900, 240600, 39900, '2019-03-15 01:42:44', 70600, 1353, 1353, 0, 2, 1, 2, 0, 0),
  (166442, 165000, 166900, 10900, '2019-03-15 01:42:44', 12800, 1227, 1227, 0, 2, 1, 2, 0, 0),
  (166625, 75500, 80000, 7500, '2019-03-15 01:42:44', 12000, 1045, 1045, 0, 2, 1, 2, 0, 0),
  (166605, 535100, 630900, 34200, '2019-03-15 01:42:44', 130000, 1053, 1053, 0, 2, 1, 2, 0, 0),
  (166418, 45600, 39600, 29500, '2019-03-15 01:42:44', 23500, 1197, 1197, 0, 2, 1, 2, 0, 0),
  (166372, 175400, 157600, 35700, '2019-03-15 01:42:44', 17900, 1039, 1039, 0, 2, 1, 2, 0, 0),
  (166477, 47000, 51000, 27000, '2019-03-15 01:42:44', 31000, 1091, 1091, 0, 2, 1, 2, 0, 0),
  (166406, 238300, 248200, 37000, '2019-03-15 01:42:44', 46900, 1440, 1440, 0, 2, 1, 2, 0, 0),
  (166413, 595100, 592500, 44800, '2019-03-15 01:42:44', 42200, 1540, 1540, 0, 2, 1, 2, 0, 0),
  (166577, 127200, 174800, 39200, '2019-03-15 01:42:44', 86800, 1448, 1448, 0, 2, 1, 2, 0, 0),
  (166597, 231500, 221600, 11300, '2019-03-15 01:42:44', 1400, 862, 862, 0, 2, 1, 2, 0, 0),
  (166439, 481400, 461900, 36300, '2019-03-15 01:42:44', 16800, 1527, 1527, 0, 2, 1, 2, 0, 0),
  (166648, 224600, 222300, 17700, '2019-03-15 01:42:44', 15400, 1674, 1674, 0, 2, 1, 2, 0, 0),
  (166405, 279700, 269500, 33000, '2019-03-15 01:42:44', 22800, 1501, 1501, 0, 2, 1, 2, 0, 0),
  (166638, 484300, 484800, 13100, '2019-03-15 01:42:44', 13600, 1679, 1679, 0, 2, 1, 2, 0, 0),
  (166601, 22700, 18700, 12700, '2019-03-15 01:42:44', 8700, 1350, 1350, 0, 2, 1, 2, 0, 0),
  (166573, 356600, 404400, 23100, '2019-03-15 01:42:44', 70900, 1376, 1376, 0, 2, 1, 2, 0, 0),
  (166379, 228300, 248600, 56700, '2019-03-15 01:42:44', 77000, 1242, 1242, 0, 2, 1, 2, 0, 0),
  (166428, 168500, 159000, 39300, '2019-03-15 01:42:44', 29800, 1166, 1166, 0, 2, 1, 2, 0, 0),
  (166590, 174000, 170700, 24000, '2019-03-15 01:42:44', 20700, 1302, 1302, 0, 2, 1, 2, 0, 0),
  (166454, 169800, 214000, 60400, '2019-03-15 01:42:44', 104600, 981, 981, 0, 2, 1, 2, 0, 0),
  (166462, 216000, 211500, 49500, '2019-03-15 01:42:44', 45000, 1069, 1070, 1, 2, 1, 2, 0, 0),
  (166569, 449700, 445800, 6100, '2019-03-15 01:42:44', 2200, 1558, 1558, 0, 2, 1, 2, 0, 0),
  (166466, 193000, 191400, 21400, '2019-03-15 01:42:44', 19800, 1129, 1129, 0, 2, 1, 2, 0, 0),
  (166470, 118700, 138300, 32100, '2019-03-15 01:42:44', 51700, 1374, 1374, 0, 2, 1, 2, 0, 0),
  (166385, 178500, 172200, 23400, '2019-03-15 01:42:44', 17100, 1430, 1430, 0, 2, 1, 2, 0, 0),
  (166580, 437200, 434500, 34200, '2019-03-15 01:42:44', 31500, 1652, 1652, 0, 2, 1, 2, 0, 0),
  (166627, 192500, 210400, 12600, '2019-03-15 01:42:44', 30500, 1424, 1424, 0, 2, 1, 2, 0, 0),
  (166422, 737600, 697800, 76400, '2019-03-15 01:42:44', 36600, 1396, 1396, 0, 2, 1, 2, 0, 0),
  (166527, 88800, 99800, 12300, '2019-03-15 01:42:44', 23300, 1161, 1161, 0, 2, 1, 2, 0, 0),
  (166376, 109900, 105500, 18400, '2019-03-15 01:42:44', 14000, 1438, 1438, 0, 2, 1, 2, 0, 0),
  (166514, 174500, 163900, 24600, '2019-03-15 01:42:44', 14000, 1125, 1125, 0, 2, 1, 2, 0, 0),
  (166424, 309400, 295400, 20100, '2019-03-15 01:42:44', 6100, 1438, 1438, 0, 2, 1, 2, 0, 0),
  (166401, 463500, 458800, 6500, '2019-03-15 01:42:44', 1800, 1495, 1495, 0, 2, 1, 2, 0, 0),
  (166488, 143300, 144600, 11300, '2019-03-15 01:42:44', 12600, 1288, 1288, 0, 2, 1, 2, 0, 0),
  (166486, 241500, 253600, 34000, '2019-03-15 01:42:44', 46100, 1567, 1568, 1, 2, 1, 2, 0, 0),
  (166538, 156700, 141800, 38700, '2019-03-15 01:42:44', 23800, 1380, 1380, 0, 2, 1, 2, 0, 0),
  (166354, 1014700, 1041200, 44100, '2019-03-15 01:42:44', 70600, 1691, 1692, 1, 2, 1, 2, 0, 0),
  (166459, 142100, 152900, 19200, '2019-03-15 01:42:44', 30000, 1418, 1418, 0, 2, 1, 2, 0, 0),
  (166612, 383500, 401800, 22900, '2019-03-15 01:42:44', 41200, 1364, 1365, 1, 2, 1, 2, 0, 0),
  (166390, 320900, 294200, 50400, '2019-03-15 01:42:44', 23700, 1417, 1417, 0, 2, 1, 2, 0, 0),
  (166436, 46100, 123800, 76400, '2019-03-15 01:42:44', 154100, 1583, 1583, 0, 2, 1, 2, 0, 0),
  (166596, 212700, 212900, 6400, '2019-03-15 01:42:44', 6600, 1225, 1225, 0, 2, 1, 2, 0, 0),
  (166402, 346200, 345400, 7000, '2019-03-15 01:42:44', 6200, 1095, 1095, 0, 2, 1, 2, 0, 0),
  (166453, 158500, 152900, 6600, '2019-03-15 01:42:44', 1000, 1545, 1545, 0, 2, 1, 2, 0, 0),
  (166475, 173400, 177800, 8200, '2019-03-15 01:42:44', 12600, 1049, 1049, 0, 2, 1, 2, 0, 0),
  (166579, 187100, 208100, 12600, '2019-03-15 01:42:44', 33600, 1414, 1415, 1, 2, 1, 2, 0, 0),
  (166471, 220400, 241800, 41100, '2019-03-15 01:42:44', 62500, 1406, 1407, 1, 2, 1, 2, 0, 0),
  (166482, 149700, 147600, 7400, '2019-03-15 01:42:44', 5300, 1590, 1590, 0, 2, 1, 2, 0, 0),
  (166355, 842800, 846700, 22500, '2019-03-15 01:42:44', 26400, 1509, 1509, 0, 2, 1, 2, 0, 0),
  (166566, 121400, 126500, 7000, '2019-03-15 01:42:44', 12100, 1269, 1270, 1, 2, 1, 2, 0, 0),
  (166395, 233900, 237300, 33900, '2019-03-15 01:42:44', 37300, 1296, 1296, 0, 2, 1, 2, 0, 0),
  (166426, 137500, 120200, 62100, '2019-03-15 01:42:44', 44800, 1357, 1357, 0, 2, 1, 2, 0, 0),
  (166447, 293500, 291200, 6600, '2019-03-15 01:42:44', 4300, 1403, 1403, 0, 2, 1, 2, 0, 0),
  (166568, 113900, 110300, 28800, '2019-03-15 01:42:44', 25200, 1257, 1257, 0, 2, 1, 2, 0, 0),
  (166440, 97700, 89800, 47700, '2019-03-15 01:42:44', 39800, 1195, 1195, 0, 2, 1, 2, 0, 0),
  (166619, 505900, 545300, 39500, '2019-03-15 01:42:44', 78900, 1133, 1133, 0, 2, 1, 2, 0, 0),
  (166386, 197600, 195800, 7000, '2019-03-15 01:42:44', 5200, 1166, 1167, 1, 2, 1, 2, 0, 0)

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

