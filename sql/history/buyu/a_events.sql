DELIMITER //
CREATE definer = root@`192.168.100.3` event agent_promotion_reward
  on schedule
    every '1' DAY
      starts '2018-05-02 13:08:10'
  on completion preserve
  enable
do
  BEGIN
CALL agent_promotion_reward();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event agent_reward
  on schedule
    every '1' DAY
      starts '2018-06-21 18:18:27'
  enable
do
  BEGIN
CALL agent_recharge_reward();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event day_coin
  on schedule
    every '24' HOUR
      starts '2018-07-17 00:05:00'
  enable
do
  BEGIN
CALL day_coin();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event day_longzhu_log
  on schedule
    every '24' HOUR
      starts '2018-07-17 00:05:00'
  enable
do
  BEGIN
CALL day_longzhu_log();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event day_pool_log
  on schedule
    every '1' DAY
      starts '2018-11-21 00:05:00'
  on completion preserve
  enable
do
  BEGIN
CALL day_pool_log();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event day_summary_log_new
  on schedule
    every '24' HOUR
      starts '2018-07-17 00:10:00'
  enable
do
  BEGIN
CALL day_summary_log_new();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event day_yule_summary
  on schedule
    every '24' HOUR
      starts '2018-11-01 00:20:00'
  enable
do
  BEGIN
CALL day_yule_summary();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event stats_coin
  on schedule
    every '1' HOUR
      starts '2018-11-19 13:00:00'
  on completion preserve
  enable
do
  BEGIN
CALL stats_coin();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event stats_remain
  on schedule
    every '1' DAY
      starts '2018-10-12 00:05:00'
  on completion preserve
  enable
do
  BEGIN
call stat_remain_player();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event stats_summary
  on schedule
    every '1' DAY
      starts '2018-11-14 00:05:00'
  on completion preserve
  enable
do
  BEGIN
CALL stats_summary();
END
//
DELIMITER ;

DELIMITER //
CREATE definer = root@`192.168.100.3` event stats_user_coin
  on schedule
    every '1' HOUR
      starts '2018-07-26 11:20:00'
  on completion preserve
  enable
do
  BEGIN
CALL stats_user_coin();
END
//
DELIMITER ;

