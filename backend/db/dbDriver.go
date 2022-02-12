package dbDriver

import (
	"database/sql"
	"log"
	"thefreepress/db/gpostgres"
	"thefreepress/tool/setting"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}
var err error

func Setup() *DB {

	// Connect PostgreSql
	dbConn.SQL, err = NewDatabase(gpostgres.Dsn())
	if err != nil {
		panic(err)
	}

	err = testDB(dbConn.SQL)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
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
	db.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenDbConn)
	db.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleDbConn)
	db.SetConnMaxLifetime(time.Duration(setting.DatabaseSetting.MaxDbLifetime) * time.Minute)

	return db, nil
}
