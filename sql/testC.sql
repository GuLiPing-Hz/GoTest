select sum(money) as s from pay_log where uid=165272
		and result=0 and channel in(1,2,3) and addtime>='2019-09-15'


