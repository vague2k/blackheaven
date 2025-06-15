package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

func Init() (*Queries, error) {
	// FIXME: delete this later when development is finalized
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("couldnt get working dir", err)
	}
	file := filepath.Join(dir, "/services/database/dummy_data/dummies.db")

	// NOTE: remember to use absolute path for db file
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("could not init blackheaven db: \n%s", err)
	}

	// if _, err := db.ExecContext(context.Background(), schema); err != nil {
	// 	return nil, fmt.Errorf("could not create 'inquiries' table in rummage db: \n%s", err)
	// }
	//
	return New(db), nil
}
