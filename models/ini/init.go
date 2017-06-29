package ini

import (
	"github.com/iteny/hmgo/ini"
	"go-hmcms/models/common"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("./ini/hmcms.ini")
	if err != nil {
		common.Log.Error(err.Error())
	}

}
