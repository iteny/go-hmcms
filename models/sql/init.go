package sql

import (
	"database/sql"
	"fmt"
	"go-hmcms/models/common"

	"go-hmcms/models/ini"

	_ "github.com/mattn/go-sqlite3"
)

var sqlitedb *sql.DB

func init() {
	var err error
	//sql type
	switch ini.Value("sql", "sqltype") {
	case "sqlite":
		sqlitedb, err = sql.Open("sqlite3", "./sql/hmcms.db")
		if err != nil {
			common.Log.Error(err.Error())
		} else {
			fmt.Println("自身sql开启")
		}
		// defer sqlitedb.Close()
	case "mysql":
	}

}
