package mybase

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mohae/deepcopy"
	"mybase/netdata"
	"reflect"
	"time"
)

type DBMgrBase struct {
	DbInst    *sql.DB
	RedisInst *redis.Client
}

const (
	//redis 缓存信息的版本号，从数据库第一次取出来放进去是0
	VerInRedis = "ver_in_redis"

	//redis写入标志位 请求写入的时候 incrBy 1,如果返回1表示可以写入，大于1表示已有其他服务器正在写入，要么等待要么直接返回
	WritingInRedis = "writing_in_redis"
)

func (imp *DBMgrBase) CheckDBConnect() error {
	return imp.CheckDBConnectEx(false)
}

func (imp *DBMgrBase) CheckDBConnectEx(withRedis bool) error {
	//这里成功，并不能代表真的成功。。。，可能这个数据库服务器压根访问不到
	//所以我们这里尝试ping一下
	fmt.Println("Check DB...", imp.DbInst)
	err := imp.DbInst.Ping()
	if err != nil {
		fmt.Println("Check DB error=", err)
		return err
	}
	fmt.Println("Check DB OK")

	if withRedis && imp.RedisInst != nil {
		_, err := imp.RedisInst.Ping().Result()
		if err != nil {
			fmt.Println("Check Redis error=", err)
			return err
		}
		fmt.Println("Check Redis OK")
	}
	return nil
}

// 读取多个数据库表到 [][]map[string]interface{} 中
func (imp *DBMgrBase) LoadTableEx(query string, args ...interface{}) ([][]map[string]interface{}, error) {
	// 将数据填入mapUsrLv中
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("query=%s | args=%v, Prepare error=%s", query, args, err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		E("query=%s | args=%v, Query error=%s", query, args, err.Error())
		return nil, err
	}
	defer rows.Close()

	result := make([][]map[string]interface{}, 0)
	dealResult := func() error {
		cols, err := rows.Columns()
		if err != nil {
			E("query=%s | args=%v, Columns error=%s", query, args, err.Error())
			return err
		}

		colCnt := len(cols)
		tableData := make([]map[string]interface{}, 0)
		values := make([]interface{}, colCnt)
		valuePtrs := make([]interface{}, colCnt)
		for rows.Next() {
			for i := 0; i < colCnt; i++ {
				valuePtrs[i] = &values[i]
			}
			err = rows.Scan(valuePtrs...)
			if err != nil {
				E("query=%s | args=%v, Scan error=%s", query, args, err.Error())
				break
			}

			entry := make(map[string]interface{})
			for i, col := range cols {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
			}
			tableData = append(tableData, entry)
		}
		result = append(result, tableData)
		return nil
	}

	err = dealResult()
	if err != nil {
		return nil, err
	}
	for rows.NextResultSet() {
		err = dealResult()
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// 读取单个数据库表到 []map[string]interface{} 中
func (imp *DBMgrBase) LoadTable(query string, args ...interface{}) ([]map[string]interface{}, error) {
	// 将数据填入mapUsrLv中
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("query=%s | args=%v, Prepare error=%s", query, args, err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		E("query=%s | args=%v, Query error=%s", query, args, err.Error())
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		E("query=%s | args=%v, Columns error=%s", query, args, err.Error())
		return nil, err
	}

	col_cnt := len(cols)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, col_cnt)
	valuePtrs := make([]interface{}, col_cnt)
	for rows.Next() {
		for i := 0; i < col_cnt; i++ {
			valuePtrs[i] = &values[i]
		}
		err = rows.Scan(valuePtrs...)
		if err != nil {
			E("query=%s | args=%v, Scan error=%s", query, args, err.Error())
			break
		}

		entry := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

func (imp *DBMgrBase) GetCnt(query string, args ...interface{}) int64 {
	cnt, _ := imp.GetCntEx(query, args...)
	return cnt
}

func (imp *DBMgrBase) GetCntEx(query string, args ...interface{}) (int64, bool) {
	tableData, err := imp.LoadTable(query, args...)
	if err != nil || len(tableData) == 0 {
		return 0, false
	}

	d := tableData[0]
	data := netdata.NetData(d)
	cnt, _ := data.GetInt64("cnt")
	return cnt, true
}

func (imp *DBMgrBase) GetSum(query string, args ...interface{}) int64 {
	var sum int64
	tableData, err := imp.LoadTable(query, args...)
	if err == nil {
		if len(tableData) > 0 {
			d := tableData[0]
			data := netdata.NetData(d)
			sum, _ = data.GetInt64("s")
		}
	}
	return sum
}

func (imp *DBMgrBase) GetSumFloat64(query string, args ...interface{}) float64 {
	var sum float64
	tableData, err := imp.LoadTable(query, args...)
	if err == nil {
		if len(tableData) > 0 {
			d := tableData[0]
			data := netdata.NetData(d)
			sum, _ = data.GetFloat64("s")
		}
	}
	return sum
}

// the v must be a pointer to a map or struct.
func (imp *DBMgrBase) SelectObject(v interface{}, query string, args ...interface{}) error {
	return imp.SelectObjectWithLog(true, v, query, args...)
}

/**
@param v 必须是数据结构指针.
*/
func (imp *DBMgrBase) SelectObjectWithLog(log bool, v interface{}, query string, args ...interface{}) error {
	dataType := reflect.TypeOf(v) //获取数据类型
	if dataType.Kind() != reflect.Ptr {
		E("query=%s | args=%v, Kind error=need Ptr", query, args)
		return ErrParam
	}

	tableData, err := imp.LoadTable(query, args...)
	if err != nil {
		return err
	}

	if len(tableData) > 0 {
		row := tableData[0]
		if err = Decode(row, v); err != nil {
			E("query=%s | args=%v, DecodePath error=%s", query, args, err.Error())
			return err
		} else {
			return nil
		}
	} else {
		if log {
			W("query=%s | args=%v, error=No Data", query, args)
		}
		return ErrNoData
	}
}

/**
跟SelectObject比，这个SelectObjectsEx是返回一个数组的。

@param v 必须是存放map或是struct的数组的指针.
*/
func (imp *DBMgrBase) SelectObjectsEx(v interface{}, query string, args ...interface{}) error {
	tableData, err := imp.LoadTable(query, args...)
	if err != nil {
		return err
	}

	if len(tableData) > 0 {
		if err = Decode(tableData, v); err != nil {
			E("query=%s | args=%v, DecodePath error=%s", query, args, err.Error())
			return err
		} else {
			//fmt.Printf("SelectObjectsEx DecodePath obj=%v\n", v)
			return nil
		}
	} else {
		W("query=%s | args=%v, error=No Data", query, args)
		return ErrNoData
	}
}

// 通用的update方法
func (imp *DBMgrBase) Update(query string, args ...interface{}) bool {
	//I("Update,query=[%s] args=%v", query, args)
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("Update=%s | args=%v Prepare error=%s", query, args, err.Error())
		return false
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		E("Update=%s | args=%v Exec error=%s", query, args, err.Error())
		return false
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		E("Update=%s | args=%v RowsAffected error=%s", query, args, err.Error())
		return false
	} else if rowsAffected <= 0 {
		W("Update=%s | args=%v RowsAffected is 0", query, args)
		return true
	}
	return true
}

// 通用的update方法
func (imp *DBMgrBase) UpdateNoWarn(query string, args ...interface{}) bool {
	//I("Update,query=[%s] args=%v", query, args)
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("UpdateNoWarn=%s | args=%v Prepare error=%s", query, args, err.Error())
		return false
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		E("UpdateNoWarn=%s | args=%v Exec error=%s", query, args, err.Error())
	}
	_, err = r.RowsAffected()
	if err != nil {
		E("UpdateNoWarn=%s | args=%v RowsAffected error=%s", query, args, err.Error())
		return false
	}
	return true
}

// Insert 通用方法
func (imp *DBMgrBase) Insert(query string, args ...interface{}) (int64, error) {
	insertId, err, _ := imp.InsertEx(query, args...)
	return insertId, err
}

func (imp *DBMgrBase) InsertEx(query string, args ...interface{}) (int64, error, bool) {
	return imp.InsertExWithLastId(true, query, args...)
}

func (imp *DBMgrBase) InsertExWithLastId(lastId bool, query string, args ...interface{}) (int64, error, bool) {
	//I("enter Insert,query=%s,args=%v", query, args)
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("InsertEx=%s | args=%v Prepare error=%s", query, args, err.Error())
		return -1, err, false
	}

	defer stmt.Close()
	r, err := stmt.Exec(args...)
	if err != nil {
		E("InsertEx=%s | args=%v Exec error=%s", query, args, err.Error())
		return -1, err, false
	}
	n, err := r.RowsAffected()
	if err != nil {
		E("InsertEx=%s | args=%v RowsAffected error=%s", query, args, err.Error())
		return -1, err, false
	}
	if lastId {
		id, err := r.LastInsertId()
		if err != nil {
			E("InsertEx=%s | args=%v LastInsertId error=%s", query, args, err.Error())
			return -1, err, false
		}
		return id, nil, n > 0
	}
	return -1, nil, n > 0
}

//执行存储过程
func (imp *DBMgrBase) CallExec(query string, args ...interface{}) (int64, error) {
	stmt, err := imp.DbInst.Prepare(query)
	if err != nil {
		E("CallExec=%s | args=%v Prepare error=%s", query, args, err.Error())
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		E("CallExec=%s | args=%v Exec error=%s", query, args, err.Error())
		return 0, err
	}
	return result.LastInsertId()
}

/**
@param v 必须是指针类型 *Struct1{}
*/
func (imp *DBMgrBase) CallQuery(v interface{}, query string, args ...interface{}) error {
	return imp.SelectObject(v, query, args...)
}

/**
@param v 必须是指针类型 //*[]Struct1{}
*/
func (imp *DBMgrBase) CallQuerys(v interface{}, query string, args ...interface{}) error {
	return imp.SelectObjectsEx(v, query, args...)
}

/**
存储过程，返回多个数据表的信息,每个表信息需要单独存放。

多张表返回，每张表含有多行数据
@v 存储数据结构指针，每个位置对应了存储过程返回的指定位置的表结构信息 是一个二维数组结构指针 类似  *[*[]Struct1,*[]Struct2,*[]Struct3...]

如果只需要取第一张表的数据，可以直接调用CallQuery或CallQuerys
*/
func (imp *DBMgrBase) CallQueryResultSets(v []interface{}, query string, args ...interface{}) error {
	resultSets, err := imp.LoadTableEx(query, args...)
	if err != nil {
		return err
	}

	for i := range v {
		if i < len(resultSets) {
			if err = Decode(resultSets[i], v[i]); err != nil {
				E("query=%s | args=%v, Decode error=%s", query, args, err.Error())
				return err
			}
		}
	}
	return nil
}

//多张表返回，每张表只取一行数据
//@v 是一个数组结构指针 类似  *[*Struct1,*Struct2,*Struct3...]
func (imp *DBMgrBase) CallQueryResultSetsOnlyFirst(v []interface{}, query string, args ...interface{}) error {
	resultSets, err := imp.LoadTableEx(query, args...)
	if err != nil {
		return err
	}

	for i := range v {
		if i < len(resultSets) {
			if len(resultSets[i]) == 0 {
				return ErrNoData
			}
			if err = Decode(resultSets[i][0], v[i]); err != nil {
				E("query=%s | args=%v, Decode error=%s", query, args, err.Error())
				return err
			}
		}
	}
	return nil
}

/**
****************************************************************************************
redis opt
*/

func (imp *DBMgrBase) RedisExists(key string) bool {
	if imp.RedisInst == nil {
		panic("Redis Instance is nil")
	}

	sc := imp.RedisInst.Exists(key)
	//D("RedisExists sc=%v", sc)
	exits, _ := sc.Result()
	return exits == 1
}

func (imp *DBMgrBase) RedisDel(key string) bool {
	ok, _ := imp.RedisInst.Del(key).Result()
	return ok == 1
}

/**
@expiration -1 永不失效
*/
func (imp *DBMgrBase) RedisSetEx(key string, value interface{}, expiration time.Duration) (bool, error) {
	ok, err := imp.RedisInst.Set(key, value, expiration).Result()
	return ok == "OK", err
}

func (imp *DBMgrBase) RedisSet(key string, value interface{}) (bool, error) {
	return imp.RedisSetEx(key, value, -1)
}

func (imp *DBMgrBase) RedisGet(key string) (string, error) {
	if !imp.RedisExists(key) {
		return "", ErrNoData
	}

	return imp.RedisInst.Get(key).Result()
}

func (imp *DBMgrBase) RedisIncrBy(key string, incr int64) (int64, error) {
	return imp.RedisInst.IncrBy(key, incr).Result()
}

/**
多个服务器分布式，读取redis中的标志位
*/
func (imp *DBMgrBase) RedisIncrByFlagAdd(key string) (int64, error) {
	return imp.RedisIncrBy(key, 1)
}
func (imp *DBMgrBase) RedisIncrByFlagCheck(key string) (int64, error) {
	return imp.RedisIncrBy(key, 0)
}
func (imp *DBMgrBase) RedisIncrByFlagRelease(key string) (int64, error) {
	return imp.RedisIncrBy(key, -1)
}

/**
不存在的key或者field，redis中默认存0
*/
func (imp *DBMgrBase) RedisHIncrBy(key, field string, incr int64) (int64, error) {
	return imp.RedisInst.HIncrBy(key, field, incr).Result()
}

/**
多个服务器分布式，读取redis中的标志位
*/
func (imp *DBMgrBase) RedisHIncrByFlagAdd(key string) (int64, error) {
	return imp.RedisHIncrBy(key, WritingInRedis, 1)
}
func (imp *DBMgrBase) RedisHIncrByFlagCheck(key string) (int64, error) {
	return imp.RedisHIncrBy(key, WritingInRedis, 0)
}
func (imp *DBMgrBase) RedisHIncrByFlagRelease(key string) (int64, error) {
	if !imp.RedisExists(key) { //减的时候我们需要检查这个key还是否有效。
		return 0, ErrNoData
	}

	return imp.RedisHIncrBy(key, WritingInRedis, -1)
}
func (imp *DBMgrBase) RedisHIncrByGetVer(key string) (int64, error) {
	return imp.RedisHIncrBy(key, VerInRedis, 0)
}

/**
正常读取redis的数据，不管是否有其他服务器正在写入
*/
func (imp *DBMgrBase) RedisHGetAll(key string, dataPtr interface{}) error {
	sc := imp.RedisInst.HGetAll(key)
	//D("RedisHGetAll sc=%v", sc)
	if sc.Err() != nil {
		if sc.Err() == redis.Nil {
			return ErrNoData
		}
		E("GetHM err=%v", sc.Err())
		return sc.Err()
	}
	return DecodeRedis(sc.Val(), dataPtr)
}

/**
正常读取redis的数据，如果有其他服务器正在写入，那么直接返回错误
*/
func (imp *DBMgrBase) RedisHGetAllEx(key string, dataPtr interface{}) error {
	//这里，即使redis中没有该key，也会返回一个空的map,所以要用exists判断一下
	if !imp.RedisExists(key) { //在检查写入标志之前优先检查该值是否存在于redis中。
		return ErrNoData
	}

	wir, _ := imp.RedisHIncrByFlagCheck(key)
	if wir != 0 {
		return ErrIsWriting
	}

	return imp.RedisHGetAll(key, dataPtr)
}

/**
正常读取redis的数据，如果有其他服务器正在写入，那么等待写完毕，获取最新的数据
@tryCnt 重试次数，如果为-1表示永远等待
*/
func (imp *DBMgrBase) RedisHGetAllExLoop(key string, dataPtr interface{}, tryCnt int, force bool) error {
	for {
		err := imp.RedisHGetAllEx(key, dataPtr)
		if err == ErrIsWriting {
			if tryCnt >= 0 {
				if tryCnt == 0 {
					if force {
						imp.RedisHGetAll(key, dataPtr)
						return nil
					}
					return ErrTryMax
				}
				tryCnt--
			}
			time.Sleep(time.Microsecond * 50)
			continue
		}
		return err
	}
}

/**
正常写入redis的数据，不管是否有其他服务器正在写入

@注意
谨慎使用该函数，该函数会导致之前的缓存数据被改写！！！
@expire 单位秒，传入负数表示永不过期或者不想改变原本的有效时间
*/
func (imp *DBMgrBase) RedisHMSet(key string, data interface{}, expire time.Duration) error {
	var dataMap map[string]interface{}
	err := Decode(data, &dataMap)
	if err != nil {
		return err
	}
	sc := imp.RedisInst.HMSet(key, dataMap)
	//D("RedisHMSet sc=%v", sc)

	if sc.Err() == nil && expire >= 0 {
		imp.RedisExpire(key, expire)
	}

	return sc.Err()
}

/**
正常写入redis的数据，如果其他服务器正在写入，那么直接返回错误,如果缓存的信息比本地版本高也返回错误
*/
func (imp *DBMgrBase) RedisHMSetEx(key string, data interface{}, ver int64, ignoreVer bool, expire time.Duration) error {
	wflag, _ := imp.RedisHIncrByFlagAdd(key)
	defer imp.RedisHIncrByFlagRelease(key)

	if wflag > 1 {
		return ErrIsWriting
	}

	if !ignoreVer {
		cur, _ := imp.RedisHIncrByGetVer(key)
		if cur > ver {
			return ErrDataOld
		}
	}

	return imp.RedisHMSet(key, data, expire)
}

type DealWithConflict func(oldPtr, newPtr interface{}) (interface{}, int64)

/**
正常写入redis的数据，如果其他服务器正在写入或者缓存中的信息版本比较高，那么我们等待写完毕把他取出来

@dataPtr 需要数据结构指针
@tryCnt 重试次数 -1表示永远尝试直到成功为止
*/
func (imp *DBMgrBase) RedisHMSetExLoop(key string, dataPtr interface{}, ver int64, expire time.Duration,
	hook DealWithConflict, tryCnt int) error {

	if tryCnt >= 0 {
		if tryCnt == 0 {
			return ErrTryMax
		}
		tryCnt--
	}

	for {
		err := imp.RedisHMSetEx(key, dataPtr, ver, false, expire)
		if err == ErrIsWriting || err == ErrDataOld {
			if hook == nil {
				return ErrAbort
			}

			newPtr := deepcopy.Copy(dataPtr)                    //深拷贝一个备份
			err = imp.RedisHGetAllExLoop(key, newPtr, 3, false) //默认尝试3次获取最新的数据
			if err != nil {
				return err
			}

			result, newVer := hook(dataPtr, newPtr)
			if result == nil {
				return ErrAbort
			}

			return imp.RedisHMSetExLoop(key, result, newVer, expire, hook, tryCnt)
		}
		return err
	}
}

func (imp *DBMgrBase) RedisSAdd(key string, members ...interface{}) error {
	return imp.RedisInst.SAdd(key, members...).Err()
}

func (imp *DBMgrBase) RedisSRem(key string, members ...interface{}) error {
	return imp.RedisInst.SRem(key, members...).Err()
}

func (imp *DBMgrBase) RedisSMembers(key string) ([]string, error) {
	if !imp.RedisExists(key) {
		return nil, ErrNoData
	}

	return imp.RedisInst.SMembers(key).Result()
}

func (imp *DBMgrBase) RedisSCard(key string) int64 {
	len, _ := imp.RedisInst.SCard(key).Result()
	return len
}

/**
设置过期时间
*/
func (imp *DBMgrBase) RedisExpire(key string, duration time.Duration) {
	imp.RedisInst.Expire(key, duration)
}

func (imp *DBMgrBase) RedisExpireAt(key string, tm time.Time) {
	imp.RedisInst.ExpireAt(key, tm)
}

/**
@time.Duration -1表示永不过期，-2表示已经过期，大于0表示还有X秒过期。
*/
func (imp *DBMgrBase) RedisTTL(key string) (time.Duration, error) {
	return imp.RedisInst.TTL(key).Result()
}

func (imp *DBMgrBase) RedisPersist(key string) {
	imp.RedisInst.Persist(key)
}

func (imp *DBMgrBase) RedisKeys(pattern string) ([]string, error) {
	return imp.RedisInst.Keys(pattern).Result()
}
