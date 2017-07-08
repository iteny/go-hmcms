package ini

import (
	"go-hmcms/models/common"

	"github.com/go-ini/ini"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("./ini/hmcms.ini")
	if err != nil {
		common.Log.Error(err.Error())
	}

}
