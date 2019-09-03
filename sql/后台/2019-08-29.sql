-- auto-generated definition
create table agent_award_log
(
    uid        bigint(16)                           not null comment '用户id',
    reward     bigint(16) default 0 comment '返利',
    createTime timestamp  default CURRENT_TIMESTAMP comment '创建时间',
    canGetTime timestamp  default CURRENT_TIMESTAMP not null comment '可领取时间',
    updateTime timestamp  default CURRENT_TIMESTAMP comment '领取时间',
    type       tinyint    default 3                 not null comment '
福利类型：
1.推广用户充值
2.鱼塘充值
3.出售金币',
    status     tinyint    default 1                 not null comment '领取状态：1.未领取2.已领取',
    primary key (uid, canGetTime, type)
)
    comment '福利记录';

create table agent_award_cfg
(
    type    tinyint null comment '1 推广用户充值，2鱼塘重置，3出售金币',
    percent int     null comment '%，百分数'
);

create unique index agent_award_cfg_type_uindex
    on agent_award_cfg (type);

alter table agent_award_cfg
    add constraint agent_award_cfg_pk
        primary key (type);

-- ----------------------------
-- Procedure structure for `proc_agent_insert_sell` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_agent_insert_sell;
CREATE PROCEDURE proc_agent_insert_sell(in vNow datetime, in vGold bigint, in vUid bigint)
BEGIN
    #代理发售金币需要调用的存储过程，实时汇结发售金币的返利
    declare nextMonday datetime;
    declare nnextMonday datetime;
    declare nnnextMonday datetime;
    declare nnnnextMonday datetime;
    declare rewardPercent int;
    declare todayOfWeek int;
    declare rewardPerWeek bigint;

    select percent into rewardPercent from agent_award_cfg where type = 3;

    set vNow = date(vNow);
    set todayOfWeek = (2 + 7 - (select dayofweek(vNow)));
    if todayOfWeek > 7 then
        set todayOfWeek = todayOfWeek - 7;
    end if;
    set nextMonday = date_add(vNow, INTERVAL todayOfWeek DAY);
    set nnextMonday = date_add(nextMonday, INTERVAL 7 DAY);
    set nnnextMonday = date_add(nextMonday, INTERVAL 14 DAY);
    set nnnnextMonday = date_add(nextMonday, INTERVAL 21 DAY);

    set rewardPerWeek = vGold * rewardPercent / 4 / 100;
    insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, vNow, nextMonday, vNow, 3, 1)
    on duplicate key update reward=reward + rewardPerWeek, updateTime=vNow;
    insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, vNow, nnextMonday, vNow, 3, 1)
    on duplicate key update reward=reward + rewardPerWeek, updateTime=vNow;
    insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, vNow, nnnextMonday, vNow, 3, 1)
    on duplicate key update reward=reward + rewardPerWeek, updateTime=vNow;
    insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, vNow, nnnnextMonday, vNow, 3, 1)
    on duplicate key update reward=reward + rewardPerWeek, updateTime=vNow;
END;

-- ----------------------------
-- Procedure structure for `proc_agent_insert_sell` END
-- ----------------------------
call proc_agent_insert_sell(now(), 1000, 188831);
-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_week` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_agent_reward_per_week;
create procedure proc_agent_reward_per_week(IN vNow datetime, IN vUid bigint, IN vWay tinyint)
BEGIN
    #@param way 1表示需要创建一个额外临时表，给其他存储过程调用
    #           0 表示不用创建额外临时表
    #计算上次领取返利的时间/注册时间 到 当前时间的充值返利
    #
    # @return
    #    rewardTG 当前可领用户充值返利，rewardTG2 下周可领用户充值返利，
    #    rewardSell1 下周可领售卖返利，rewardSell2 下下周可领售卖返利，
    #    rewardSell3 下下下周可领售卖返利，rewardSell4 下下下下周可领售卖返利

    declare lastTime1 datetime;
    declare lastMonday datetime;
    declare nextMonday datetime;
    declare nnextMonday datetime;
    declare nnnextMonday datetime;
    declare nnnnextMonday datetime;

    declare limitTime datetime;
    declare todayOfWeek int;
    declare nextMondayOfWeek int;
    declare ytUid bigint;
    declare rewardTG bigint;
    declare rewardTG2 bigint;
    declare rewardSell1 bigint;
    declare rewardSell2 bigint;
    declare rewardSell3 bigint;
    declare rewardSell4 bigint;
    declare rewardPercent1 int;

    set rewardTG = 0;
    set rewardTG2 = 0;
    set rewardSell1 = 0;
    set rewardSell2 = 0;
    set rewardSell3 = 0;
    set rewardSell4 = 0;
    select percent into rewardPercent1 from agent_award_cfg where type = 1;

    set vNow = date(vNow);
    set todayOfWeek = dayofweek(vNow);
#     select todayOfWeek;
    set nextMondayOfWeek = (2 + 7 - todayOfWeek);
    if nextMondayOfWeek > 7 then
        set nextMondayOfWeek = nextMondayOfWeek - 7;
    end if;
    set nextMonday = date_add(vNow, INTERVAL nextMondayOfWeek DAY);
    set nnextMonday = date_add(nextMonday, INTERVAL 7 DAY);
    set nnnextMonday = date_add(nextMonday, INTERVAL 14 DAY);
    set nnnnextMonday = date_add(nextMonday, INTERVAL 21 DAY);
#     select nextMonday;
    if todayOfWeek = 2 then #如果今天是周一取今天
        set limitTime = vNow;
        set lastMonday = vNow;
    else #如果今天不是周一，取最近的上次周一
        set lastMonday = date_add(nextMonday, INTERVAL -7 DAY);
        set limitTime = lastMonday;
        -- select limitTime limitTime;
        -- select lastMonday lastMonday;
    end if;

    select ytid into ytUid from yt where tgy = vUid limit 1;

    #推广用户充值返利
    select canGetTime into lastTime1
    from agent_award_log
    where uid = vUid
      and status = 2
      and type = 1
    order by createTime desc
    limit 1;
    if isnull(lastTime1) then
        select reg_tm into lastTime1 from user where uid = vUid;
#         select reg_tm from user where uid=188895;
        set lastTime1 = date(lastTime1);
    end if;

#     推广奖励1 - 可领取
    select ifnull(sum(money), 0) into rewardTG
    from pay_log
    where reid = vUid
      and result = 0
      and channel in (1, 2, 3, 6)
      and addtime >= lastTime1
      and addtime < limitTime;

#     推广奖励2 - 下周领取
    select ifnull(sum(money), 0) into rewardTG2
    from pay_log
    where reid = vUid
      and result = 0
      and channel in (1, 2, 3, 6)
      and addtime >= limitTime;

    select ifnull(reward, 0) into rewardSell1
    from agent_award_log
    where canGetTime = nextMonday
      and status = 1
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell2
    from agent_award_log
    where canGetTime = nnextMonday
      and status = 1
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell3
    from agent_award_log
    where canGetTime = nnnextMonday
      and status = 1
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell4
    from agent_award_log
    where canGetTime = nnnnextMonday
      and status = 1
      and uid = vUid;

    if vWay = 1 then
        drop temporary table if exists tmp_pay_log2;
        create temporary table tmp_pay_log2
        select floor(rewardTG * rewardPercent1 / 100)  as rewardTG,
               floor(rewardTG2 * rewardPercent1 / 100) as rewardTG2,
               rewardSell1,
               rewardSell2,
               rewardSell3,
               rewardSell4,
               lastMonday,
               vUid                                    as uid;
    end if;

    select floor(rewardTG * rewardPercent1 / 100)  as rewardTG,
           floor(rewardTG2 * rewardPercent1 / 100) as rewardTG2,
           rewardSell1,
           rewardSell2,
           rewardSell3,
           rewardSell4;

END;
-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_week` END
-- ----------------------------
# call proc_agent_reward_per_week(now(), 188895, 1);
call proc_agent_reward_per_week(now(), 188895, 0);

-- ----------------------------
-- Procedure structure for `proc_agent_reward_get` begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_agent_reward_get;
CREATE PROCEDURE proc_agent_reward_get(in vNow datetime, in vUid bigint, in vType tinyint)
exec:
BEGIN
    #     @param vType 1 充值返利 给后台调用
#                  3 金币返利 给前台代理调用

    #代理领取充值返利
    declare vLastMonday datetime;
    declare vTG bigint;
    declare vYT bigint;
    declare vSell bigint;

    if vType = 1 then
        call proc_agent_reward_per_week(vNow, vUid, 1);
        select rewardTG, rewardTG2, lastMonday into vTG,vYT,vLastMonday
        from tmp_pay_log2
        where uid = vUid;

        if vTG is null then
            select 1 as code, 0 as reward;#请求失败
            leave exec;
        end if;

        insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
            value (vUid, vTG, vNow, vLastMonday, vNow, 1, 2);
        insert agent_award_log(uid, reward, createTime, canGetTime, updateTime, type, status)
            value (vUid, vYT, vNow, vLastMonday, vNow, 2, 2);

        select 0 as code, vTG + vYT as reward;
    elseif vType = 3 then
        select sum(reward) into vSell
        from agent_award_log
        where uid = vUid
          and status = 1
          and type = 3
          and canGetTime <= vNow;
        update agent_award_log
        set status = 2
        where uid = vUid
          and status = 1
          and type = 3
          and canGetTime <= vNow;
        select 0 as code, vSell as reward;
    else
        select 1 as code, 0 as reward;#请求失败
    end if;
END;

-- ----------------------------
-- Procedure structure for `proc_agent_reward_get` END
-- ----------------------------

