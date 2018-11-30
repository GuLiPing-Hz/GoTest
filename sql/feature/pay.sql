alter table Buyu.wares_cfg
  drop column id;
ALTER TABLE Buyu.wares_cfg
  ADD PRIMARY KEY (wares_id);