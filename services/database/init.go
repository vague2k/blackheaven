package database

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

func Init() (*Queries, error) {
	db, err := sql.Open("sqlite3", "dummy_data/dummies.db")
	if err != nil {
		return nil, fmt.Errorf("could not init rummage db: \n%s", err)
	}

	// if _, err := db.ExecContext(context.Background(), schema); err != nil {
	// 	return nil, fmt.Errorf("could not create 'inquiries' table in rummage db: \n%s", err)
	// }
	//
	return New(db), nil
}
