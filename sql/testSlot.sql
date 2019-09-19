create table coin_change_today
(
    uid         bigint not null
        primary key,
    coin_change bigint null
);

create index coin_change_today_uid_index
    on coin_change_today (uid);

create table fee_log
(
    uid bigint   not null,
    tm  datetime not null,
    fee bigint   null comment '手续费',
    primary key (uid, tm)
);
create index fee_log_uid_index
    on fee_log (uid);
create index fee_log_tm_index
    on fee_log (tm);

insert into coin_change_today(uid, coin_change) value (?, ?)
                on duplicate key update coin_change = ?;
| args=[165272 -15000 -15000] RowsAffected error=not 1 row affected