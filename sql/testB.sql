
select sum(money) promotionAwardMonth
from invite_log i
         left join pay_log l
                   on i.uid = l.uid
where result = 0
  and finish = 1
  and issandbox = 0
  and channel in (1, 2, 3, 6)
  and i.time <= addtime
  and i.code = '188895';
;

