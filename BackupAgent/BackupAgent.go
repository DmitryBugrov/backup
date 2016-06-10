package main

import (
	//"errors"
	//	"fmt"
	"backup/BackupAgent/ba"
	"backup/BackupAgent/cfg"
	"backup/BackupAgent/com"
	"backup/BackupAgent/db"
	"os"
	//	"backup/client/nc"
	//	"backup/client/log"
	"github.com/DmitryBugrov/log"
)

const config_file_name = "config.json"

var (
	config *cfg.Cfg //config
	ba     *BA.BA   //Backup Agent
	DB     *db.BackupDB
	Log    *log.Log
	Com    *com.Communications
)

func main() {
	//Init logging
	Log = new(log.Log)
	Log.Init(log.LogLevelTrace, true, true, true)

	//Init and load config
	config = new(cfg.Cfg)
	err := config.Init(Log, config_file_name)
	if err != nil {
		Log.Print(log.LogLevelError, "No configuration file loaded: ", config_file_name)
		os.Exit(1)
	}

	//Init DB, create, if it need
	DB = new(db.BackupDB)
	//	err = DB.Init("ba_client.db")
	if DB.Init(Log, "ba_client.db") != nil {
		os.Exit(1)
	}
	defer DB.Close()
	if DB.CreateDB() != nil {
		os.Exit(1)
	}

	//Init communications module
	Com := new(com.Communications)
	if Com.Init(Log, config) != nil {
		os.Exit(1)
	}
	defer Com.Close()

	Com.SendHello("Hello from BackupAgent")

	//Init Backup Agent
	ba = new(BA.BA)
	err = ba.Init(Log, config, DB)
	if err != nil {
		Log.Print(log.LogLevelError, "Error of initialisation Backup Agent")
		os.Exit(1)
	}

	//	ba.StartBackup()
}
