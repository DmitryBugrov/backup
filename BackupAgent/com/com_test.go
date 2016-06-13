package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	//	goczmq "github.com/zeromq/goczmq"

	"github.com/DmitryBugrov/log"

	"backup/BackupAgent/cfg"
)

const config_file_name = "./config_test.json"

var (
	Com *Communications
)

func TestModuleCom(t *testing.T) {

	//Init logging
	Log = new(log.Log)

	//Init and load config
	config := new(cfg.Cfg)
	err := config.Init(Log, config_file_name)
	if err != nil {
		Log.Print(log.LogLevelError, "No configuration file loaded: ", config_file_name)

	}

	Convey("Init com module", t, func() {
		//Init communications module
		Com = new(Communications)
		So(Com.Init(Log, config), ShouldEqual, nil)

	})

	Convey("Send message Hello", t, func() {
		So(Com.SendToAdmin([]byte("Hello")), ShouldEqual, nil)
		reply, err := Com.dealer.RecvMessage()
		So(err, ShouldEqual, nil)
		So(reply[0], ShouldEqual, []byte("Hello"))

	})

	Convey("Close socket", t, func() {
		Com.Close()

	})

}
