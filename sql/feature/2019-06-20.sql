alter table pay_log
    modify channel tinyint default 0 not null comment
        '支付渠道，0苹果，1微信（含第三方），2支付宝(含第三方)，3云闪付(含第三方), 4红包提现，5企业付款，6公众号充值，8话费提现';

