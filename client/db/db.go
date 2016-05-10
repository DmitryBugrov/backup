package db

import (
	"github.com/DmitryBugrov/log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"encoding/hex"
	"strconv"
)

type BackupDB struct {
	*sql.DB //database
	//	name    string //file name
	//	err  error
}

var (
	err error
)
	
func (self *BackupDB) Init(db_name string) error {
	log.Print(log.LogLevelTrace, "Enter to db.Init")
	self.DB, err = sql.Open("sqlite3", db_name)
	if err != nil {
		log.Print(log.LogLevelError, "Error opening DB")
		return err
	}
	
	err = self.DB.Ping()
	if err != nil {
		log.Print(log.LogLevelError, "Error opening DB")
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
	log.Print(log.LogLevelTrace, "Enter to db.CreateDB")

	err = self.DB.Ping()
	if err != nil {
		log.Print(log.LogLevelError, "Error opening DB")
		return err
	}
	self.DB.Exec(`CREATE TABLE file_for_backup (path TEXT,
		filename TEXT, hash TEXT, lastbackup TEXT)`)
	if err != nil {
		log.Print(log.LogLevelError, "Failed to initialize tables")
		return err
	}
	return nil

}

func (self *BackupDB) AddFile(path string, filename string, hash []byte) error {
	log.Print(log.LogLevelTrace, "Enter to db.AddFile")
	//args := self.DB.NamedArgs{"$path": path, "$filename": filename, "$change": change}
	ret,err:=self.fileIsExist(path,filename)
	if err !=nil {
		return err
	}
	if !ret {
		self.DB.Ping()
		log.Print(log.LogLevelTrace, "\tAdd file: ",filename)
		_, err = self.DB.Exec("INSERT INTO file_for_backup(path,filename,hash) VALUES(?, ?, ?	)", path, filename, hex.EncodeToString(hash))
		if err != nil {
			log.Print(log.LogLevelError, "Error adding date to DB")
			return err
		}
	}
	return nil
}

func (self *BackupDB)fileIsExist(path string,filename string) (bool,error) {
	log.Print(log.LogLevelTrace, "Enter to db.fileIsExist")
	rows,err:=self.DB.Query("SELECT COUNT (*) FROM file_for_backup WHERE path=? AND filename=?", path, filename)
	defer rows.Close()
	var n int
	rows.Next()
	rows.Scan(&n)
	log.Print(log.LogLevelTrace, 	strconv.Itoa(n))
	if err != nil {
	    log.Print(log.LogLevelError, "Error Select to db")
		return false,err
    }	
	if n!=0 {
		log.Print(log.LogLevelTrace,"File exist in DB")
		return true,nil
	}
		log.Print(log.LogLevelTrace,"File does not exist in DB")	
	return false,nil
}

//Get store hash and last backup time file
func (self *BackupDB)GetHashAndBackupTimeFile(path string,filename string) (hash string, lastbackup string, err error) {
	log.Print(log.LogLevelTrace, "Enter to db.GetHashAndModTimeFile")
	rows,err:=self.Query("Select hash,lastbackup FROM file_for_backup WHERE path=? AND filename=?", path, filename)
	if err!=nil {
		return "","",err
	}
	defer rows.Close()
	rows.Next()
	rows.Scan(&hash,&lastbackup)
	return
}