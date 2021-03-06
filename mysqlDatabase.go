package wooUtils

import (
	"database/sql"
	"fmt"
	"strings"

	"errors"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlDatabase mysql数据库对象
type MysqlDatabase struct {
	connStr string
}

var dbo *sql.DB

//Query query from mysqldb by sql
func (db *MysqlDatabase) Query(sqlstr string) ([]map[string]string, error) {
	if dbo == nil {
		return nil, errors.New("Dbo is nil")
	}
	rows, err := dbo.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	return rowsToSlice(rows), nil
}

//Config set config of mysqldb instance
func (db *MysqlDatabase) Config(
	uname string,
	passwd string,
	host string,
	port string,
	database string,
	charset string) (err error) {
	db.connStr = uname + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset
	dbo, err = sql.Open("mysql", db.connStr)
	if err != nil {
		dbo = nil
		return err
	}
	return nil
}

func rowsToSlice(rows *sql.Rows) []map[string]string {

	var slice []map[string]string
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		row := make(map[string]string)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			columName := strings.ToLower(columns[i])
			row[columName] = value
		}
		if slice == nil {
			slice = []map[string]string{row}
		} else {
			slice = append(slice, row)
		}
	}
	return slice
}

var _instance *MysqlDatabase

//GetDb 获取mysql单例
func GetDb() *MysqlDatabase {
	if _instance == nil {
		_instance = new(MysqlDatabase)
	}
	return _instance
}
