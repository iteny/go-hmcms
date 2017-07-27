package sql

import (
	"database/sql"
	"fmt"
	"reflect"
)

type AuthRule struct {
	Id       int        `field:"id"`
	Url      string     `field:"url"`
	Name     string     `field:"name"`
	Pid      int        `field:"pid"`
	Isshow   int        `field:"isshow"`
	Sort     int        `field:"sort"`
	Icon     string     `field:"icon"`
	Level    int        `field:"level"`
	Children []AuthRule `field:"children"`
}
type Sqler interface {
	SelectAll(sqlstr string, args ...interface{}) ([]map[string]interface{}, error)
}
type SqlDb struct {
}

//插入
func Insert(sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := sqlitedb.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.LastInsertId()
}

//修改或删除
func Exec(sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := sqlitedb.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.RowsAffected()
}

//取一行数据，注意这类取出来的结果都是string
func SelectOne(sqlstr string, args ...interface{}) (map[string]string, error) {
	stmtOut, err := sqlitedb.Prepare(sqlstr)
	if err != nil {
		// panic(err.Error())
		return map[string]string{}, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		// panic(err.Error())
		return map[string]string{}, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		// panic(err.Error())
		return map[string]string{}, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			// panic(err.Error())
			return map[string]string{}, err
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break //get the first row only
	}
	return ret, nil
}

//取多行，注意这类取出来的结果都是string
func SelectAll(sqlstr string, args ...interface{}) ([]map[string]interface{}, error) {
	stmtOut, err := sqlitedb.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]interface{}, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]interface{}, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}

//取多行，注意这类取出来的结果都是string
func (s *AuthRule) SelectAll(sqlstr string, args ...interface{}) ([]map[string]interface{}, error) {
	stmtOut, err := sqlitedb.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]interface{}, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]interface{}, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}

func SelectStruct(strt interface{}, sqlstr string, args ...interface{}) {
	t := reflect.TypeOf(strt)
	ss := reflect.New(t)
	fmt.Println("构建：", ss)
	// fmt.Println(t.Kind())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
	// stmt, err := sqlitedb.Prepare(sqlstr)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer stmt.Close()
	// rows, err := stmt.Query(args...)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer rows.Close()
}
func (ar *AuthRule) Find(sqlstr string, args ...interface{}) *AuthRule {
	rows, err := ar.SelectAll(sqlstr, args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range rows {
		for ks, vs := range v {
			refs := reflect.ValueOf(ar).Elem()
			// ss := reflect.New(t)
			// fmt.Println("构建：", ss)
			// fmt.Println(t.Kind())
			for i := 0; i < refs.NumField(); i++ {
				fieldInfo := refs.Type().Field(i)
				fieldSet := refs.Field(i)
				// fmt.Println(t.Field(i).Tag.Get("field"))
				// t.Field(i).Tag
				if ks == fieldInfo.Tag.Get("field") {
					fieldSet.Set(reflect.ValueOf(vs))

				}
			}
		}

	}
	return ar
}
