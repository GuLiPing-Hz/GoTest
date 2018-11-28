ALTER TABLE Buyu.alms_log
  ADD exp bigint DEFAULT 0 NOT NULL
comment '破产的时候的流水值，用于破产保护机制';

ALTER TABLE Buyu.alms_log
  ADD charge tinyint DEFAULT 0 NOT NULL
COMMENT '是否是充值领取破产救济金;0没充值，1充值';

-- ----------------------------
-- #当日获取破产次数，以及最近破产的经验值是否充值破产 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_select_collapse_count;
DELIMITER //
CREATE PROCEDURE proc_select_collapse_count(in uidIn bigint, in uuidIn text)
  BEGIN
    #Routine body goes here...
    declare cnt int;
    declare cnt1 int; #声明变量c
    declare cnt2 int; #声明变量c
    declare expMax bigint;
    declare chargeLast int;

    #查询uid领取破产记录
    select
      count(1),
      max(exp)
    into cnt1, expMax
    from Buyu.alms_log
    where uid = uidIn and date(tm) = current_date();

    select
      exp,
      charge
    into expMax, chargeLast
    from Buyu.alms_log
    where uid = uidIn and date(tm) = current_date()
    order by tm desc
    limit 1;

    #查询uuid领取破产记录
    select count(1)
    into cnt2
    from Buyu.alms_log
    where uuid = uuidIn and date(tm) = current_date();

    #判断哪个值最大
    if cnt1 > cnt2
    then
      set cnt = cnt1;
    else
      set cnt = cnt2;
    end if;

    if isnull(expMax) then
      set expMax = 0;
    end if;

    if isnull(chargeLast) then
      set chargeLast = 0;
    end if;

    #把最后的结果返回
    select
      cnt,
      expMax     as exp,
      chargeLast as charge;
  END
//
DELIMITER ;
-- ----------------------------
-- #当日获取破产次数，以及最近破产的经验值是否充值破产 end
-- ----------------------------
call proc_select_collapse_count(165338, '0874E0CA-AECF-4D2F-C0FA-00717FA4FAC7');


create table if not exists Buyu.pay_exp_log
(
  id      bigint PRIMARY KEY    NOT NULL AUTO_INCREMENT,
  uid     bigint                NOT NULL
  COMMENT '用户id',
  oid     bigint                NOT NULL
  COMMENT '订单号',
  paycoin int                   NOT NULL
  COMMENT '充值对应的金币数值',
  exp     bigint COMMENT '用户充值的时候对应的经验值',
  time    datetime                       DEFAULT now()
  COMMENT '写入时间'
);
ALTER TABLE pay_exp_log
  COMMENT = '充值的时候对应的用户经验值';

create index uid
  on Buyu.pay_exp_log (uid);

-- ----------------------------
-- #获取当日最近一笔充值金额 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_select_payexp_last;
DELIMITER //
CREATE PROCEDURE proc_select_payexp_last(in uidIn bigint)
  BEGIN
    #Routine body goes here...
    select
      paycoin,
      exp
    from Buyu.pay_exp_log
    where uid = uidIn and date(time) = curdate()
    order by time desc
    limit 1;
  END
//
DELIMITER ;
-- ----------------------------
-- #获取当日最近一笔充值金额 end
-- ----------------------------
# call proc_select_payexp_last(170652);