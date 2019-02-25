INSERT INTO `activities` (`id`, `title`, `activityDes`, `sub_id`, `imgUrl`, `activityUrl`, `startTime`, `endTime`, `sortNo`, `createTime`, `isDeleted`, `isEnabled`)
VALUES ('9', '精彩瞬间', '精彩瞬间', '11',
             'https://www.fanyu123.cn/active_pic/%E7%AD%89%E7%BA%A7%E4%BB%BB%E5%8A%A1%E5%B0%8F%E5%9B%BE.png',
             'https://www.fanyu123.cn/active_pic/%E7%AD%89%E7%BA%A7%E4%BB%BB%E5%8A%A1%E5%B0%8F%E5%9B%BE.png',
             '2019-01-01 00:00:00', '2020-01-01 00:00:00', '9', '2019-01-01 00:00:00', 'N', 'Y');

UPDATE `Buyu`.`props_cfg` t
SET t.`duration` = 120
WHERE t.`id` = 5;
INSERT INTO `Buyu`.`props_cfg` (`type`, `coin`, `diamond`, `duration`, `times_per_day`) VALUES (19, 0, 50, 0, 0);
INSERT INTO `Buyu`.`props_cfg` (`type`, `coin`, `diamond`, `duration`, `times_per_day`) VALUES (44, 0, 100000, 0, 0);

ALTER TABLE user_stat
  CHANGE coin_limit auto_card int NOT NULL DEFAULT 0
COMMENT '自动发炮体验卡。';
UPDATE user_stat
set auto_card = 0;

UPDATE `props_cfg`
SET `diamond` = '2'
WHERE (`id` = '3');
UPDATE `props_cfg`
SET `diamond` = '2'
WHERE (`id` = '4');


ALTER TABLE wares_cfg
  ADD ext_value int DEFAULT 0 NOT NULL
COMMENT '官方充值渠道额外的赠送金币/钻石赠送。';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 1, t.`price` = 6, t.`coin` = 60000, t.`first_reward_coin` = 60000, t.`ext_value` = 48000
WHERE t.`wares_id` LIKE '7' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 2, t.`price` = 12, t.`coin` = 120000, t.`first_reward_coin` = 120000, t.`ext_value` = 96000
WHERE t.`wares_id` LIKE '8' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 4, t.`price` = 30, t.`coin` = 300000, t.`first_reward_coin` = 300000, t.`ext_value` = 240000
WHERE t.`wares_id` LIKE '9' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 8, t.`price` = 98, t.`coin` = 980000, t.`first_reward_coin` = 980000, t.`ext_value` = 784000
WHERE t.`wares_id` LIKE '10' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 16, t.`price` = 198, t.`coin` = 1980000, t.`first_reward_coin` = 1980000, t.`ext_value` = 1584000
WHERE t.`wares_id` LIKE '11' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 32, t.`price` = 648, t.`coin` = 6480000, t.`first_reward_coin` = 6480000, t.`ext_value` = 5184000
WHERE t.`wares_id` LIKE '12' ESCAPE '#';


UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 65536, t.`price` = 6, t.`diamond` = 60, t.`first_reward_diamond` = 60,
  t.`ext_value`   = 48
WHERE t.`wares_id` LIKE '101' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 131072, t.`price` = 12, t.`diamond` = 120, t.`first_reward_diamond` = 120,
  t.`ext_value`   = 96
WHERE t.`wares_id` LIKE '102' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 262144, t.`price` = 30, t.`diamond` = 300, t.`first_reward_diamond` = 300,
  t.`ext_value`   = 240
WHERE t.`wares_id` LIKE '103' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 524288, t.`price` = 98, t.`diamond` = 980, t.`first_reward_diamond` = 980,
  t.`ext_value`   = 784
WHERE t.`wares_id` LIKE '104' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 1048576, t.`price` = 198, t.`diamond` = 1980, t.`first_reward_diamond` = 1980,
  t.`ext_value`   = 1584
WHERE t.`wares_id` LIKE '105' ESCAPE '#';
UPDATE `Buyu`.`wares_cfg` t
SET t.`first_pay` = 2097152, t.`price` = 648, t.`diamond` = 6480, t.`first_reward_diamond` = 6480,
  t.`ext_value`   = 5184
WHERE t.`wares_id` LIKE '106' ESCAPE '#';


UPDATE `growth_task_cfg`
SET `type` = '6', `value` = '10000', `tip` = '累计打鱼获得10000金币'
WHERE (`id` = '3');
UPDATE `growth_task_cfg`
SET `type` = '10', `value` = '1', `tip` = '使用双倍射速道具1次'
WHERE (`id` = '4');
UPDATE `growth_task_cfg`
SET `type` = '9', `value` = '1', `tip` = '使用双倍金币道具1次'
WHERE (`id` = '7');
UPDATE `growth_task_cfg`
SET `type` = '2', `gift_type` = '19', `tip` = '人物等级到6级'
WHERE (`id` = '8');
UPDATE `growth_task_cfg`
SET `type` = '6', `value` = '100000', `tip` = '累计打鱼获得100000金币'
WHERE (`id` = '12');
UPDATE `growth_task_cfg`
SET `type` = '1', `value` = '400', `gift_count` = '30', `tip` = '捕获400条鱼'
WHERE (`id` = '16');
UPDATE `growth_task_cfg`
SET `type` = '2', `value` = '25', `tip` = '任务等级到25级'
WHERE (`id` = '19');
UPDATE `growth_task_cfg`
SET `type` = '6', `value` = '1000000', `gift_count` = '50', `tip` = '累计打鱼获得1000000金币'
WHERE (`id` = '24');
UPDATE `growth_task_cfg`
SET `type` = '6', `value` = '2000000', `gift_count` = '100', `tip` = '累计打鱼获得2000000金币'
WHERE (`id` = '25');
UPDATE `growth_task_cfg`
SET `type` = '6', `value` = '4000000', `gift_count` = '200', `tip` = '累计打鱼获得4000000金币'
WHERE (`id` = '26');

INSERT INTO `Buyu`.`sub_act_cfg` (`activity_id`, `sub_id`, `level`, `total_cnt`, `sub_title`, `sub_content`)
VALUES (9, 1, 0, 99999999, DEFAULT, DEFAULT);
INSERT INTO `Buyu`.`sub_act_gift` (`sub_id`, `activity_id`, `gift_type`, `gift_count`) VALUES (1, 9, 19, 1);