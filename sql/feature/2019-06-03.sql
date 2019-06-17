alter table gift_cfg
	add vip_for_db1 tinyint default 4 not null;

alter table gift_cfg
	add id int default 1 not null;

alter table gift_cfg
	add constraint gift_cfg_pk
		primary key (id);


-- auto-generated definition
create table user_value_props_log
(
    id        bigint auto_increment
        primary key,
    uid       bigint                                not null,
    type      int         default -1                null comment '见trello 字段说明物品类型',
    cur       int         default 0                 null comment '当前的道具最新值',
    variation int         default 0                 null comment '道具变化量 >0表示增加，小余0表示减少',
    optUid    bigint                                not null comment '操作者ID，0表示系统，其他表示用户ID，如果是sendType是红包提现（61），
这个表示订单ID；如果是是sendType是商店兑换(60)或者新年集福活动(61)，这个表示兑换ID',
    optedUid  bigint                                not null comment '被操作者用户ID，0表示自己',
    sendType  smallint(6) default -1                not null comment '参见Trello。数据库字段说明',
    time      datetime    default CURRENT_TIMESTAMP null comment '变化时间'
)
    comment '用户有价值道具变化表';

create index user_value_props_log_time_index
    on user_value_props_log (time);

create index user_value_props_log_uid_index
    on user_value_props_log (uid);

create index user_value_props_log_uid_sendType
    on user_value_props_log (sendType);

#这个版本更新需要更新 myConf.json
