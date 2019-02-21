INSERT INTO `activities` (`id`, `title`, `activityDes`, `sub_id`, `imgUrl`, `activityUrl`, `startTime`, `endTime`, `sortNo`, `createTime`, `isDeleted`, `isEnabled`)
VALUES ('9', '精彩瞬间', '精彩瞬间', '11',
             'https://www.fanyu123.cn/active_pic/%E7%AD%89%E7%BA%A7%E4%BB%BB%E5%8A%A1%E5%B0%8F%E5%9B%BE.png',
             'https://www.fanyu123.cn/active_pic/%E7%AD%89%E7%BA%A7%E4%BB%BB%E5%8A%A1%E5%B0%8F%E5%9B%BE.png',
             '2019-01-01 00:00:00', '2020-01-01 00:00:00', '9', '2019-01-01 00:00:00', 'N', 'Y');

UPDATE `Buyu`.`props_cfg` t
SET t.`duration` = 120
WHERE t.`id` = 5;
INSERT INTO `Buyu`.`props_cfg` (`type`, `coin`, `diamond`, `duration`, `times_per_day`) VALUES (19, 0, 50, 0, 0);
INSERT INTO `Buyu`.`props_cfg` (`type`, `coin`, `diamond`, `duration`, `times_per_day`) VALUES (44, 0, 20, 0, 0);

ALTER TABLE user_stat
  CHANGE coin_limit auto_card int NOT NULL DEFAULT 0
COMMENT '自动发炮体验卡。';
UPDATE user_stat
set auto_card = 0;