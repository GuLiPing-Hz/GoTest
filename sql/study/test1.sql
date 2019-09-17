#初步接触数据库
#创建数据库和表，
#对表字段进行操作

#设置mysql超时断开时间
show global variables like '%timeout%';
#这个只能设置当前会话，退出登录又没了
# set interactive_timeout = 604800; #60*60*8=28800 8小时不断
set global interactive_timeout = 604800;
set global wait_timeout = 604800;
#60*60*24=604800 7*24小时不断

#批量插入数据库限制
show global variables like '%max_al%';
set global max_allowed_packet = 1073741824;

#查看mysql连接数上限
show global variables like '%max_connections%';
set global max_connections = 1000;

#查看sql_mode
show global variables like '%sql_mode%';
set global sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION';

#添加远程访问权限
#https://www.jianshu.com/p/92fdef0dab32

#非常重要的优化命令
#查看mysql运行状态
show status;
#在数据库服务器高峰期，可用此命令查看当前比较费时的sql语句，+explain优化之
show full processlist;
# 如果某个查询一直在执行可以直接通过命令 kill 指定ID : kill 192610;


#重启mysql服务器
#/etc/init.d/mysqld restart

#创建数据库
# DataBaseTest1 是数据库名字
CREATE DATABASE DataBaseTest1;
#如果上一步已经做好了,那么在命令行中敲入:

#使用数据库 DataBaseTest1
USE DataBaseTest1;

#创建数据库表 TabTest1
CREATE TABLE TabTest1
(
    #定义字段名 类型 默认值 主键 是否可空 自动增加
    rowId BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,

    #用户ID
    uid   TEXT               NOT NULL,

    #strict mode导致
    #BLOB, TEXT, GEOMETRY or JSON column 'name' can't have a default value
    name  TEXT,          #用户昵称  ##default ""
    time  TIMESTAMP,     #时间戳

    #整数
    #TINYINT SMALLINT MEDIUMINT INT   BIGINT
    #1字节    2字节    3字节     4字节  8字节

    #小数
    #FLOAT DOUBLE DECIMAL
    score INT DEFAULT 0, #得分

    #日期
    #DATE       TIME      YEAR   DATETIME             TIMESTAMP
    #YYYY-MM-DD HH:MM:SS  YYYY   YYYY-MM-DD HH:MM:SS  YYYYMMDD HHMMSS
    age   DATE,          #年龄

    #字符串
    #CHAR         0~255字节    定长字符串 CHAR(10)
    #VARCHAR      0~65535字节  变长字符串 VARCHAR(200)
    #TINYTEXT	    0-255字节
    #TEXT	        0-65535字节
    #MEDIUMTEXT	  0-16777215字节
    #LONGTEXT	    0-4294967295字节

    #二进制字符串
    #TINYBLOB     0-255字节
    #BLOB	        0-65535字节
    #MEDIUMBLOB	  0-16777215字节
    #LONGBLOB	    0-4294967295字节
    phone CHAR(11)#手机号
);
ALTER TABLE TabTest1
    COMMENT = '测试数据库1';
#增加表注释

#删除表
DROP TABLE TabTest1;

#删除数据库
DROP DATABASE DataBaseTest1;

CREATE DATABASE DataBaseTest;
USE DataBaseTest;

#字段 增删改
#增加字段：
ALTER TABLE tabtest1
    ADD area TINYINT Default 86;
#移除字段
ALTER TABLE DataBaseTest.tabtest1
    DROP COLUMN area;

#修改字段名：
ALTER TABLE tabtest1
    RENAME COLUMN uid TO uuid;
# uuid to uid
#修改字段类型：
ALTER TABLE tabtest1
    MODIFY COLUMN name VARCHAR(200);

#创建临时表
USE databasetest;
CREATE TEMPORARY TABLE tabUserTemp
(
    rowId   INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    account TEXT            NOT NULL,
    pwd     TEXT            NOT NULL
);
#临时表的插入查询删除跟普通表操作一样

#生成用户表
USE databasetest;
CREATE TABLE tabUser
(
    rowId   INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    account TEXT            NOT NULL,
    pwd     TEXT            NOT NULL
);
# CREATE UNIQUE INDEX tabUser_rowId_uindex ON tabUser (rowId);
CREATE UNIQUE INDEX tabUser_account_uindex
    ON tabUser (account);
ALTER TABLE tabUser
    COMMENT = '用户表';


