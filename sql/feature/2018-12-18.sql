ALTER TABLE pool_log ADD bill bigint DEFAULT 0 NULL COMMENT '房间累积流水';
ALTER TABLE pool_log ADD roomWin bigint DEFAULT 0 NULL COMMENT '房间累积输赢值';

-- auto-generated definition
create table room_win_log
(
  id         bigint auto_increment
    primary key,
  rid        tinyint                            not null
  comment '房间ID',
  pool       double default '0'                 null
  comment '房间小奖池',
  dragonball bigint default '0'                 null
  comment '房间龙珠奖池',
  bill       bigint default '0'                 null
  comment '房间累积流水',
  roomWin    bigint default '0'                 null
  comment '房间累积输赢',
  type       tinyint                            not null
  comment '0 正常记录，1实时输赢清空',
  curBill    bigint default '0'                 null
  comment '当前房间时段流水',
  curRoomWin bigint default '0'                 null
  comment '当前房间时段房间输赢',
  tm         datetime default CURRENT_TIMESTAMP null
)
  comment '房间实时输赢日志';



