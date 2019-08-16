# 优化每日任务，防止出现插入重复的数据

create table online_reward_log2
(
    uid       bigint  not null comment '用户ID',
    reward_id tinyint not null comment 'online_reward_cfg中的id',
    state     tinyint  default 0 comment '当前任务状态',
    tm        datetime default now() comment '数据更新时间',
    constraint online_reward_log2_pk
        primary key (uid, reward_id)
)
    comment '每日任务优化存储数据';


drop procedure if exists proc_reset_by_day;
create procedure proc_reset_by_day(IN tm datetime, IN tm_utc bigint, IN vip_limit int)
BEGIN
    #Routine body goes here...
    #     DECLARE sumTotal INT; #声明变量sumTotal
    #每日0点清空在线日志
    call online_log_summary(tm); #数据统计。
    DELETE FROM online_log2 where true;
    DELETE from online_reward_log2 where true;
    #清理表
    #更新每日捕获龙的记录
    INSERT INTO catch_dragon_log (uid, dragon_cnt, tm)
    SELECT uid,
           value,
           tm
    FROM mission
    WHERE value > 0
      AND mid = 4;
    #用户每日状态信息清空
    UPDATE user_stat
    SET lz_points     = 0,
        transfer_coin = 0
    where true;

    #更新VIP每日奖励领取状态
    update user_stat
    set isgetvipreward = 1
    WHERE vip >= vip_limit;
    #     update user_stat
    #     set isgetvipreward = 0
    #     WHERE vip < vip_limit;

    #更新任务状态
    UPDATE mission
    set value = 0,
        state = 0
    where true;

    #更新活跃度金币
    UPDATE active_coin
    SET state     = 0,
        update_tm = tm
    where true;

    #把炮台过期的炮置为默认炮台
    UPDATE user
        inner join fort_log on user.uid = fort_log.uid and weapon = fort_id
    SET weapon = 0
    WHERE expire_utc > 0
      AND expire_utc < tm_utc;

    #更新周月卡记录，把过期的置为无效状态。
    UPDATE card_log
    SET state = 0
    WHERE end_tm < tm_utc
      and state = 1;
END;



