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
delete from user_exp where uid between 166668 and 167250;
update user_stat
set multi_rate = 1;

select count(1) from gift_pkg_log;
SELECT * FROM growth_task_log WHERE uid=166847 ORDER BY task_id;

SELECT * FROM portal_pay_order WHERE oid=144;
SELECT * FROM portal_pay_order WHERE orderId=144;

select * from user where token='f162fe500f22f01415039cc50f218e10';

call xianwan_user(167348,"MARKET_xw");
call xianwan_user2(167348,"MARKET_xw");

select variation,count(variation) from user_props_log inner join user on user.uid = user_props_log.uid where flavors='MARKET_sp'
 group by variation;