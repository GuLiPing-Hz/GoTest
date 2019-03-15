#学习定时任务

#查看event是否开启
show variables like '%event_sche%';
#开启event
set global event_scheduler = 1;
#关闭event
set global event_scheduler = 0;

#查看当前所有事务
show events;

#创建一个定时任务

# ON SCHEDULE 设置时间规则
#决定event的执行时间和频率（注意时间一定要是将来的时间，过去的时间会出错），有两种形式 AT和EVERY。
#AT 是指定时间执行一次
# 时间戳需要大于当前时间
#EVERY 是循环执行。
# 在重复的计划任务中，时间（单位）的数量可以是任意非空（Not Null）的整数式，
# 时间单位是关键词：YEAR，MONTH，DAY，HOUR，MINUTE 或者SECOND

# STARTS 决定了什么时候开启这个定时任务

# ON COMPLETION参数表示"当这个事件不会再发生的时候"，
# 即当单次计划任务执行完毕后或当重复性的计划任务执行到了ENDS阶段。
# 而PRESERVE的作用是使事件在执行完毕后不会被Drop掉，建议使用该参数，以便于查看EVENT具体信息

-- ----------------------------
-- #每日数据统计定时任务 begin
-- ----------------------------
DROP EVENT IF EXISTS event_test;
DELIMITER //
CREATE EVENT event_test
  ON SCHEDULE
    EVERY 2 second
      STARTS TIMESTAMP '2018-11-28 17:06:00'
  ON COMPLETION PRESERVE
  COMMENT '每日测试任务'
DO
  BEGIN
    insert into databasetest.event_test_table values ();
    #call proc(); #调用某个存储过程
  END
//
DELIMITER ;
-- ----------------------------
-- #每日数据统计定时任务 end
-- ----------------------------


#关闭事件任务 :
ALTER EVENT event_test
ON COMPLETION PRESERVE
DISABLE;

#开启事件任务 :
ALTER EVENT event_test
ON COMPLETION PRESERVE
ENABLE;



