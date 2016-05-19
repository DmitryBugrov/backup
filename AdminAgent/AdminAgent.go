package main

import (
	//	"backup/BackupAgent/ba"
	"backup/AdminAgent/cfg"
	"backup/AdminAgent/com"
	//	"backup/BackupAgent/db"
	"os"

	"github.com/DmitryBugrov/log"
)

const config_file_name = "config.json"

var (
	config *cfg.Cfg //config
	Log    *log.Log
	Com    *com.Communications
)

func main() {
	//Init lgging
	Log = new(log.Log)
	Log.Init(log.LogLevelTrace, true, true, true)

	//Init and load config
	config = new(cfg.Cfg)
	err := config.Init(Log, config_file_name)
	if err != nil {
		Log.Print(log.LogLevelError, "No configuration file loaded: ", config_file_name)
		os.Exit(1)
	}

	Com := new(com.Communications)
	err = Com.Init(Log, config)
	if err != nil {
		Log.Print(log.LogLevelError, "Error init communications module: ")
		os.Exit(1)
	}
	defer Com.Close()
}
