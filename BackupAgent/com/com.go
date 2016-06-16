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
	dealer *goczmq.Sock //struct for send message
	router *goczmq.Sock //struct for receve message
}

//Initialization of package, setup up log, server url, timeout
//return an error if goczmq.NewDealer return any errors
func (c *Communications) Init(_Log *log.Log, config *cfg.Cfg) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to com.Init")
	url := "tcp://" + config.BAcfg.MessageServerAddress + ":" + config.BAcfg.MessageServerPort
	c.dealer, err = goczmq.NewDealer(url)
	c.dealer.SetRcvtimeo(config.BAcfg.TimeoutForResponse)
	if err != nil {
		Log.Print(log.LogLevelError, "Error creating NewDealer: ")
		return err
	}
	// Create a router socket and bind it to port.
	endpoint := "tcp://*:" + config.BAcfg.MessageServerPort
	Log.Print(log.LogLevelTrace, "Creating endpoint: ", endpoint)
	c.router, err = goczmq.NewRouter(endpoint)
	if err != nil {
		Log.Print(log.LogLevelError, "Error creating NewRouter: ")
		return err
	}

	return err

}

//close connections
func (c *Communications) Close() {
	Log.Print(log.LogLevelTrace, "Enter to com.Close")
	c.dealer.Destroy()
	c.router.Destroy()
}

//send a message
//return an error if response not equal "ok"
func (c *Communications) Send(msg []byte) error {
	err = c.dealer.SendFrame(msg, goczmq.FlagNone)
	return err
}

// Receve the message.
func (c *Communications) Receve() ([][]byte, error) {
	return c.router.RecvMessage()

}
