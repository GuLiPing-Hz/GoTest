create table z_huizong_qudao_hour
(
    tm        datetime         not null,
    money     bigint default 0 null comment '距离上一次统计时间段的充值，默认一小时',
    hb        bigint default 0 null comment '距离上一次统计时间段的红包提现，默认一小时',
    hbq_hbc   bigint default 0 null comment '距离上一次统计时间段的红包场红包，默认一小时',
    hbq_lzc   bigint default 0 null comment '距离上一次统计时间段的龙珠场场红包，默认一小时',
    hbq_lzdb  bigint default 0 null comment '距离上一次统计时间段的龙珠夺宝红包，默认一小时',
    in_cnt    int    default 0 null comment '距离上一次统计时间段的登录次数，默认一小时',
    new_cnt   int    default 0 null comment '距离上一次统计时间段的注册人数，默认一小时',
    check_cnt int    default 0 null comment '距离上一次统计时间段的签到次数，默认一小时',
    platform  varchar(50) comment '渠道来源'
);
create unique index z_huizong_qudao_hour_tm_uindex
    on z_huizong_qudao_hour (tm);
alter table z_huizong_qudao_hour
    add constraint z_huizong_qudao_hour_pk
        primary key (tm);

-- ----------------------------
-- Procedure structure for `proc_huizong_by_platform` BEGIN
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_huizong_by_flavor;
CREATE PROCEDURE proc_huizong_by_flavor(in vFlavor text)
exec:
BEGIN
    #按渠道单位时间分类汇总
    declare vTmBegin datetime;
    declare vTmEnd datetime;
    declare vMoney int;
    declare vHB int;
    declare vHbqHBC int;
    declare vHbqLZC int;
    declare vHbqLZDB int;
    declare vInCnt int;
    declare vNewCnt int;
    declare vCheckCnt int;

    set vTmEnd = date_add(date(now()), interval hour(now()) hour);
    select tm into vTmBegin from z_huizong_qudao_hour order by tm desc limit 1;
    if vTmBegin is null then
        set vTmBegin = '2017-01-01';
    end if;

    select sum(s) into vMoney
    from (select uid, sum(money) as s
          from pay_log
          where addtime >= vTmBegin
            and addtime < vTmEnd
            and result = 0
            and channel in (1, 2, 3, 6)
          group by uid) a
             inner join user on a.uid = user.uid
        and flavors = vFlavor;
END;
-- ----------------------------
-- Procedure structure for `proc_huizong_by_flavor` END
-- ----------------------------
call proc_huizong_by_flavor('AOfficial');