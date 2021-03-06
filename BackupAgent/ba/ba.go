package BA

import (
	"backup/BackupAgent/cfg"
	//	"backup/client/nc"
	"backup/BackupAgent/db"
	"crypto/md5"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/DmitryBugrov/log"

	"encoding/hex"
)

var (
	self *BA
	err  error
	Log  *log.Log
)

type BA struct {
	*cfg.Cfg
	*db.BackupDB
	//fileList []os.FileInfo
	fileListForBackup []fileWithParam
}

type fileWithParam struct {
	path     string
	filename string
	modTime  time.Time
	hash     []byte
}

type block struct {
	source    string   `json:"source"`
	path      string   `json:"path"`
	filename  string   `json:"filename"`
	num       int64    `json:"num"`
	timestamp string   `json:"timestamp"`
	data      []byte   `json:"data"`
	size      int64    `json:"size"`
	hash      [16]byte `json:"hash"`
}

func (self *BA) Init(_Log *log.Log, config *cfg.Cfg, db *db.BackupDB) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to BA.Init")
	self.Cfg = config
	self.BackupDB = db
	return nil
}

//get file list for backup with param
func (self *BA) getFileList() error {
	Log.Print(log.LogLevelTrace, "Enter to BA.GetFileList")
	//	err := errors.New("")
	for i := range self.BAcfg.BackupGroup[0].Path_for_backup {
		path := self.BAcfg.BackupGroup[0].Path_for_backup[i]
		Log.Print(log.LogLevelTrace, "	Get file list: ", path)
		fileList, _ := ioutil.ReadDir(path)

		for j := range fileList {
			Log.Print(log.LogLevelTrace, "		", fileList[j].Name())
			filename := fileList[j].Name()
			hash, err := getHashFile(path, filename)
			if err != nil {
				Log.Print(log.LogLevelError, "Error calculate hash", path+"/"+filename)
				//return err

			} else {
				Log.Print(log.LogLevelTrace, "\t hash: ", hex.EncodeToString(hash))
				newelement := fileWithParam{path, filename, fileList[j].ModTime(), hash}
				self.fileListForBackup = append(self.fileListForBackup, newelement)
			}

		}

	}

	return nil
}

func (self *BA) UpdateFileListInDB() {
	Log.Print(log.LogLevelTrace, "Enter to BA.UpdateFileListInDB")
	self.getFileList()
	for i := range self.fileListForBackup {
		path := self.fileListForBackup[i].path
		filename := self.fileListForBackup[i].filename
		//	hash,_:=self.getHashFile(path,filename)
		hash := self.fileListForBackup[i].hash
		self.BackupDB.AddFile(path, filename, hash)
	}
}

func getHashFile(path string, filename string) ([]byte, error) {
	Log.Print(log.LogLevelTrace, "Enter to BA.getHashFile")
	var result []byte
	file, err := os.Open(path + "/" + filename)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil

}

//select modified files in array fileListForBackup
func (self *BA) findModFile() error {
	Log.Print(log.LogLevelTrace, "Enter to BA.findModFile")
	max := len(self.fileListForBackup)
	for i := 0; i < max; i++ {
		path := self.fileListForBackup[i].path
		filename := self.fileListForBackup[i].filename
		hash_now := self.fileListForBackup[i].hash
		modtime := self.fileListForBackup[i].modTime
		hashDB, lastbackup, _ := self.BackupDB.GetHashAndBackupTimeFile(path, filename)
		if hex.EncodeToString(hash_now) == hashDB {
			Log.Print(log.LogLevelTrace, "\t file:", filename, " hashes are the same")
			//	Log.Print(log.LogLevelTrace,self.fileListForBackup)
			self.fileListForBackup = append(self.fileListForBackup[:i], self.fileListForBackup[i+1:]...)
			max--
			i--
			//	Log.Print(log.LogLevelTrace,self.fileListForBackup)
		}
		Log.Print(log.LogLevelTrace, "\t file hash ", filename, " :", hex.EncodeToString(hash_now))
		Log.Print(log.LogLevelTrace, "\t file hash in DB ", filename, " :", hashDB)
		Log.Print(log.LogLevelTrace, "\t file mod time:", filename, modtime.String())
		Log.Print(log.LogLevelTrace, "\t file lastbackup:", filename, lastbackup)
	}
	return nil
}

//backup changed files
func (self *BA) backup() error {
	Log.Print(log.LogLevelTrace, "Enter to BA.backup")
	for i := range self.fileListForBackup {
		self.backupFile(self.fileListForBackup[i].path, self.fileListForBackup[i].filename)
	}

	return nil
}

//backup file
func (self *BA) backupFile(path string, filename string) error {
	Log.Print(log.LogLevelTrace, "Enter to BA.backupFile")
	file, err := os.Open(path + "/" + filename)
	if err != nil {
		Log.Print(log.LogLevelWarning, "Error reading file: ", path, "/", filename)
		return err
	}
	var file_block block
	file_block.data = make([]byte, self.BAcfg.MaxFileSegment)
	var i int64
	fi, _ := file.Stat()
	fileSize := fi.Size()
	//Log.Print(log.LogLevelTrace, "________",self.BAcfg.MaxFileSegment)
	for i = 0; i < fileSize/self.BAcfg.MaxFileSegment+1; i++ {
		n, err := file.ReadAt(file_block.data, i*self.BAcfg.MaxFileSegment)
		if (err != nil) && (err != io.EOF) {
			Log.Print(log.LogLevelWarning, "Error reading file: ", path, "/", filename)
			return err
		}
		file_block.hash = md5.Sum(file_block.data)
		file_block.size = int64(n)
		file_block.source = self.BAcfg.Cid
		file_block.path = path
		file_block.filename = filename
		file_block.num = i
		file_block.timestamp = time.Now().Format(time.StampMilli)
		SendFileBlockToServer(file_block)
		Log.Print(log.LogLevelTrace, "__", file_block.filename, " ", i, " ", file_block.size)
	}
	return nil
}

func (self *BA) StartBackup() error {
	Log.Print(log.LogLevelTrace, "Enter to BA.StartBackup")
	err = self.getFileList()
	if err != nil {
		return err
	}
	//self.UpdateFileListInDB()
	err = self.findModFile()
	if err != nil {
		return err
	}
	err = self.backup()
	if err != nil {
		return err
	}
	return nil
}

func SendFileBlockToServer(file_block block) error {
	return nil
}
