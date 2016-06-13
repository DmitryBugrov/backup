package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBackupAgent(t *testing.T) {

	Convey("Backup Agent should work", t, func() {

		Convey("Init", func() {
			So(Init(), ShouldBeTrue)
		})

		Convey("Listening messages", nil)

	})

}
