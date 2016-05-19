package cfg

import (
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

	//AAcfg - config for Admin Agent
	AAcfg struct {
		//Unique client id
		Aid                  string
		MaxFileSegment       int64
		MessageServerAddress string
		MessageServerPort    string
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
		return err
	}

	return nil
}
