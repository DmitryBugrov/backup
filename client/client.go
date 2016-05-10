package main

import (
	//"errors"
	//	"fmt"
	"os"
	"backup/client/ba"
	"backup/client/cfg"
	"backup/client/db"
//	"backup/client/nc"
	//	"backup/client/log"
	"github.com/DmitryBugrov/log"
	
)

const config_file_name = "config.json"

var (
	config *cfg.Cfg //config
	ba     *BA.BA   //Backup Agent
	DB     *db.BackupDB
)

func main() {
	//Init lgging
	log.Init(log.LogLevelTrace,true,true,true)

	//Init and load config
	config = new(cfg.Cfg)
	err := config.Init(config_file_name)
	if err != nil {
		log.Print(log.LogLevelError, "No configuration file loaded: ", config_file_name)
		os.Exit(1)
	}

	//Init DB, create, if it need
	DB = new(db.BackupDB)
	//	err = DB.Init("ba_client.db")
	if DB.Init("ba_client.db") != nil {
		os.Exit(1)
	}
	defer DB.Close()
	if DB.CreateDB() != nil {
		os.Exit(1)
	}

	//Init Backup Agent
	ba = new(BA.BA)
	err = ba.Init(config, DB)
	if err != nil {
		log.Print(log.LogLevelError, "Error of initialisation Backup Agent")
		os.Exit(1)
	}

	ba.StartBackup()
}
