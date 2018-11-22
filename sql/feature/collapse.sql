ALTER TABLE Buyu.alms_log
  ADD exp bigint DEFAULT 0 NOT NULL
comment '破产的时候的流水值，用于破产保护机制';

-- ----------------------------
-- #当日获取破产次数，已经最近破产的经验值 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_select_collapse_count;
DELIMITER //
#@param uid 可以为空，为空查询整个联盟的贡献数据
CREATE PROCEDURE proc_select_collapse_count(in uidIn bigint, in uuidIn text)
  BEGIN
    #Routine body goes here...
    declare cnt int;
    declare cnt1 int; #声明变量c
    declare cnt2 int; #声明变量c
    declare expMax bigint;

    #查询uid领取破产记录
    select
      count(1),
      max(exp)
    into cnt1, expMax
    from Buyu.alms_log
    where uid = uidIn and date(tm) = current_date();

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

    #把最后的结果返回
    select
      cnt,
      expMax as exp;
  END
//
DELIMITER ;
-- ----------------------------
-- #当日获取破产次数，已经最近破产的经验值 end
-- ----------------------------
# call proc_select_collapse_count(170652, '0874E0CA-AECF-4D2F-C0FA-00717FA4FAC7')
