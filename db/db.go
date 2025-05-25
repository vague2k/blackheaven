package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vague2k/blackheaven/utils"
)

// A database wrapper for the web server that pertains to it's actions
type DB struct {
	Sqlite   *sql.DB // Pointer to the underlying sqlite database
	Dir      string  // The parent directory of the database
	FilePath string  // the database path
}

// Initializes the blackheaven db, returning a pointer to the db instance.
func Init(path string) (*DB, error) {
	if path == "" {
		dataDir := utils.UserDataDir()
		path = dataDir
	}

	var dir string
	var dbFile string
	if path == ":memory:" {
		dbFile = ":memory:"
	} else {
		dir = filepath.Join(path, "blackheaven")
		dbFile = filepath.Join(dir, "blackheaven.db")
		err := os.MkdirAll(dir, 0o777)
		if err != nil {
			return nil, fmt.Errorf("could not create db dir: \n%s", err)
		}
	}

	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("could not init blackheaven db: \n%s", err)
	}

	// _, err = database.Exec(`
	//        CREATE TABLE IF NOT EXISTS items ()
	//    `)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not create 'items' table in blackheaven db: \n%s", err)
	// }

	instance := &DB{
		Sqlite:   database,
		Dir:      dir,
		FilePath: dbFile,
	}

	return instance, nil
}

func (db *DB) Close() {
	db.Sqlite.Close()
}
