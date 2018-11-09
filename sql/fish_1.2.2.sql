#捕鱼数据库更新SQL
#v1.2.2 sql
CREATE TABLE IF NOT EXISTS ylc_cfg (
  #定义字段名 类型 默认值 主键 是否可空 自动增加
  id         BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT
  COMMENT '自增字段',
  gameId     INT COMMENT '游戏ID 1捕鱼 2英雄传 3渔乐',
  vipLimit   INT COMMENT 'vip等级限制',
  levelLimit INT COMMENT '等级限制'
)
  COMMENT '游戏进入的配置信息'
  DEFAULT CHARSET = utf8mb4;
INSERT INTO ylc_cfg (gameId, vipLimit, levelLimit) VALUES (2, 2, 0);
INSERT INTO ylc_cfg (gameId, vipLimit, levelLimit) VALUES (3, 0, 1);


