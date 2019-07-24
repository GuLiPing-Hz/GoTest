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
