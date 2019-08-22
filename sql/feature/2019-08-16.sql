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

INSERT INTO `Buyu`.`activities` (`id`, `title`, `endTime`, `sortNo`, `isEnabled`)
VALUES (13, '找回账号', '2019-10-01 00:00:00', 13, 1);
INSERT INTO `Buyu`.`activities` (`id`, `title`, `endTime`, `sortNo`, `isEnabled`)
VALUES (14, '游戏推广', '2019-10-01 00:00:00', 14, 1);

alter table user
    modify type tinyint default 0 null comment '用户类型：
0.普通用户
1.普通推广员(不可登录游戏，可登后台)
2.普通推广员(可登录游戏，不可登后台)
3.渔乐机器人，
4捕鱼机器人，
5后台推广（待审核,不可登游戏，不可登后台）';

alter table invite_log
    add status tinyint default 2 null comment '推广成员的奖励是否已经领取，默认已领取。';

# 更新邀请奖励
update invite_cfg
set inviterGift=50000
where id = 1;

-- ----------------------------
-- Procedure structure for `proc_add` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_get_my_promotion;
CREATE PROCEDURE proc_get_my_promotion(in vUid bigint)
BEGIN
    declare cnt int;
    declare cnt2 int;
    #查询总共推广了多少用户
    select count(1) into cnt from invite_log where code = vUid;
    #查询当前还有多少推广用户已经升到5级，但是还没有领取的推广奖励
    select count(1) into cnt2 from invite_log where code = vUid and status = 1;
    select cnt, cnt2;
END;
-- ----------------------------
-- Procedure structure for `proc_get_my_promotion` END
-- ----------------------------
# call proc_get_my_promotion(165272);

alter table yt
    add tgy bigint default 0 null comment '鱼塘推广员。';

-- ----------------------------
-- Procedure structure for `proc_create_yt` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_create_yt;
CREATE PROCEDURE proc_create_yt(in vUid bigint, in vName text, in vIntro text,
                                in vReward int, in vTgy bigint, in vTm datetime, in vPool bigint)
exec:
BEGIN
    # 创建一个鱼塘
    # @return status:
    #     0 成功
    #     10123, --您已加入一个鱼塘

    declare vCnt int;
    select count(1) into vCnt from yt_user where uid = vUid and apply = 0 and ytid > 0;
    if vCnt > 0 then
        select 10123 as status;
        leave exec;
    end if;

    #删除所有的申请消息
    delete from yt_user where uid = vUid;
    #加入到鱼塘中
    insert into yt_user(uid, ytid, tm, apply) value (vUid, vUid, vTm, 0);
    #创建一个鱼塘
    insert into yt(ytid, uid, name, intro, reward, tm, pool, tgy)
        value (vUid, vUid, vName, vIntro, vReward, vTm, vPool, vTgy);
    select 0 as status;
END;
-- ----------------------------
-- Procedure structure for `proc_create_yt` END
-- ----------------------------

drop view if exists view_yt_rank_act_last;
create view view_yt_rank_act_last as
select a.ytid,
       nickname,
       name,
       a.act,
       num
from yt_rank_last a
         inner join view_yt_rank_act b on a.ytid = b.ytid
where a.act > 0;

UPDATE `Buyu`.`yt_create_cfg` t
SET t.`reward` = 0
WHERE t.`id` = 1;
