package db

import (
	"database/sql"
	"encoding/hex"
	"strconv"

	"github.com/DmitryBugrov/log"
	_ "github.com/mattn/go-sqlite3"
)

type BackupDB struct {
	*sql.DB //database
	//	name    string //file name
	//	err  error
}

var (
	err error
	Log *log.Log
)

func (self *BackupDB) Init(_Log *log.Log, db_name string) error {
	Log = _Log
	Log.Print(log.LogLevelTrace, "Enter to db.Init")
	self.DB, err = sql.Open("sqlite3", db_name)
	if err != nil {
		Log.Print(log.LogLevelError, "Error opening DB")
		return err
	}

	err = self.DB.Ping()
	if err != nil {
		Log.Print(log.LogLevelError, "Error opening DB")
		return err
	}
	return err

}

//Close DB
func (self *BackupDB) Close() {
	self.DB.Close()
}

//Create empty DB
func (self *BackupDB) CreateDB() error {
	Log.Print(log.LogLevelTrace, "Enter to db.CreateDB")

	err = self.DB.Ping()
	if err != nil {
		Log.Print(log.LogLevelError, "Error opening DB")
		return err
	}
	self.DB.Exec(`CREATE TABLE file_for_backup (path TEXT,
		filename TEXT, hash TEXT, lastbackup TEXT)`)
	if err != nil {
		Log.Print(log.LogLevelError, "Failed to initialize tables")
		return err
	}
	return nil

}

func (self *BackupDB) AddFile(path string, filename string, hash []byte) error {
	Log.Print(log.LogLevelTrace, "Enter to db.AddFile")
	//args := self.DB.NamedArgs{"$path": path, "$filename": filename, "$change": change}
	ret, err := self.fileIsExist(path, filename)
	if err != nil {
		return err
	}
	if !ret {
		self.DB.Ping()
		Log.Print(log.LogLevelTrace, "\tAdd file: ", filename)
		_, err = self.DB.Exec("INSERT INTO file_for_backup(path,filename,hash) VALUES(?, ?, ?	)", path, filename, hex.EncodeToString(hash))
		if err != nil {
			Log.Print(log.LogLevelError, "Error adding date to DB")
			return err
		}
	}
	return nil
}

func (self *BackupDB) fileIsExist(path string, filename string) (bool, error) {
	Log.Print(log.LogLevelTrace, "Enter to db.fileIsExist")
	rows, err := self.DB.Query("SELECT COUNT (*) FROM file_for_backup WHERE path=? AND filename=?", path, filename)
	defer rows.Close()
	var n int
	rows.Next()
	rows.Scan(&n)
	Log.Print(log.LogLevelTrace, strconv.Itoa(n))
	if err != nil {
		Log.Print(log.LogLevelError, "Error Select to db")
		return false, err
	}
	if n != 0 {
		Log.Print(log.LogLevelTrace, "File exist in DB")
		return true, nil
	}
	Log.Print(log.LogLevelTrace, "File does not exist in DB")
	return false, nil
}

//Get store hash and last backup time file
func (self *BackupDB) GetHashAndBackupTimeFile(path string, filename string) (hash string, lastbackup string, err error) {
	Log.Print(log.LogLevelTrace, "Enter to db.GetHashAndModTimeFile")
	rows, err := self.Query("Select hash,lastbackup FROM file_for_backup WHERE path=? AND filename=?", path, filename)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()
	rows.Next()
	rows.Scan(&hash, &lastbackup)
	return
}
