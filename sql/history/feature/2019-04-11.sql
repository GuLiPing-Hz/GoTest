#停服，把redis数据推送到数据库，清理redis的usr_数据


#把老玩家的龙珠红包引导置为已完成
update user_stat
set guide_flag = guide_flag | 2;
