#移除vip每日签到的金币奖励。
delete
from vip_gift_cfg
where gift_type = 10
  and once = 0;

#新增vip升级一次性奖励。
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (5, 10, 100000, 1);
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (6, 10, 150000, 1);
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (7, 10, 250000, 1);
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (8, 10, 500000, 1);
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (9, 10, 1500000, 1);
INSERT INTO `Buyu`.`vip_gift_cfg` (`vip`, `gift_type`, `gift_count`, `once`)
VALUES (10, 10, 5000000, 1);

#更新充值类型
alter table wares_cfg
    modify wares_type tinyint default 1 not null comment '商品类型，1购买金币，2购买钻石，3提现，4周卡/月卡，
5限购礼包，6破产礼包，7新手，8炮倍礼包 9CDKey兑换 10vip每日特惠礼包';

#配置vip特惠礼包价格
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip5', 10, DEFAULT, DEFAULT, 6, DEFAULT, 120000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip6', 10, DEFAULT, DEFAULT, 12, DEFAULT, 240000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip7', 10, DEFAULT, DEFAULT, 30, DEFAULT, 600000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip8', 10, DEFAULT, DEFAULT, 30, DEFAULT, 630000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip9', 10, DEFAULT, DEFAULT, 50, DEFAULT, 1050000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`wares_cfg` (`wares_id`, `wares_type`, `first_pay`, `mail_giftid`, `price`, `discount`, `coin`,
                                `first_reward_coin`, `diamond`, `first_reward_diamond`, `room_ids`, `countdown`,
                                `ext_value`)
VALUES ('vip10', 10, DEFAULT, DEFAULT, 50, DEFAULT, 1100000, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT);
