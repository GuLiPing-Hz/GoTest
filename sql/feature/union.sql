CREATE TABLE IF NOT EXISTS Buyu.union_cfg (
  #定义字段名 类型 默认值 主键 是否可空 自动增加
  id            BIGINT PRIMARY KEY            NOT NULL AUTO_INCREMENT,
  vip           INT                           NOT NULL
  COMMENT '创建需要的vip等级',
  level         INT                           NOT NULL
  COMMENT '创建需要的等级',
  gold          BIGINT                        NOT NULL
  COMMENT '创建需要的金币',
  rewardCreater BIGINT                        NOT NULL
  COMMENT '新用户第一次进入联盟的奖励,这是发放给创建者的',
  rewardUser    BIGINT                        NOT NULL
  COMMENT '新用户第一次进入联盟的奖励,发放给用户的'
)
  COMMENT '推广联盟创建配置'
  DEFAULT CHARSET = utf8mb4;
#流水计算？？，玩家当天最后一秒退出，今日的流水怎么计算，是否返利给联盟

CREATE TABLE IF NOT EXISTS Buyu.union_log (
  unionId  bigint PRIMARY KEY            NOT NULL
  COMMENT '推广联盟ID',
  name     VARCHAR(255) CHARSET utf8mb4  NOT NULL
  COMMENT '联盟名称',
  uid      BIGINT                        NOT NULL
  COMMENT '联盟创建者(会长)',
  time     DATETIME DEFAULT NOW()
  COMMENT '创建时间',
  status   tinyint default 1             not null
  comment '联盟状态：1正常；0解散',
  integral bigint   default 0
  comment '联盟积分'
)
  COMMENT '推广联盟记录表'
  DEFAULT CHARSET = utf8mb4;

create table if not exists Buyu.union_member (
  id      bigint primary key not null auto_increment,
  unionId bigint             not null
  comment '联盟ID',
  uid     bigint             not null
  comment '用户iD',
  time    datetime                    default now()
  comment '加入时间'
)
  COMMENT '推广联盟成员表'
  DEFAULT CHARSET = utf8mb4;

create table if not exists Buyu.union_member_log (
  id           bigint primary key not null auto_increment,
  unionId      bigint             not null
  comment '联盟ID',
  uid          bigint             not null
  comment '用户id',
  time         datetime                    default now()
  comment '贡献时间',
  contribution bigint             not null
  comment '贡献流水'
)
  COMMENT '推广联盟成员每日贡献'
  DEFAULT CHARSET = utf8mb4;
create index unionId
  on Buyu.union_member_log (unionId);
create index uid
  on Buyu.union_member_log (uid);
create index time
  on Buyu.union_member_log (time);

create table if not exists Buyu.union_member_inout (
  id      bigint primary key not null auto_increment,
  unionId bigint             not null
  comment '联盟ID',
  uid     bigint             not null
  comment '用户iD',
  time    datetime                    default now()
  comment '加入/离开时间',
  type    tinyint            not null
  comment '类型：0离开联盟，1加入联盟'
)
  COMMENT '推广联盟成员进出联盟记录表'
  DEFAULT CHARSET = utf8mb4;
create index uid
  on Buyu.union_member_inout (uid);

-- ----------------------------
-- #添加用户每日贡献 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_add_user_contirbution;
DELIMITER //
CREATE PROCEDURE proc_add_user_contirbution(in uidIn bigint, in unionIdIn bigint, in valIn bigint)
  BEGIN
    #Routine body goes here...
    declare total bigint;
    declare con bigint;

    select contribution
    into con
    from Buyu.union_member_log
    where uid = uidIn and unionId = unionIdIn and curdate() = date(time);

    if con is null
    then
      set total = valIn;
      insert into Buyu.union_member_log (unionId, uid, contribution) values (unionIdIn, uidIn, total);
    else
      set total = valIn + con;
      update Buyu.union_member_log
      set contribution = total
      where uid = uidIn and unionId = unionIdIn and curdate() = date(time);
    end if;


  END
//
DELIMITER ;
-- ----------------------------
-- #添加用户每日贡献 end
-- ----------------------------

-- ----------------------------
-- #查询帮会列表 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_select_unions;
DELIMITER //
CREATE PROCEDURE proc_select_unions()
  BEGIN
    #Routine body goes here...

    #mysql 如果想返回多多行数据，直接select就行，返回某个数据需要放入到out参数中
    select
      name,
      Buyu.union_log.unionId,
      Buyu.user.nick_name,
      Buyu.union_log.uid,
      count(1) as cnt
    #     into nameOut, unionIdOut, uidOut, countOut
    from Buyu.union_log
      inner join Buyu.user on Buyu.union_log.uid = Buyu.user.uid
      inner join Buyu.union_member on Buyu.union_member.unionId = Buyu.union_log.unionId
    group by Buyu.union_member.unionId
    order by cnt desc;

  END
//
DELIMITER ;
-- ----------------------------
-- #查询帮会列表 end
-- ----------------------------

# call proc_select_unions();

-- ----------------------------
-- #查询帮会列表 begin
-- ----------------------------
DROP PROCEDURE IF EXISTS proc_select_contributions;
DELIMITER //
#@param uid 可以为空，为空查询整个联盟的贡献数据
CREATE PROCEDURE proc_select_contributions(in unionIdIn bigint, in uidIn bigint)
  BEGIN
    #Routine body goes here...
    #     创建临时表，进行查询操作
    if uidIn is null
    then
      DROP TEMPORARY TABLE IF EXISTS Buyu.tmp_table;
      CREATE TEMPORARY TABLE Buyu.tmp_table
          SELECT
            contribution as yesterday,
            uid,
            unionId
          FROM Buyu.union_member_log
          where date(time) = date(DATE_SUB(NOW(), INTERVAL '1 0:0:0' DAY_SECOND)) and unionId = unionIdIn;
      #     select *
      #     from Buyu.tmp_table;


      #mysql 如果想返回多多行数据，直接select就行，返回某个数据需要放入到out参数中
      select
        Buyu.user.nick_name,
        Buyu.union_member_log.uid,
        yesterday,
        sum(Buyu.union_member_log.contribution)
      from Buyu.union_member_log
        inner join Buyu.user on Buyu.union_member_log.uid = Buyu.user.uid
        left join Buyu.tmp_table on Buyu.union_member_log.uid = Buyu.tmp_table.uid
      where Buyu.union_member_log.unionId = unionIdIn
      group by Buyu.union_member_log.uid;


    else
      #       select 1;

      DROP TEMPORARY TABLE IF EXISTS Buyu.tmp_table;
      CREATE TEMPORARY TABLE Buyu.tmp_table
          SELECT
            contribution as yesterday,
            uid,
            unionId
          FROM Buyu.union_member_log
          where date(time) = date(DATE_SUB(NOW(), INTERVAL '1 0:0:0' DAY_SECOND))
                and unionId = unionIdIn and uid = uidIn;

      #mysql 如果想返回多多行数据，直接select就行，返回某个数据需要放入到out参数中
      select
        Buyu.user.nick_name,
        Buyu.union_member_log.uid,
        yesterday,
        sum(Buyu.union_member_log.contribution) as total
      from Buyu.union_member_log
        inner join Buyu.user on Buyu.union_member_log.uid = Buyu.user.uid
        left join Buyu.tmp_table on Buyu.union_member_log.uid = Buyu.tmp_table.uid
      where Buyu.union_member_log.unionId = unionIdIn and Buyu.union_member_log.unionId = unionIdIn
            and Buyu.union_member_log.unionId = unionIdIn and Buyu.union_member_log.uid = uidIn
      group by Buyu.union_member_log.uid;

    end if;

    -- 删除临时表
    DROP TEMPORARY TABLE IF EXISTS Buyu.tmp_table;
  END
//
DELIMITER ;
-- ----------------------------
-- #查询帮会列表 end
-- ----------------------------

# call proc_select_contributions(10000, 177851);


