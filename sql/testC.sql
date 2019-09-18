select sum(money) as s from pay_log where uid=165272
		and result=0 and channel in(1,2,3) and addtime>='2019-09-15'
create table flavor_cfg
(
    flavor varchar(50) not null
        primary key
);

INSERT INTO `Buyu`.`flavor_cfg` (`flavor`) VALUES ('MARKET_xw');
INSERT INTO `Buyu`.`flavor_cfg` (`flavor`) VALUES ('MARKET_xw1');
INSERT INTO `Buyu`.`flavor_cfg` (`flavor`) VALUES ('MARKET_sp');
INSERT INTO `Buyu`.`flavor_cfg` (`flavor`) VALUES ('AOfficial');

