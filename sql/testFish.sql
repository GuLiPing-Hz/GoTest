SELECT *
FROM Buyu.user_stat
WHERE uid = 170652;
SELECT *
FROM Buyu.user
WHERE uid = 177851;
SELECT *
FROM Buyu.user_stat
WHERE uid = 177851;

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



SELECT * FROM ip_rule WHERE id>9 OR tm>ADDDATE(NOW(),INTERVAL -2 MINUTE);
select * from mail_reward where id = 21338;

