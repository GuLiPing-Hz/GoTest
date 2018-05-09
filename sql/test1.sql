#创建数据库
# DataBaseTest1 是数据库名字
CREATE DATABASE DataBaseTest1; #如果上一步已经做好了,那么在命令行中敲入:

#使用数据库 DataBaseTest1
USE DataBaseTest1;

#创建数据库 TabTest1
CREATE TABLE TabTest1
(

  #定义字段名 类型 默认值 主键 是否可空 自动增加
  row   BIGINT DEFAULT 0 PRIMARY KEY NOT NULL AUTO_INCREMENT,

  uid   TEXT                         NOT NULL, #用户ID

  #strict mode导致
  #BLOB, TEXT, GEOMETRY or JSON column 'name' can't have a default value
  name  TEXT, #用户昵称  ##default ""
  time  TIMESTAMP, #时间戳

  #整数
  #TINYINT SMALLINT MEDIUMINT INT   BIGINT
  #1字节    2字节    3字节     4字节  8字节

  #小数
  #FLOAT DOUBLE DECIMAL
  score INT, #得分

  #日期
  #DATE       TIME      YEAR   DATETIME             TIMESTAMP
  #YYYY-MM-DD HH:MM:SS  YYYY   YYYY-MM-DD HH:MM:SS  YYYYMMDD HHMMSS
  age   DATE, #年龄

  #字符串
  #CHAR         0~255字节    定长字符串 CHAR(10)
  #VARCHAR      0~65535字节  变长字符串
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
  COMMENT = '测试数据库1'; #增加表名注释

#删除表
DROP TABLE TabTest1;

#删除数据库
DROP DATABASE DataBaseTest1;