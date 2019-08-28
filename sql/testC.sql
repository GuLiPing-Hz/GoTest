#计算下周一算法。
set @now = '2019-08-18';
select dayofweek(@now);
set @day = (2 + 7 - (select dayofweek(@now)));
select @day;
if @day > 7 then
    @day = @day -7
end if;
select date_add(@now, INTERVAL @day DAY);
