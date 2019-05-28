CREATE TABLE online_log2
(
  uid   bigint PRIMARY KEY NOT NULL,
  vip   tinyint
  COMMENT '今日vip等级',
  value int COMMENT '大厅在线时长 单位(分)',
  lobby int comment '大厅在线时长 单位(秒)',
  game  int comment '游戏在线时长 单位(秒)'
);
ALTER TABLE online_log2
  COMMENT = '今日用户在线时长统计';

# select
#   user.avatar,
#   user.nick_name name,
#   vip,
#   value,
#   online_log2.uid
# from online_log2
#   inner join user on online_log2.uid = user.uid
# ORDER BY value DESC
# LIMIT 50 OFFSET 0;


-- ----------------------------
-- Procedure structure for `proc_update_onlinelog` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_update_onlinelog;
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# DEFINER指定权限的存储过程
CREATE PROCEDURE proc_update_onlinelog(in uidIn   bigint, in vipIn tinyint, in seconds int,
                                       in lobbyIn int, in gameIn int)
  BEGIN
    declare cnt int;
    select count(1)
    into cnt
    from online_log2
    where uid = uidIn;

    if cnt > 0
    then
      update online_log2
      set value = seconds, lobby = lobbyIn, game = gameIn, vip = vipIn
      where uid = uidIn;
    else
      insert into online_log2 (uid, vip, value, lobby, game) values (uidIn, vipIn, seconds, lobbyIn, gameIn);
    end if;
  END
//
#分隔符还原
DELIMITER ;
-- ----------------------------
-- Procedure structure for `proc_update_onlinelog` END
-- ----------------------------
# call proc_update_onlinelog(167367, 1, 10);

ALTER TABLE user_stat MODIFY online_time bigint(20) DEFAULT '1' COMMENT '用户在线总时长 废弃';