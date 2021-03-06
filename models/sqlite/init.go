package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type ArCoherent struct {
	ArWhere     string
	ArWhereArgs interface{}
	ArAsc       string
	ArCols      string
}

// ORM 引擎
var x *xorm.Engine

func init() {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("sqlite3", "./sql/hmcms.db")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	} else {
		fmt.Println("xorm已经打开")
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "hm_")
	x.SetTableMapper(tbMapper)
	location, _ := time.LoadLocation("Asia/Shanghai")
	x.TZLocation = location
	// x.ShowSQL(true)
	// x.Logger().SetLevel(core.LOG_DEBUG)
	// 同步结构体与数据表
	// if err = x.Sync(new(User), new(LoginLog), new(AuthRule)); err != nil {
	// 	log.Fatalf("Fail to sync database: %v\n", err)
	// }
}

//执行sql语句
func Execute(sql string) (sql.Result, error) {
	res, err := x.Exec(sql)
	return res, err
}
