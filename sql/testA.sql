update coin_log
set add_time='2019-05-28 07:00:00'
where change_type = 74
  and add_time > '2019-05-29';
call proc_update_yuhuo_by_day('2019-05-29');

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


select count(1) as cnt from hbq_dui_log where uid=199287 and tm>='2019-06-02'
    and hbqType=2
