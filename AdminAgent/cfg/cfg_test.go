package cfg

import (
	"testing"

	"github.com/DmitryBugrov/log"
	. "github.com/smartystreets/goconvey/convey"
)

const config_file_name = "config_test.json"

var (
	config *Cfg
)

func TestAdminCfg(t *testing.T) {

	Convey("Init and load config", t, func() {
		Log = new(log.Log)
		Log.Init(log.LogLevelTrace, true, true, true)

		config = new(Cfg)
		So(config.Init(Log, config_file_name), ShouldEqual, nil)

		Convey("Check correct data", func() {
			So(config.AAcfg.Aid, ShouldEqual, "1234567890")
			So(config.AAcfg.MaxFileSegment, ShouldEqual, 1048576)
			So(config.AAcfg.MessageServerAddress, ShouldEqual, "127.0.0.1")
			So(config.AAcfg.MessageServerPort, ShouldEqual, "5672")
		})

	})

}
