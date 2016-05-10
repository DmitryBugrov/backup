package cfg

import (
	//	"errors"
	//	"log"
	//	"fmt"
	//	"github.com/spf13/viper"
	//"backup/client/log"
	"encoding/json"
	"github.com/DmitryBugrov/log"
	"os"
)

var (
	c *Cfg
)

type Cfg struct {
	//	cid	string
	//	paths_list_for_backup [] string
	//	Type map[string]string

	//Path to config file
	config_file_name string

	//BAcfg - config for Backup Agent
	BAcfg struct {
		//Unique client id
		Cid             string
		MaxFileSegment	int64
		BackupGroup []struct {
			Path_for_backup []string
			Schedule struct {
				Type byte
				DayOfWeek byte
				Day int
				Hour int
				Min int
				HALB int //Hour after last backup
				
			}
		}
	}
}

func (c *Cfg) Init(file_name string) error {
	log.Print(log.LogLevelTrace, "Enter to cfg.Init")
	c.config_file_name = file_name
	err := c.load()
	return err
}

func (c *Cfg) load() error {
	log.Print(log.LogLevelTrace, "Enter to cfg.Load")
	file, err := os.Open(c.config_file_name)
	if err != nil {
		log.Print(log.LogLevelError, "Configuration file cannot be loaded: ", c.config_file_name)
		return err
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c)
	if err != nil {
		log.Print(log.LogLevelError, "Unable to decode config into struct", err.Error())
	}

	return nil
}
