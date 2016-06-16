package main

import (
	//	"backup/BackupAgent/cfg"
	//	"backup/BackupAgent/com"

	//	"github.com/DmitryBugrov/log"

	"testing"
	//	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const test_config_file_name = "./test/config_test.json"

func TestBackupAgent(t *testing.T) {

	Convey("Backup Agent should work", t, func() {

		Convey("Init", func() {
			So(initWithConf(test_config_file_name), ShouldBeTrue)

			Convey("Listening messages", func() {

				//	Convey("Test message 'Exit'", func() {
				So(Com.Send([]byte("Exit")), ShouldEqual, nil)
				So(ListenerMessages(), ShouldBeFalse)
				//	})

				//	Convey("Test an error message 'error message'", func() {
				So(Com.Send([]byte("error message")), ShouldEqual, nil)
				So(ListenerMessages(), ShouldBeTrue)
				//	})
			})
		})

	})

}
