package cfg

import (
	//	"errors"
	//	"log"
	//	"fmt"
	//	"github.com/spf13/viper"
	"backup/client/mylog"
	"encoding/json"
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

	//BA - config for Backup Agent
	BA struct {
		//Unic client id
		Cid             string
		Path_for_backup []string
	}
}

func (c *Cfg) Init(file_name string) error {

	mylog.Print(mylog.LogLevelTrace, "Enter to cfg.Init")
	c.config_file_name = file_name
	err := c.load()
	return err
}

func (c *Cfg) load() error {
	mylog.Print(mylog.LogLevelTrace, "Enter to cfg.Load")
	file, err := os.Open(c.config_file_name)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Configuration file cannot be loaded: ", c.config_file_name)
		return err
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Unable to decode config into struct", err.Error())
	}

	return nil
}
