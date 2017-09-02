/*
 *
 * log help class
 *
 */

package utility

import (
	"time"

	"github.com/uber-go/zap"
)

var (
	Flog zap.Logger
)

const (
	LocalTimeFormate string = "2006-01-02 15:04:05"
	//timeformate string = "freelancer-server"
)

func init() {

	Flog = zap.New(zap.NewJSONEncoder(LocalFormatter("ts")), zap.AddCaller())
}

func LocalFormatter(key string) zap.TimeFormatter {
	return zap.TimeFormatter(func(t time.Time) zap.Field {
		return zap.String(key, t.Local().Format(LocalTimeFormate))
	})
}
