SELECT * FROM Buyu.user_stat WHERE uid = 170652;
SELECT * FROM Buyu.user WHERE uid = 170652;

SELECT * FROM Buyu.yule_playerlog WHERE uid = 177775;
SELECT SUM(changeCoin) FROM Buyu.yule_playerlog WHERE uid = 177775;
SELECT * FROM user_stat WHERE uid=170652;
SELECT * FROM user_stat WHERE uid=170662;
SELECT * FROM lzdb_log WHERE uid=170652;
SELECT * FROM lzdb_log WHERE uid=170662;

# UPDATE lb_points()
UPDATE user_stat SET lb_points = ? WHERE uid = ?;
UPDATE user_stat SET lb_points = 10 WHERE uid = 100027;

