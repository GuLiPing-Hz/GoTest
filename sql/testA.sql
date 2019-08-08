update coin_log
set add_time='2019-06-16 07:00:00'
where change_type = 74
  and uid in (167369, 167392, 165272)
  and add_time > '2019-06-17';
call proc_update_yuhuo_by_day('2019-06-17');
select uid, -sum(fee)
from coin_log
where uid in (167392);
select a.uid, b.ytid, floor(-sum(fee) / 1000) as bill
from coin_log a
         inner join yt_user b on a.uid = b.uid
where b.apply = 0
  and a.change_type = 74
  and a.add_time >= '2019-06-16'
  and a.add_time < '2019-06-17'
group by a.uid, b.ytid;

update yt
set act=0
where true;

select a.uid, floor(-sum(fee) / 1000) as bill, -1
from coin_log a
         inner join yt_user b on a.uid = b.uid and a.change_type = 74
where b.apply = 0
  and a.add_time >= '2019-05-20'
  and a.add_time < '2019-05-21'
group by a.uid
having bill > 1000;

call proc_reset_by_day('2019-05-28 20:39:05', 1559047145700, 4);

select uid, floor(-sum(fee) / 1000) as bill
from coin_log a
where uid = 165272
  and a.change_type = 74
  and a.add_time >= '2019-05-28'
  and a.add_time < '2019-05-29'
group by a.uid
having bill >= 1000;


update yt
set ver=ver + 1
where true;

select a.uid, ytid, ifnull(yuhuoutc, 0) as yuhuoutc, utc
from yt_user a
         left join yt_yuhuo c on a.uid = c.uid
where a.uid = 165272
  and a.apply = 0;


select count(1) as cnt
from hbq_dui_log
where uid = 199287
  and tm >= '2019-06-02'
  and hbqType = 2;

insert into online_count(p_cnt, tp, tms, addtime) value (0, 2, 1, '2019-06-04 21:32:25');

call proc_reset_by_day('2019-06-18 00:00:00',1560787200000,4);
call online_log_summary('2019-06-18 00:00:00'); #数据统计。

select * from view_yt_apply where ytid=165272 limit 50;

select sum(money) as s from pay_log where uid=188801 and result=0 and channel in(1,2,3) and addtime>='2019-07-15'

explain select code from invite_log where uid=167338;


update growth_task_log set task_type=1 where task_id in(1,8,11,13,16);
update growth_task_log set task_type=2 where task_id in(5,9,15,17,19);
update growth_task_log set task_type=4 where task_id in(6);
update growth_task_log set task_type=5 where task_id in(10);
update growth_task_log set task_type=6 where task_id in(2,3,12,14,18,20,22,24,25,26);
update growth_task_log set task_type=7 where task_id in(21);
update growth_task_log set task_type=8 where task_id in(23);
update growth_task_log set task_type=9 where task_id in(7);
update growth_task_log set task_type=10 where task_id in(4);

explain select * from user where flavors like 'MARKET_xw%';

update user set sex=0 where true;
update user set sex=1 where flavors like 'MARKET_xw%';

alter table user modify sex int(10) null comment '指代是否是闲玩用户，0表示非闲玩渠道用户，1表示闲玩用户';

select * from user where ID_valid='A1000055435B2B' and sex=1 order by uid asc limit 1;