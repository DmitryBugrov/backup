package main

import (
	"backup/BackupAgent/ba"
	"backup/BackupAgent/cfg"
	"backup/BackupAgent/com"
	"backup/BackupAgent/db"
	"bytes"
	"os"

	"github.com/DmitryBugrov/log"
)

const const_config_file_name = "config.json"

var (
	config *cfg.Cfg //config
	ba     *BA.BA   //Backup Agent
	DB     *db.BackupDB
	Log    *log.Log
	Com    *com.Communications
)

func main() {

	//init agent
	if !Init() {
		Log.Print(log.LogLevelError, "Error initialization")
		os.Exit(1)
	}
	defer DB.Close()
	defer Com.Close()

	//start message loop
	for {
		if !ListenerMessages() {
			Log.Print(log.LogLevelTrace, "Receve the message 'Exit' ")
			os.Exit(0)
		}
	}

}

func Init() bool {
	return initWithConf(const_config_file_name)
}

func initWithConf(config_file_name string) bool {
	//Init logging
	Log = new(log.Log)
	Log.Init(log.LogLevelTrace, true, true, true)

	//Init and load config
	config = new(cfg.Cfg)
	err := config.Init(Log, config_file_name)
	if err != nil {
		Log.Print(log.LogLevelError, "No configuration file loaded: ", config_file_name)
		return false
	}

	//Init DB, create, if it need
	DB = new(db.BackupDB)
	if DB.Init(Log, config.BAcfg.LocalDB) != nil {
		Log.Print(log.LogLevelError, "Error of initialization local database file: ", config.BAcfg.LocalDB)
		return false
	}

	if DB.CreateDB() != nil {
		Log.Print(log.LogLevelError, "Error creating local database file: ", config.BAcfg.LocalDB)
		return false
	}

	//Init communications module
	Com = new(com.Communications)
	if Com.Init(Log, config) != nil {
		Log.Print(log.LogLevelError, "Error of initialization communications module")
		return false
	}

	//Init Backup Agent
	ba = new(BA.BA)
	err = ba.Init(Log, config, DB)
	if err != nil {
		Log.Print(log.LogLevelError, "Error of initialization Backup Agent")
		return false
	}
	return true
}

func ListenerMessages() bool {

	reply, err := Com.Receve()
	Log.Print(log.LogLevelTrace, "Received a message")
	if err != nil {
		Log.Print(log.LogLevelError, "Error receve a message")
		return true
	}
	if bytes.Compare(reply[1], []byte("Exit")) == 0 {
		Log.Print(log.LogLevelTrace, "Receve Exit")
		return false
	}

	Log.Print(log.LogLevelTrace, "received unidentified message")

	return true
}
