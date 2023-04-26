package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupSQLDatabase() (*sql.DB, error) {

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)
	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetConnMaxLifetime(time.Minute * 60)

	return db, err
}
