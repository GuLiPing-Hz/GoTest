update coin_log
set add_time='2019-05-21 07:00:00'
where change_type = 74 and add_time = '2019-05-20 07:00:00';
call proc_update_yuhuo_by_day('2019-05-22');

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

select count(1)
    from yt_coin_log
    where (uuid = 'ABBA1969-3E3A-84D9-39CA-A685A9476255' or uid = 165272)
      and tm >= '2019-05-27'
      and type = 0;
call proc_yt_checkin(165272,,165272,'2019-05-27 21:38:39')