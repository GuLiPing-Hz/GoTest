#学习数据库数据格式日期，还有数据的操作（增删改查）


#日期 相关操作

#下面的 utc 当做是取别名
SELECT UTC_TIMESTAMP() utc; #获取UTC标准日期+时间             2018-05-10 07:42:58

SELECT NOW(); #获取本地日期+时间，已经经过时区转换              2018-05-10 15:42:50
#下面三个等价于NOW();,建议使用NOW()
#SELECT CURRENT_TIMESTAMP(),LOCALTIME(),LOCALTIMESTAMP()

SELECT SYSDATE(); #日期时间跟 NOW() 类似，不同之处在于：now() 在执行开始时值就得到了， sysdate() 在函数执行时动态得到值。

#查看下面的例子，明确不同之处，两个 NOW()的时间一致，SYSDATE() 相差3秒，就是我们SLEEP的3秒
SELECT
  NOW(),
  SYSDATE(),
  SLEEP(3),
  NOW(),
  SYSDATE();

SELECT
  CURDATE(),
  CURRENT_DATE(); #获取本地日期                      2018-05-10

SELECT
  CURTIME(),
  CURRENT_TIME(); #获取本地时间                      15:42:08

SELECT
  UTC_DATE(),
  UTC_TIME();

SET @dt = '2018-05-10 07:15:30.123456';
SELECT
  DATE(@dt),
  # 日期 2018-05-10
  TIME(@dt),
  # 时间 07:15:30
  YEAR(@dt),
  # 年 2018
  QUARTER(@dt),
  # 季度 2
  MONTH(@dt),
  # 月 5
  WEEK(@dt),
  # 周 18
  DAY(@dt),
  # 日 9
  HOUR(@dt),
  # 小时 7
  MINUTE(@dt),
  # 分 15
  SECOND(@dt),
  # 秒 30
  MICROSECOND(@dt); # 123456

# EXTRACT 函数使用截取日期
SELECT EXTRACT(YEAR FROM @dt); #选取年份 2018
SELECT EXTRACT(QUARTER FROM @dt); #选取季度 2
SELECT EXTRACT(MONTH FROM @dt); #选取月份 5
SELECT EXTRACT(WEEK FROM @dt); #选取周 18
SELECT EXTRACT(DAY FROM @dt); #选取日 10
SELECT EXTRACT(HOUR FROM @dt); #选取小时 7
SELECT EXTRACT(MINUTE FROM @dt); #选取分钟 15
SELECT EXTRACT(SECOND FROM @dt); #选取秒 30
SELECT EXTRACT(MICROSECOND FROM @dt); #选取毫秒 123456

#选取 某个值到另一个值的闭区间  只提供下面几个关键字，，没有更多了。。。
SELECT EXTRACT(YEAR_MONTH FROM @dt); #选取年到月 201805
SELECT EXTRACT(DAY_HOUR FROM @dt); #选取日到小时 1007
SELECT EXTRACT(DAY_MINUTE FROM @dt); # 100715
SELECT EXTRACT(DAY_SECOND FROM @dt); # 10071530
SELECT EXTRACT(DAY_MICROSECOND FROM @dt); # 10071530123456  从日期读取到毫秒的数据
SELECT EXTRACT(HOUR_MINUTE FROM @dt); # 715
SELECT EXTRACT(HOUR_SECOND FROM @dt); # 71530
SELECT EXTRACT(HOUR_MICROSECOND FROM @dt); # 71530123456
SELECT EXTRACT(MINUTE_SECOND FROM @dt); # 1530
SELECT EXTRACT(MINUTE_MICROSECOND FROM @dt); # 1530123456
SELECT EXTRACT(SECOND_MICROSECOND FROM @dt); # 30123456

#日期 DAYOF 函数  DAYOFWEEK（1 = SunDAY, 2 = MonDAY, …, 7 = SaturDAY）
SELECT
  DAYOFWEEK(@dt),
  DAYOFMONTH(@dt),
  DAYOFYEAR(@dt);

SHOW VARIABLES LIKE 'default_WEEK_format';
set @dt = '2018-05-10';
SELECT WEEK(@dt); # 18
SELECT WEEK(@dt, 3); # 19
SELECT WEEKOFYEAR(@dt); # 19  等价于 WEEK(日期,3)
SELECT DAYOFWEEK(@dt); # 5
SELECT WEEKDAY(@dt); # 3
SELECT YEARWEEK(@dt); # 201818 返回 year(2018) + week 位置(18)。

#名称
select DAYNAME(@dt); # Friday
select MONTHNAME(@dt); # August

#返回传入日期月份的最后一天的日期
select LAST_DAY('2018-02-01'); # 2008-02-29
select LAST_DAY('2018-08-08'); # 2008-08-31

#计算某个日期的月份有多少天
SELECT DAY(LAST_DAY('2019-02-01'));

#日期增加/减少
SELECT DATE_ADD(NOW(), INTERVAL '1 1:1:1' DAY_SECOND); #增加1天1小时1分1秒 替代 ADDDATE() ADDTIME()
SELECT DATE_ADD(NOW(), INTERVAL '0 0:1:0' DAY_SECOND); #增加0天0小时1分0秒
SELECT DATE_SUB(NOW(), INTERVAL '1 1:1:1' DAY_SECOND); #减少1天1小时1分1秒 替代 SUBDATE() SUBTIME()
SELECT DATE_SUB(NOW(), INTERVAL '1-1' YEAR_MONTH); #减少1年1月


#PERIOD_ADD 日期格式必须是 YYYYMM YYMM
SELECT PERIOD_ADD(EXTRACT(YEAR_MONTH FROM NOW()), 2); # 对日期增加N个月 N可以为负数
SELECT PERIOD_ADD(EXTRACT(YEAR_MONTH FROM NOW()), -2);

#计算两个日期间隔
#PERIOD_DIFF(P1,P2)：日期 P1-P2，返回 N 个月。 日期格式必须是 YYYYMM YYMM
SELECT PERIOD_DIFF('201805', '201804'); #计算相差几个月
SELECT DATEDIFF(NOW(), UTC_TIMESTAMP()); #计算相差几天
#TIMEDIFF 只支持传入时间格式的，不能含日期，如果含日期，必须日期一致
SELECT TIMEDIFF('22:11:11', '20:10:10'); #('2018-05-10 22:10:10', '2018-04-10 20:10:10'); 报错
SELECT TIMEDIFF(NOW(), UTC_TIMESTAMP()); #计算相差多少时间，只是计算时间差值


#数据 增删改查

#插入数据
#自增列可以插入 NULL 或者 0
INSERT INTO databaSETest.tabtest1 VALUES (0, '100001', 'Aaa', UTC_TIMESTAMP(), 99, '2000-01-01', '13711111111');
INSERT INTO databaSETest.tabtest1 VALUES (NULL, '100001', 'Aaa', UTC_TIMESTAMP(), 99, '2000-01-01', '13711111111');
INSERT INTO databaSETest.tabtest1 (uuid, name) values ('100002', 'Bbb');

#删
DELETE FROM databasetest.tabtest1
WHERE rowId = 2 AND name = 'Aaa';

#改 改多列的时候，中间需要加上逗号
UPDATE databasetest.tabtest1
SET name = 'Ccc', time = UTC_TIMESTAMP()
WHERE rowId = 6;

#查
#查所有数据
SELECT *
FROM databasetest.tabtest1;
#查多少条记录
SELECT COUNT(*)
FROM databasetest.tabtest1;
#查名字是Aaa的记录
SELECT COUNT(*)
FROM databasetest.tabtest1
WHERE name = 'Aaa' OR score >= 60;
SELECT
  name  as n,
  score as s
FROM databasetest.tabtest1;
#LIMIT 指定获取前面多少行数据
SELECT *
FROM databasetest.tabtest1
WHERE score >= 60
LIMIT 3;
#LIMIT 指定获取第2行到第4行的数据，索引从0开始
SELECT *
FROM databasetest.tabtest1
WHERE score >= 60
LIMIT 2, 4;
#SELECT TOP 2 * FROM databasetest.tabtest1;  不支持TOP语句，用LIMIT代替

#ORDER BY排序
SELECT *
FROM databasetest.tabtest1
ORDER BY name DESC; #默认升序，可以用DESC指定降序,升序是ASC

#GROUP BY 会把字段值一样的值组合在一起，在GROUP BY中即可使用聚合函数，求个数COUNT,求和SUM
SELECT
  name,
  #数量
  COUNT(*),
  COUNT(name),
  COUNT(DISTINCT name),
  #求和
  SUM(score),
  #求平均数
  AVG(score),
  #最小值
  MIN(score),
  #最大值
  MAX(score),
  #标准差
  STD(score),
  #把不同的数据组合成一列中，用逗号分隔
  GROUP_CONCAT(score)
FROM databasetest.tabtest1
GROUP BY name DESC;

#DISTINCT 去重 如果指定多个字段，则表示需要每个字段都一样才会去重
SELECT DISTINCT
  name,
  score
FROM databasetest.tabtest1;
SELECT
  name,
  score
FROM databasetest.tabtest1;

# 语法：
# SELECT
#   select_expr [,
#   select_expr...]
#
# FROM table_name
#
# [WHERE where_condition]   #where子句在分组前对记录进行过滤
#
# [GROUP BY {col_name | expr} [ASC | DESC], ... [WITH ROLLUP]]
#
# [HAVING where_condition]  #having子句在分组后对记录进行过滤

