package db

import (
	"backup/client/mylog"
	"database/sql"
	//"database/sql"
	//"code.google.com/p/go-sqlite/go1/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	//	"github.com/mattn/go-sqlite3"
	//"log"
)

type DB struct {
	*sql.DB //database
	//	name    string //file name
	//	err  error
}

var (
	err error
)

func (_db *DB) Init(db_name string) error {
	mylog.Print(mylog.LogLevelTrace, "Enter to db.Init")
	_db.DB = new(sql.DB)
	//	_db.name = db_name
	_db.DB, err = sql.Open("sqlite3", db_name)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error opening DB")
		return err
	}
	defer _db.Close()

	err = _db.Ping()
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error opening DB")
		return err
	}
	return err

}

//Create empty DB for test
func (_db *DB) CreateDB() {
	mylog.Print(mylog.LogLevelTrace, "Enter to db.CreateDB")
	/*transaction, err := _db.DB.Begin()
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error creating table in DB")
	}*/
	/*err = _db.Ping()
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error opening DB")
		//		return err
	}*/
	_db.DB.Exec(`CREATE TABLE file_for_backup (id integer primary key, path varchar(512),
	 	filename varchar(256), backup INTEGER,
	 	change INTEGER)`)

	if err != nil {
		mylog.Print(mylog.LogLevelError, "Failed to initialize tables")

	}
	//	_db.DB.Close()
	//query.Exec()
	//	return err

}

func (_db *DB) AddFile(path string, filename string, change int) {
	//args := _db.DB.NamedArgs{"$path": path, "$filename": filename, "$change": change}
	_db.DB.Ping()
	_, err = _db.DB.Exec("INSERT INTO file_for_backup(path,filename,change) VALUES(?, ?, ?	)", path, filename, change)
	if err != nil {
		mylog.Print(mylog.LogLevelError, "Error adding date to DB")
		//		return err
	}
	_db.DB.Close()
}
