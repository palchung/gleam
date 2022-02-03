package dbDriver

import (
	"database/sql"
	"thefreepress/tool/setting"
	"time"
	"log"
	"thefreepress/db/gpostgres"
	
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func Setup() *DB{

	// Connect PostgreSql
	d, err := NewDatabase(gpostgres.Dsn())
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenDbConn)
	d.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleDbConn)
	d.SetConnMaxLifetime(time.Duration(setting.DatabaseSetting.MaxDbLifetime) * time.Minute)

	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	return dbConn
}

// test database connection
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// create a new database for the app
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}