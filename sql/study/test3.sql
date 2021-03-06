#学习存储过程

# 存储过程 语法
# CREATE PROCEDURE  过程名([[IN|OUT|INOUT] 参数名 数据类型[,[IN|OUT|INOUT] 参数名 数据类型…]]) [特性 ...] 过程体
#
# DELIMITER //
#   CREATE PROCEDURE myproc(OUT s int)
#     BEGIN
#       SELECT COUNT(*) INTO s FROM students;
#     END
#     //
# DELIMITER ;

#创建存储过程或者执行存储过程必须指定某个数据库schemas
USE databasetest;

-- ----------------------------
-- Procedure structure for `proc_add` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_test_opt;
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# DEFINER指定权限的存储过程
# CREATE DEFINER =`root`@`localhost` PROCEDURE `proc_adder`(IN a int, IN b int, OUT sum int)
CREATE PROCEDURE proc_test_opt(IN a INT, INOUT b INT, IN opt INT, OUT sum INT)
BEGIN
    #Routine body goes here...
    DECLARE c INT;
    #声明变量c

    #由于指定sum为OUT参数，所以这里sum打印出来是NULL
    SELECT a,
           b,
           opt,
           sum;

    #IF语句
    IF a IS NULL
    THEN
        SET a = 0;
    END IF;

    #IF ELSEIF ELSE语句
    IF b IS NULL
    THEN
        SET b = 0;
    ELSEIF b = 0
    THEN
        SET b = 11;
    ELSE
        SET b = b + 1;
    END IF;

    SELECT a,
           b,
           sum;

    #CASE语句
    CASE opt
        WHEN 1
            THEN
                SET sum = a * b;
        WHEN 2
            THEN
                SET sum = a / b;
        WHEN 3
            THEN
                SET sum = a % b;
        WHEN 4
            THEN
                SET sum = a - b;
        ELSE
            SET sum = a + b;
        END CASE;

    SELECT a,
           b,
           sum;

    SET b = b * 10; #更新变量的值

    SELECT a,
           b,
           sum;

    SET c = 0;
    SELECT c;
    #WHILE 循环
    WHILE c < 3 DO
    SET c = c + 1;
    END WHILE;
    SELECT c;

    #REPEAT 循环
    REPEAT
        SET c = c + 1;
    UNTIL c > 5 END REPEAT;
    SELECT c;

    #LOOP 循环
    LOOP_1:
        LOOP
            IF c > 10
            THEN
                LEAVE LOOP_1;
            ELSE
                SET c = c + 1;
            END IF;
        END LOOP;
    SELECT c;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_add` END
-- ----------------------------

#调用
SET @a_in = 1, @b_in = 0;
SET @sum_out = 0;
#调用存储过程
CALL proc_test_opt(@a_in, @b_in, 5, @sum_out);
SELECT @b_in,
       @sum_out;

#移除存储过程
#DROP PROCEDURE [过程1[,过程2…]]
#DataGrip见schemas->database->routines 右击Drop


-- ----------------------------
-- Procedure structure for `proc_add` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_test_tmp_table;
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# DEFINER指定权限的存储过程
# CREATE DEFINER =`root`@`localhost` PROCEDURE `proc_adder`(IN a int, IN b int, OUT sum int)
CREATE PROCEDURE proc_test_tmp_table()
BEGIN
    #Routine body goes here...
    DECLARE sumTotal INT;
    #声明变量sumTotal

    #创建临时表，进行查询操作
    DROP TEMPORARY TABLE IF EXISTS tmp_t_;
    CREATE TEMPORARY TABLE tmp_t_
    select uuid,
           name,
           sum(score) as sum
    from databasetest.tabtest1
    group by uuid, name;

    #ifnull判断，如果为空，填入后面一个值
    select ifnull(sum(score), 0) as s into sumTotal -- 只接受一行数据
    from databasetest.tabtest1;
    #查询到的结果
    select *
    from tmp_t_;

    select sumTotal;

    -- 删除临时表
    DROP TEMPORARY TABLE IF EXISTS tmp_t_;
END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_add` END
-- ----------------------------
call proc_test_tmp_table();

#存储过程模板
-- ----------------------------
-- Procedure structure for `proc_insert_pay` BEGIN
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_insert_pay;
CREATE PROCEDURE proc_insert_pay(in vUid bigint, vTm datetime)
exec:
BEGIN
    leave exec;
END;
-- ----------------------------
-- Procedure structure for `proc_insert_pay` END
-- ----------------------------

#游标调用模板
drop procedure if exists StatisticStore;
CREATE PROCEDURE StatisticStore()
BEGIN
    -- 创建接收游标数据的变量
    declare c int;
    declare n varchar(20);
    -- 创建总数变量
    declare total int default 0;
    -- 创建结束标志变量
    declare done int default false;
    -- 创建游标
    declare cur cursor for select name, count from store where name = 'iphone';
    -- 指定游标循环结束时的返回值
    declare continue HANDLER for not found set done = true;
    -- 设置初始值
    set total = 0;
    -- 打开游标
    open cur;
    -- 开始循环游标里的数据
    read_loop:
        loop
            -- 根据游标当前指向的一条数据
            fetch cur into n,c;
            -- 判断游标的循环是否结束
            if done then
                leave read_loop; -- 跳出游标循环
            end if;
            -- 获取一条数据时，将count值进行累加操作，这里可以做任意你想做的操作，
            set total = total + c;
            -- 结束游标循环
        end loop;
    -- 关闭游标
    close cur;

    -- 输出结果
    select total;
END;

# call proc_get_my_promotion(165272);
# MySQL存储过程的基本函数
# 字符串类
# CHARSET(str) //返回字串字符集
# CONCAT (string2 [,... ]) //连接字串
# INSTR (string ,substring ) //返回substring首次在string中出现的位置,不存在返回0
# LCASE (string2 ) //转换成小写
# LEFT (string2 ,length ) //从string2中的左边起取length个字符
# LENGTH (string ) //string长度
# LOAD_FILE (file_name ) //从文件读取内容
# LOCATE (substring , string [,start_position ] ) 同INSTR,但可指定开始位置
# LPAD (string2 ,length ,pad ) //重复用pad加在string开头,直到字串长度为length
# LTRIM (string2 ) //去除前端空格
# REPEAT (string2 ,count ) //重复count次
# REPLACE (str ,search_str ,replace_str ) //在str中用replace_str替换search_str
# RPAD (string2 ,length ,pad) //在str后用pad补充,直到长度为length
# RTRIM (string2 ) //去除后端空格
# STRCMP (string1 ,string2 ) //逐字符比较两字串大小,
# SUBSTRING (str , position [,length ]) //从str的position开始,取length个字符,
# 注：mysql中处理字符串时，默认第一个字符下标为1，即参数position必须大于等于1
# TRIM([[BOTH|LEADING|TRAILING] [padding] FROM]string2) //去除指定位置的指定字符
# UCASE (string2 ) //转换成大写
# RIGHT(string2,length) //取string2最后length个字符
# SPACE(count) //生成count个空格

# 数学类
# ABS (number2 ) //绝对值
# BIN (decimal_number ) //十进制转二进制
# CEILING (number2 ) //向上取整
# CONV(number2,from_base,to_base) //进制转换
# FLOOR (number2 ) //向下取整
# FORMAT (number,decimal_places ) //保留小数位数
# HEX (DecimalNumber ) //转十六进制
# 注：HEX()中可传入字符串，则返回其ASC-11码，如HEX('DEF')返回4142143
# 也可以传入十进制整数，返回其十六进制编码，如HEX(25)返回19
# LEAST (number , number2 [,..]) //求最小值
# MOD (numerator ,denominator ) //求余
# POWER (number ,power ) //求指数
# RAND([seed]) //随机数
# ROUND (number [,decimals ]) //四舍五入,decimals为小数位数] 注：返回类型并非均为整数，如：
# SIGN (number2 ) // 正数返回1，负数返回-1

# 日期时间类 - 有些函数可以参考test2.sql中的例子
# ADDTIME (date2 ,time_interval ) //将time_interval加到date2
# CONVERT_TZ (datetime2 ,fromTZ ,toTZ ) //转换时区
# CURRENT_DATE ( ) //当前日期
# CURRENT_TIME ( ) //当前时间
# CURRENT_TIMESTAMP ( ) //当前时间戳
# DATE (datetime ) //返回datetime的日期部分
# DATE_ADD (date2 , INTERVAL d_value d_type ) //在date2中加上日期或时间
# DATE_FORMAT (datetime ,FormatCodes ) //使用formatcodes格式显示datetime
# DATE_SUB (date2 , INTERVAL d_value d_type ) //在date2上减去一个时间
# DATEDIFF (date1 ,date2 ) //两个日期差
# DAY (date ) //返回日期的天
# DAYNAME (date ) //英文星期
# DAYOFWEEK (date ) //星期(1-7) ,1为星期天
# DAYOFYEAR (date ) //一年中的第几天
# EXTRACT (interval_name FROM date ) //从date中提取日期的指定部分
# MAKEDATE (year ,day ) //给出年及年中的第几天,生成日期串
# MAKETIME (hour ,minute ,second ) //生成时间串
# MONTHNAME (date ) //英文月份名
# NOW ( ) //当前时间
# SEC_TO_TIME (seconds ) //秒数转成时间
# STR_TO_DATE (string ,format ) //字串转成时间,以format格式显示
# TIMEDIFF (datetime1 ,datetime2 ) //两个时间差
# TIME_TO_SEC (time ) //时间转秒数]
# WEEK (date_time [,start_of_week ]) //第几周
# YEAR (datetime ) //年份
# DAYOFMONTH(datetime) //月的第几天
# HOUR(datetime) //小时
# LAST_DAY(date) //date的月的最后日期
# MICROSECOND(datetime) //微秒
# MONTH(datetime) //月
# MINUTE(datetime) //分返回符号,正负或0
# SQRT(number2) //开平方
