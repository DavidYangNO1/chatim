package file

import (
	freeUtility "lotteryim/tcp/utility"

	ini "gopkg.in/ini.v1"
)

// app config
type AppInfo struct {
	CONN_HOST           string
	CONN_PORT           string
	CONN_TYPE           string
	CONN_HeatBeatEnable bool
}

var FlAppInfo *AppInfo

func init() {
	cfg, err := ini.Load("../conf/app.ini")
	if err != nil {
		freeUtility.Flog.Info(err.Error())
	}
	FlAppInfo = new(AppInfo)
	cfg.MapTo(FlAppInfo)
}
