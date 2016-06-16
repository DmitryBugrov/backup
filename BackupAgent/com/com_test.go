package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	//	goczmq "github.com/zeromq/goczmq"

	"github.com/DmitryBugrov/log"

	"backup/BackupAgent/cfg"
)

const config_file_name = "../test/config_test.json"

var (
	Com *Communications
)

func TestModuleCom(t *testing.T) {
	Convey("Test com module", t, func() {
		//Init logging
		Log = new(log.Log)

		//Init and load config
		config := new(cfg.Cfg)
		So(config.Init(Log, config_file_name), ShouldEqual, nil)

		Convey("Init com module", func() {
			//Init communications module
			Com = new(Communications)
			So(Com.Init(Log, config), ShouldEqual, nil)

			Convey("Send message Hello", func() {
				So(Com.Send([]byte("Hello")), ShouldEqual, nil)

				Convey("Receve message Hello ", func() {
					request, err := Com.Receve()
					So(err, ShouldEqual, nil)
					So(request[1], ShouldResemble, []byte("Hello"))
				})

			})

		})
		Convey("Close socket", func() {
			Com.Close()

		})
	})
}
