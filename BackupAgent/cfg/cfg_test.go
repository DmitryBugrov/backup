package cfg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/DmitryBugrov/log"
)

const config_file_name = "config_test.json"

var (
	config *Cfg
)

func TestCfg(t *testing.T) {
	Convey("Cfg module should work correctly", t, func() {
		Convey("Init and load config", func() {
			//Init logging
			Log = new(log.Log)
			Log.Init(log.LogLevelTrace, true, true, true)

			config = new(Cfg)
			So(config.Init(Log, config_file_name), ShouldEqual, nil)

			Convey("Check correct data", func() {
				So(config.BAcfg.Cid, ShouldEqual, "123456789")
				So(config.BAcfg.MaxFileSegment, ShouldEqual, 1048576)
				So(config.BAcfg.MessageServerAddress, ShouldEqual, "127.0.0.1")
				So(config.BAcfg.BackupGroup[0].Path_for_backup[0], ShouldEqual, "/tmp")
			})
		})

	})
}
