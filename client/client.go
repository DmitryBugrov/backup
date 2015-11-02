package main

import (
	//"errors"
	//	"fmt"
	"backup/client/ba"
	"backup/client/cfg"
	"backup/client/db"
	"backup/client/mylog"
	"os"
)

const config_file_name = "config.json"

var (
	config *cfg.Cfg //config
	ba     *BA.BA   //Backup Agent
	DB     *db.DB
)

func main() {
	//Init lgging
	mylog.Init(os.Stdout, mylog.LogLevelTrace)

	//Init and load config
	config = new(cfg.Cfg)
	err := config.Init(config_file_name)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "No configuration file loaded: ", config_file_name)
		os.Exit(1)
	}

	//Init Backup Agent
	ba = new(BA.BA)
	err = ba.Init(config)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error of initialisation Backup Agent")
		os.Exit(1)
	}

	ba.GetFileList()

	//	mylog.Print(mylog.LogLevelTrace, "config.BA=", config.BA)
	DB = new(db.DB)
	err = DB.Init("ba_client.db")
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error of initialisation DB", err.Error())
		os.Exit(1)
	}
	DB.CreateDB()
	DB.AddFile("./", "test", 0)

}
