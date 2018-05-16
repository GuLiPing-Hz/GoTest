USE databasetest;

-- ----------------------------
-- Procedure structure for `proc_user_add` begin
-- ----------------------------
# MySQL默认以";"为分隔符，如果没有声明分割符，则编译器会把存储过程当成SQL语句进行处理，
# 因此编译过程会报错，所以要事先用“DELIMITER //”声明当前段分隔符，
# 让编译器把两个"//"之间的内容当做存储过程的代码，不会执行这些代码；“DELIMITER ;”的意为把分隔符还原。
DELIMITER //
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_user_add;
# DEFINER指定权限的存储过程
# CREATE DEFINER =`root`@`localhost` PROCEDURE `proc_adder`(IN a int, IN b int, OUT sum int)
CREATE PROCEDURE proc_user_add(IN account TEXT, IN pwd TEXT)
  BEGIN

    INSERT INTO databasetest.tabuser VALUES (NULL, account, pwd);

  END
//
#分隔符还原
DELIMITER ;