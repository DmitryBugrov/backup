package BA

import (
	"backup/client/cfg"
	"backup/client/mylog"
	"errors"
	"io/ioutil"
	"os"
)

var (
	b *BA
)

type BA struct {
	c        *cfg.Cfg
	fileList []os.FileInfo
}

func (b *BA) Init(config *cfg.Cfg) error {
	mylog.Print(mylog.LogLevelTrace, "Enter to BA.Init")
	b.c = config
	return nil
}

func (b *BA) GetFileList() error {
	mylog.Print(mylog.LogLevelTrace, "Enter to BA.GetFileList")
	err := errors.New("")
	for i := range b.c.BA.Path_for_backup {
		mylog.Print(mylog.LogLevelTrace, "	Get file list: ", b.c.BA.Path_for_backup[i])
		b.fileList, err = ioutil.ReadDir(b.c.BA.Path_for_backup[i])

		//Print file list if ErrorLevel=Trace
		if mylog.LogLevel == 0 {
			for j := range b.fileList {
				mylog.Print(mylog.LogLevelTrace, "		", b.fileList[j].Name())
			}
		}
	}

	return err
}
