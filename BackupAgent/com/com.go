//communications
//This package provides client-server relationship

package com

import (
	"github.com/streadway/amqp"

	"github.com/DmitryBugrov/log"

	"backup/BackupAgent/cfg"
)

var (
	err error
	Log *log.Log
)

type Communications struct {
	conn *amqp.Connection
}

func (self *Communications) Init(_Log *log.Log, config *cfg.Cfg) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to com.Init")
	url := "amqp://" + config.BAcfg.MessageServerAddress + ":" + config.BAcfg.MessageServerPort
	self.conn, err = amqp.Dial(url)
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}
	return nil

}

func (self *Communications) Close() error {
	Log.Print(log.LogLevelTrace, "Enter to com.Close")
	err := self.conn.Close()
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}
	return nil
}
