package tool

import (
	"fmt"
	"reflect"
	"database/sql"
	"strconv"
	"github.com/mitchellh/mapstructure"
)

type DBMgr struct {
	DbInst *sql.DB
}

// 读取数据库表到 []map[string]interface{} 中
func (DM *DBMgr) LoadTable(query string, args ...interface{}) ([]map[string]interface{}, error) {
	// 将数据填入mapUsrLv中
	stmt, err := DM.DbInst.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	} else {
		colCnt := len(cols)
		tableData := make([]map[string]interface{}, 0)
		values := make([]interface{}, colCnt)
		valuePtrs := make([]interface{}, colCnt)
		if rows != nil {
			for rows.Next() {
				for i := 0; i < colCnt; i++ {
					valuePtrs[i] = &values[i]
				}
				err = rows.Scan(valuePtrs...)
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

				if err != nil {
					return nil, err
				}
			}
		}

		return tableData, nil
	}
}

func (DM *DBMgr) GetCnt(query string, args ...interface{}) int64 {
	var cnt int64
	tableData, err := DM.LoadTable(query, args...)
	if err == nil {
		if len(tableData) > 0 {
			d := tableData[0]
			cnt, _ = strconv.ParseInt(d["cnt"].(string), 10, 64) //util.GetInt64(d, "cnt")
		}
	}
	return cnt
}

func (DM *DBMgr) SelectObject(v interface{}, query string, args ...interface{}) error {
	//util.LOG.Infof("SelectObject query=%s,args=%v", query, args)
	tableData, err := DM.LoadTable(query, args...)
	if err != nil {
		return err
	}

	if len(tableData) > 0 {
		row := tableData[0]
		if err = mapstructure.Decode(row, v); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("SelectObject error, not exists")
	}
}

func (DM *DBMgr) SelectObjects(structPtrType reflect.Type, query string, args ...interface{}) (interface{}, error) {
	tableData, err := DM.LoadTable(query, args...)
	if err != nil {
		return nil, err
	}

	kind := structPtrType.Kind()
	if kind != reflect.Ptr {
		// 必须传入struct的地址（必须是指针类型）
		return nil, fmt.Errorf("Incompatible Type: %v : need struct ptr", kind)
	} else {
		structType := structPtrType.Elem()
		structTypeKind := structType.Kind()
		if structTypeKind != reflect.Struct {
			return nil, fmt.Errorf("Incompatible Type : %v : Looking For Struct", structTypeKind)
		} else {
			sliceValue := reflect.MakeSlice(reflect.SliceOf(structPtrType), 0, 0)
			for _, row := range tableData {
				structPtr := reflect.New(structType)
				if err = mapstructure.Decode(row, structPtr.Interface()); err != nil {
					return nil, err
				}
				sliceValue = reflect.Append(sliceValue, structPtr)
			}
			s := sliceValue.Interface()
			return s, nil
		}
	}
}

// 通用的update方法
func (DM *DBMgr) Update(query string, args ...interface{}) (bool, error) {
	stmt, err := DM.DbInst.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return false, err
	} else if rowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

// Insert 通用方法
func (DM *DBMgr) Insert(query string, args ...interface{}) (int64, error) {
	stmt, err := DM.DbInst.Prepare(query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	r, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return -1, err
	} else if rowsAffected != 1 {
		return -1, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
