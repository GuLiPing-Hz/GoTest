update coin_log
set add_time='2019-05-20-07:00'
where change_type = 74;
call proc_update_yuhuo_by_day('2019-05-21');

update yt set act=0 where true;


select a.uid, floor(-sum(fee) / 1000) as bill, -1
from coin_log a
         inner join yt_user b on a.uid = b.uid and a.change_type = 74
where b.apply = 0
  and a.add_time >= '2019-05-20'
  and a.add_time < '2019-05-21'
group by a.uid
having bill > 1000;

call proc_get_yuhuo(167374,165272,165272,'2019-05-21 14:05:42',1558418742)

select name from yt where ytid=165272;