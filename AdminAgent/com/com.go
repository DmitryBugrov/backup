//communications
//This package provides client-server relationship

package com

import (
	"github.com/streadway/amqp"

	"github.com/DmitryBugrov/log"

	"backup/AdminAgent/cfg"
)

var (
	err error
	Log *log.Log
)

type Communications struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	qHello amqp.Queue
}

func (self *Communications) Init(_Log *log.Log, config *cfg.Cfg) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to com.Init")
	url := "amqp://" + config.AAcfg.MessageServerAddress + ":" + config.AAcfg.MessageServerPort
	self.conn, err = amqp.Dial(url)
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}
	self.ch, err = self.conn.Channel()
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}

	self.qHello, err = self.ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	msgs, err := self.ch.Consume(
		self.qHello.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			Log.Print(log.LogLevelTrace, "Received a message: ", string(d.Body))
		}
	}()

	Log.Print(log.LogLevelTrace, " [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil

}

func (self *Communications) Close() error {
	Log.Print(log.LogLevelTrace, "Enter to com.Close")
	if self.ch != nil {
		err := self.ch.Close()
		if err != nil {
			Log.Print(log.LogLevelError, err)
			return err
		}
	}
	err := self.conn.Close()
	if err != nil {
		Log.Print(log.LogLevelError, err)
		return err
	}
	return nil
}

func (self *Communications) SendHello(text string) error {

	err = self.ch.Publish(
		"",               // exchange
		self.qHello.Name, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(text),
		})
	return err
}
