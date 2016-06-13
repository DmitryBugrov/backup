//communications
//This package provides client-server relationship

package com

import (
	goczmq "github.com/zeromq/goczmq"

	"github.com/DmitryBugrov/log"

	"backup/BackupAgent/cfg"
)

var (
	err error
	Log *log.Log
)

type Communications struct {
	dealer *goczmq.Sock
}

func (c *Communications) Init(_Log *log.Log, config *cfg.Cfg) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to com.Init")
	url := "tcp://" + config.BAcfg.MessageServerAddress + ":" + config.BAcfg.MessageServerPort
	c.dealer, err = goczmq.NewDealer(url)
	c.dealer.SetRcvtimeo(config.BAcfg.TimeoutForResponse)

	return err

}

func (c *Communications) Close() {
	Log.Print(log.LogLevelTrace, "Enter to com.Close")
	c.dealer.Destroy()
}

func (c *Communications) SendToAdmin(msg []byte) error {
	err = c.dealer.SendFrame(msg, goczmq.FlagNone)
	return err
}
