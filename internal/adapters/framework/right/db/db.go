package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSource string) (*Adapter, error) {
	// connect

	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}

func (da Adapter) AddHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil
}
