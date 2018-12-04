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
insert into alms_log (uid, uuid, login_ip, coin, tm, platform, exp,charge)
values (177863, '43BA4F0B-776F-F184-6AFE-FB5C6095EA90', '183.156.125.192'
  , 8000, '2018-11-30 12:53:20', '0', 107850, 0);

UPDATE mission SET state=1 WHERE uid=177863 AND mid=2 AND state=0 AND value>=60;
insert into notice(title,content,sender,receiver,addtime,isValid,showOrder,ntype,mail_type,mail_giftid) values ('新手七天奖励','新手七天大礼，奖励已送达，请签收。',0,177863,'2018-12-04 11:26:20',1,2,1,5,'7');
select reward_id,reward_cnt,operator from mail_reward where mail_giftid = 7 and isvalid = 1;

SELECT totalsum FROM yule_gamelog ORDER BY round DESC limit 1;

