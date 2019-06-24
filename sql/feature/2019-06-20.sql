alter table pay_log
    modify channel tinyint default 0 not null comment
        '支付渠道，0苹果，1微信（含第三方），2支付宝(含第三方)，3云闪付(含第三方), 4红包提现，5企业付款，6公众号充值，8话费提现';


drop view if exists view_yt_user;
create view view_yt_user as
select a.uid,
       ytid,
       nick_name            as nickname,
       avatar,
       a.tm                 as tm,
       ifnull(yuhuocur, 0)  as yuhuocur,
       ifnull(yuhuoutc, -2) as yuhuoutc,
       yuhuo,
       utc,
       checkin
from yt_user a
         inner join user b on a.uid = b.uid
         left join yt_yuhuo c on a.uid = c.uid
where ytid > 0
  and apply = 0;


