create table pay_log
(
  oid       bigint  auto_increment
  comment '订单号'
    primary key,
  transid   varchar(64) charset utf8                      null
  comment '计费支付平台的交易流水号',
  tradeno   varchar(64) default ''                        not null
  comment '捕鱼平台生产的唯一订单号',
  channel   tinyint default '0'                           not null
  comment '支付渠道，1苹果，2支付宝，3微信, 4红包提现，5企业付款，6公众号充值',
  uid       bigint default '0'                            not null
  comment '用户id',
  waresid   varchar(64) charset utf8 default '0'          not null
  comment '商品id（1-6apple支付，7-12爱贝支付）',
  money     float(10, 2) default '0.00'                   not null
  comment '用户实际付款金额。元',
  result    tinyint default '-1'
  comment '订单状态值。-1默认未支付，0支付成功，1等待付款，2失败或已退款',
  can_back  tinyint default 1
  comment '是否可以退款，默认1,0表示不可退款',
  issandbox tinyint default '0'
  comment '是否是沙箱测试账号充值',
  isget     tinyint default '0'
  comment '客户端是否已经获取过（用于客户端提交友盟）',
  finish    tinyint default '0'
  comment '订单标志位。按位从低到高表示：是否已经发放金币，是否已经发送红包。比如 01 表示红包金额已扣,企业红包尚未发放。
  11 表示红包金额已扣，企业红包已成功发放',
  addtime   datetime                                      not null
  comment '订单生成时间',
  transtime datetime                                      not null
  comment '订单完成时间',
  des       text                                          null
  comment '订单状态描述',
  phone     varchar(20)                                   null
  comment '充值号码'

)
  collate = utf8mb4_unicode_ci;

CREATE INDEX pay_log_oid_index
  ON pay_log (oid DESC);
CREATE INDEX pay_log_addtime_index
  ON pay_log (addtime DESC);

ALTER TABLE user_stat
  MODIFY get_coin bigint(20) DEFAULT '0'
  COMMENT '龙珠红包个人奖池';

