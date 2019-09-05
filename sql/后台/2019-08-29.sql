-- auto-generated definition
create table agent_award_log
(
    uid        bigint(16)                           not null comment '用户id',
    reward     bigint(16) default 0 comment '返利',
    rewardSG   bigint(16) default 0 comment '返利-上供',
    createTime timestamp  default CURRENT_TIMESTAMP comment '创建时间',
    canGetTime timestamp  default CURRENT_TIMESTAMP not null comment '可领取时间',
    updateTime timestamp  default CURRENT_TIMESTAMP comment '领取时间',
    type       tinyint    default 3                 not null comment '
福利类型：
1.推广用户充值
2.鱼塘充值
3.出售金币',
    status     tinyint    default 1                 not null comment '领取状态：
0|0  从右向左，第一位表示返利是否领取，第二位表示上级返利是否已领取   0未领取。1已领取。',
    primary key (uid, canGetTime, type)
)
    comment '福利记录';


create table agent_award_cfg
(
    uid      bigint        not null,
    type     tinyint       not null comment '1 推广用户充值，2鱼塘充值，3出售金币',
    percent  int default 0 null comment '% 百分数',
    percent2 int default 0 null comment '%， 上供百分数',
    constraint agent_award_cfg_pk
        primary key (uid, type)
);

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
    declare rewardPercentSG int;
    declare todayOfWeek int;
    declare rewardPerWeek bigint;
    declare rewardPerWeekSG bigint;

    select ifnull(percent, 0), ifnull(percent2, 0) into rewardPercent,rewardPercentSG
    from agent_award_cfg
    where type = 3
      and uid = vUid;

    set vNow = date(vNow);
    set todayOfWeek = (2 + 7 - (select dayofweek(vNow)));
    if todayOfWeek > 7 then
        set todayOfWeek = todayOfWeek - 7;
    end if;
    set nextMonday = date_add(vNow, INTERVAL todayOfWeek DAY);
    set nnextMonday = date_add(nextMonday, INTERVAL 7 DAY);
    set nnnextMonday = date_add(nextMonday, INTERVAL 14 DAY);
    set nnnnextMonday = date_add(nextMonday, INTERVAL 21 DAY);

    set rewardPerWeek = floor(vGold * rewardPercent / 100 / 4);
    set rewardPerWeekSG = floor(vGold * rewardPercentSG / 100 / 4);
    insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, rewardPerWeekSG, vNow, nextMonday, vNow, 3, 0)
    on duplicate key update reward=reward + rewardPerWeek, rewardSG=rewardSG + rewardPerWeekSG, updateTime=vNow;

    insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, rewardPerWeekSG, vNow, nnextMonday, vNow, 3, 0)
    on duplicate key update reward=reward + rewardPerWeek, rewardSG=rewardSG + rewardPerWeekSG, updateTime=vNow;

    insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, rewardPerWeekSG, vNow, nnnextMonday, vNow, 3, 0)
    on duplicate key update reward=reward + rewardPerWeek, rewardSG=rewardSG + rewardPerWeekSG, updateTime=vNow;

    insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
        value (vUid, rewardPerWeek, rewardPerWeekSG, vNow, nnnnextMonday, vNow, 3, 0)
    on duplicate key update reward=reward + rewardPerWeek, rewardSG=rewardSG + rewardPerWeekSG, updateTime=vNow;

END;

-- ----------------------------
-- Procedure structure for `proc_agent_insert_sell` END
-- ----------------------------
call proc_agent_insert_sell('2019-08-04', 10000, 188967);
call proc_agent_insert_sell('2019-08-11', 10000, 188967);
call proc_agent_insert_sell('2019-08-18', 10000, 188967);
call proc_agent_insert_sell('2019-08-25', 10000, 188967);
call proc_agent_insert_sell('2019-09-01', 10000, 188967);
call proc_agent_insert_sell('2019-09-05', 10000, 188967);

call proc_agent_insert_sell('2019-08-04', 10000, 188969);
call proc_agent_insert_sell('2019-08-11', 10000, 188969);
call proc_agent_insert_sell('2019-08-18', 10000, 188969);
call proc_agent_insert_sell('2019-08-25', 10000, 188969);
call proc_agent_insert_sell('2019-09-05', 10000, 188969);

call proc_agent_insert_sell('2019-08-11', 10000, 188978);
call proc_agent_insert_sell('2019-08-18', 10000, 188978);
call proc_agent_insert_sell('2019-08-25', 10000, 188978);
call proc_agent_insert_sell('2019-09-01', 10000, 188978);
call proc_agent_insert_sell('2019-09-05', 10000, 188978);
-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_week` begin
-- ----------------------------
# 如果已经存在一个同名存储过程，那么我们移除掉
DROP PROCEDURE IF EXISTS proc_agent_reward_per_week;
create procedure proc_agent_reward_per_week(in vNow datetime, in vUid bigint,
                                            out vRewardTG bigint, out vRewardTGSG bigint,
                                            out vRewardYT bigint, out vRewardYTSG bigint,
                                            out vLastMonday datetime)
BEGIN
    #@param way 1表示需要创建一个额外临时表，给其他存储过程调用
    #           0 表示不用创建额外临时表
    #计算上次领取返利的时间/注册时间 到 当前时间的充值返利
    #
    # @return

    declare rewardTG bigint; #当前可领用户充值返利
    declare rewardTGSG bigint; #当前可领用户充值返利 - 上供
    declare rewardTG2 bigint; #下周可领用户充值返利
    declare rewardTGSG2 bigint; #下周可领用户充值返利 - 上供
    declare rewardYT bigint; #当前可领鱼塘充值返利
    declare rewardYTSG bigint; #当前可领鱼塘充值返利 - 上供
    declare rewardYT2 bigint; #下周可领鱼塘充值返利
    declare rewardYTSG2 bigint; #下周可领鱼塘充值返利 - 上供
    declare rewardSell bigint; #可领售卖返利
    declare rewardSell1 bigint; #下周可领售卖返利
    declare rewardSell2 bigint; #下下周可领售卖返利
    declare rewardSell3 bigint; #下下下周可领售卖返利
    declare rewardSell4 bigint; #下下下下周可领售卖返利
    declare rewardSellSG bigint; #可领售卖返利 - 上供
    declare rewardSellSG1 bigint; #下周可领售卖返利 - 上供
    declare rewardSellSG2 bigint; #下下周可领售卖返利 - 上供
    declare rewardSellSG3 bigint; #下下下周可领售卖返利 - 上供
    declare rewardSellSG4 bigint; #下下下下周可领售卖返利 - 上供

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

    declare rewardPercent1 int;#推广充值返利百分比
    declare rewardPercent2 int;#鱼塘充值返利百分比
    declare sgCnt int;

    set rewardTG = 0;
    set rewardTGSG = 0;
    set rewardTG2 = 0;
    set rewardTGSG2 = 0;
    set rewardYT = 0;
    set rewardYTSG = 0;
    set rewardYT2 = 0;
    set rewardYTSG2 = 0;
    set rewardSell = 0;
    set rewardSell1 = 0;
    set rewardSell2 = 0;
    set rewardSell3 = 0;
    set rewardSell4 = 0;
    set rewardSellSG = 0;
    set rewardSellSG1 = 0;
    set rewardSellSG2 = 0;
    set rewardSellSG3 = 0;
    set rewardSellSG4 = 0;
    select ifnull(percent, 0) into rewardPercent1
    from agent_award_cfg
    where type = 1
      and uid = vUid;
    select ifnull(percent, 0) into rewardPercent2
    from agent_award_cfg
    where type = 2
      and uid = vUid;

    select count(1) into sgCnt from agent_user_info where bindUid = vUid;

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

# 推广用户充值返利
    select canGetTime into lastTime1
    from agent_award_log
    where uid = vUid
      and status = 2
      and type = 1
    order by createTime desc
    limit 1;
    if isnull(lastTime1) then
        select reg_tm into lastTime1 from user where uid = vUid;
        set lastTime1 = date(lastTime1);
    end if;

# 推广奖励 - 可领取 自己的部分
    select ifnull(sum(money), 0) into rewardTG
    from pay_log
    where reid = vUid
      and result = 0
      and retype = 1
      and channel in (1, 2, 3, 6)
      and addtime >= lastTime1
      and addtime < limitTime;

# 推广奖励 - 下周领取
    select ifnull(sum(money), 0) into rewardTG2
    from pay_log
    where reid = vUid
      and result = 0
      and retype = 1
      and channel in (1, 2, 3, 6)
      and addtime >= limitTime;


# 如果金推下面有其他推广员
    if sgCnt > 0 then
# 推广奖励 - 可领取 上供部分
        select ifnull(sum(money * percent2), 0) into rewardTGSG
        from (select pay_log.reid, sum(money) as money
              from pay_log
              where result = 0
                and reid2 = vUid         #199888
                and retype = 1
                and channel in (1, 2, 3, 6)
                and addtime >= lastTime1 # '2019-08-26'
                and addtime < limitTime  #'2019-09-02'  #
              group by pay_log.reid) a
                 inner join agent_award_cfg on a.reid = agent_award_cfg.uid
        where type = 1;

# 推广奖励 - 下周领取 上供部分
        select ifnull(sum(money * percent2), 0) into rewardTGSG2
        from (select pay_log.reid, sum(money) as money
              from pay_log
              where result = 0
                and reid2 = vUid         #199888
                and retype = 1
                and channel in (1, 2, 3, 6)
                and addtime >= limitTime # '2019-08-26'
              group by pay_log.reid) a
                 inner join agent_award_cfg on a.reid = agent_award_cfg.uid
        where type = 1;
    end if;

# 鱼塘推广奖励 -可领取
    select ifnull(sum(money), 0) into rewardYT
    from pay_log
    where reid = vUid
      and result = 0
      and retype = 2
      and channel in (1, 2, 3, 6)
      and addtime >= lastTime1
      and addtime < limitTime;

# 鱼塘推广奖励 -下周领取
    select ifnull(sum(money), 0) into rewardYT2
    from pay_log
    where reid = vUid
      and result = 0
      and retype = 2
      and channel in (1, 2, 3, 6)
      and addtime >= limitTime;

    if sgCnt > 0 then
# 鱼塘推广奖励 - 可领取 上供部分
        select ifnull(sum(money * percent2), 0) into rewardYTSG
        from (select pay_log.reid, sum(money) as money
              from pay_log
              where result = 0
                and reid2 = vUid         #199888
                and retype = 2
                and channel in (1, 2, 3, 6)
                and addtime >= lastTime1 # '2019-08-26'
                and addtime < limitTime  #'2019-09-02'  #
              group by pay_log.reid) a
                 inner join agent_award_cfg on a.reid = agent_award_cfg.uid
        where type = 2;

# 鱼塘推广奖励 - 下周领取 上供部分
        select ifnull(sum(money * percent2), 0) into rewardYTSG2
        from (select pay_log.reid, sum(money) as money
              from pay_log
              where result = 0
                and reid2 = vUid         #199888
                and retype = 2
                and channel in (1, 2, 3, 6)
                and addtime >= limitTime # '2019-08-26'
              group by pay_log.reid) a
                 inner join agent_award_cfg on a.reid = agent_award_cfg.uid
        where type = 2;
    end if;

    select ifnull(sum(reward), 0) into rewardSell
    from agent_award_log
    where canGetTime < nextMonday
      and status & 1 = 0
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell1
    from agent_award_log
    where canGetTime = nextMonday
      and status & 1 = 0
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell2
    from agent_award_log
    where canGetTime = nnextMonday
      and status & 1 = 0
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell3
    from agent_award_log
    where canGetTime = nnnextMonday
      and status & 1 = 0
      and uid = vUid;
    select ifnull(reward, 0) into rewardSell4
    from agent_award_log
    where canGetTime = nnnnextMonday
      and status & 1 = 0
      and uid = vUid;

    if sgCnt > 0 then
        select ifnull(sum(rewardSG), 0) into rewardSellSG
        from agent_award_log
        where canGetTime < nextMonday
          and status & 2 = 0
          and uid in (select uid from agent_user_info where bindUid = vUid);
        select ifnull(sum(rewardSG), 0) into rewardSellSG1
        from agent_award_log
        where canGetTime = nextMonday
          and status & 2 = 0
          and uid in (select uid from agent_user_info where bindUid = vUid);
        select ifnull(sum(rewardSG), 0) into rewardSellSG2
        from agent_award_log
        where canGetTime = nnextMonday
          and status & 2 = 0
          and uid in (select uid from agent_user_info where bindUid = vUid);
        select ifnull(sum(rewardSG), 0) into rewardSellSG3
        from agent_award_log
        where canGetTime = nnnextMonday
          and status & 2 = 0
          and uid in (select uid from agent_user_info where bindUid = vUid);
        select ifnull(sum(rewardSG), 0) into rewardSellSG4
        from agent_award_log
        where canGetTime = nnnnextMonday
          and status & 2 = 0
          and uid in (select uid from agent_user_info where bindUid = vUid);
    end if;

    set vRewardTG = floor(rewardTG * rewardPercent1 / 100);
    set vRewardTGSG = floor(rewardTGSG / 100);
    set vRewardYT = floor(rewardYT * rewardPercent2 / 100);
    set vRewardYTSG = floor(rewardYTSG / 100);
    set vLastMonday = lastMonday;

    select vRewardTG                               as rewardTG,
           floor(rewardTG2 * rewardPercent1 / 100) as rewardTG2,
           vRewardTGSG                             as rewardTGSG,
           floor(rewardTGSG2 / 100)                as rewardTGSG2,
           vRewardYT                               as rewardYT,
           floor(rewardYT2 * rewardPercent2 / 100) as rewardYT2,
           vRewardYTSG                             as rewardYTSG,
           floor(rewardYTSG2 / 100)                as rewardYTSG2,
           rewardSell,
           rewardSell1,
           rewardSell2,
           rewardSell3,
           rewardSell4,
           rewardSellSG,
           rewardSellSG1,
           rewardSellSG2,
           rewardSellSG3,
           rewardSellSG4;

END;
-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_week` END
-- ----------------------------
# call proc_agent_reward_per_week(now(), 188895, 1);
call proc_agent_reward_per_week(now(), 188967, @vTG, @vTGSG, @vYT, @vTYSG, @vMonday);

-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_weekex` BEGIN
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_agent_reward_per_weekex;
CREATE PROCEDURE proc_agent_reward_per_weekex(in vNow datetime, in vUid bigint)
exec:
BEGIN
    declare vTG bigint;
    declare vYT bigint;
    declare vTGSG bigint;
    declare vYTSG bigint;
    declare vLastMonday bigint;
    call proc_agent_reward_per_week(vNow, vUid, vTG, vTGSG, vYT, vYTSG, vLastMonday);
END;
-- ----------------------------
-- Procedure structure for `proc_agent_reward_per_weekex` END
-- ----------------------------
call proc_agent_reward_per_weekex(now(), 188967);
call proc_agent_reward_per_weekex(now(), 188978);
call proc_agent_reward_per_weekex(now(), 188969);

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
    declare vTGSG bigint;
    declare vYTSG bigint;
    declare vSell bigint;
    declare vSellSG bigint;
    declare sgCnt int;

    if vType = 1 then
        call proc_agent_reward_per_week(vNow, vUid, vTG, vTGSG, vYT, vYTSG, vLastMonday);
        if vTG is null then
            select 1 as code, 0 as reward;#请求失败
            leave exec;
        end if;

        insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
            value (vUid, vTG, vTGSG, vNow, vLastMonday, vNow, 1, 2);
        insert agent_award_log(uid, reward, rewardSG, createTime, canGetTime, updateTime, type, status)
            value (vUid, vYT, vYTSG, vNow, vLastMonday, vNow, 2, 2);

        select 0 as code, vTG + vYT + vTGSG + vYTSG as reward;
    elseif vType = 3 then
# 查询可领取的售卖金币
        select ifnull(sum(reward), 0) into vSell
        from agent_award_log
        where uid = vUid
          and status & 1 = 0
          and type = 3
          and canGetTime <= vNow;
#更新自己的售卖金币领取状态
        update agent_award_log
        set status = status | 1
        where uid = vUid
          and status = 1
          and type = 3
          and canGetTime <= vNow;
#查看是否有其他返利
        select count(1) into sgCnt from agent_user_info where bindUid = vUid;
        set vSellSG = 0;
        if sgCnt > 0 then
            select ifnull(sum(rewardSG), 0) into vSellSG
            from agent_award_log
            where status & 2 = 0
              and type = 3
              and canGetTime <= vNow
              and uid in (select uid from agent_user_info where bindUid = vUid);

            update agent_award_log
            set status = status | 2
            where status & 2 = 0
              and type = 3
              and canGetTime <= vNow
              and uid in (select uid from agent_user_info where bindUid = vUid);
        end if;
        select 0 as code, vSell + vSellSG as reward;
    else
        select 1 as code, 0 as reward;#请求失败
    end if;
END;
-- ----------------------------
-- Procedure structure for `proc_agent_reward_get` END
-- ----------------------------

update agent_award_log
set status = 0
where true;
call proc_agent_reward_get(now(), 188967);
call proc_agent_reward_get(now(), 188978);
call proc_agent_reward_get(now(), 188969);
