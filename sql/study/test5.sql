# mysql 用 case when可以把数表转成横表
#
create table score
(
    rowId  int auto_increment
        primary key,
    uid    int         not null,
    score  text        not null,
    course varchar(20) null
);
INSERT INTO databasetest.score (rowId, uid, score, course)
    VALUES (1, 1, '80', 'math');
INSERT INTO databasetest.score (rowId, uid, score, course)
    VALUES (2, 2, '90', 'english');
INSERT INTO databasetest.score (rowId, uid, score, course)
    VALUES (3, 2, '100', 'math');
INSERT INTO databasetest.score (rowId, uid, score, course)
    VALUES (4, 1, '95', 'english');
#
#对于上面的这张表
use databasetest;
create view view_students_ as
select uid,
       case course
           when 'math'
               then score
           else 0 end as mathscore,
       case course
           when 'english'
               then score
           else 0 end as englishscore
from score;

create view view_students as
select uid,
       max(mathscore)    as mathscore,
       max(englishscore) as englishscore
from view_students_
group by uid;

select *
from view_students;
# 通过case把数学成绩和英语成绩列成两列，本来是在同一列中(score)

# mysql 触发器
# CREATE TRIGGER trigger_name trigger_time trigger_event ON tb_name FOR EACH ROW trigger_stmt
# trigger_name：触发器的名称
# tirgger_time：触发时机，为BEFORE或者AFTER
# trigger_event：触发事件，为INSERT、DELETE或者UPDATE
# tb_name：表示建立触发器的表名，就是在哪张表上建立触发器
# trigger_stmt：触发器的程序体，可以是一条SQL语句或者是用BEGIN和END包含的多条语句
# 所以可以说MySQL创建以下六种触发器：
# BEFORE INSERT,BEFORE DELETE,BEFORE UPDATE
# AFTER INSERT,AFTER DELETE,AFTER UPDATE

#mysql 随机数
# rand() 取值范围 0~1.0
select cast(rand() * 0.01 + 0.01 as int);