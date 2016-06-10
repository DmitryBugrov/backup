package cfg

import (
	//	"errors"
	//	"log"
	//	"fmt"
	//	"github.com/spf13/viper"
	//"backup/client/log"
	"encoding/json"
	"os"

	"github.com/DmitryBugrov/log"
)

var (
	c   *Cfg
	Log *log.Log
)

type Cfg struct {

	//Path to config file
	config_file_name string

	//BAcfg - config for Backup Agent
	BAcfg struct {
		//Unique client id
		Cid                  string
		MaxFileSegment       int64
		MessageServerAddress string
		MessageServerPort    string
		BackupGroup          []struct {
			Path_for_backup []string
			Schedule        struct {
				Type      byte
				DayOfWeek byte
				Day       int
				Hour      int
				Min       int
				HALB      int //Hour after last backup

			}
		}
	}
}

func (c *Cfg) Init(_Log *log.Log, file_name string) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to cfg.Init")
	c.config_file_name = file_name
	err := c.load()
	return err
}

func (c *Cfg) load() error {
	Log.Print(log.LogLevelTrace, "Enter to cfg.Load")
	file, err := os.Open(c.config_file_name)
	if err != nil {
		Log.Print(log.LogLevelError, "Configuration file cannot be loaded: ", c.config_file_name)
		return err
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c)
	if err != nil {
		Log.Print(log.LogLevelError, "Unable to decode config into struct", err.Error())
	}

	return nil
}
