package database

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var DB *sqlx.DB

func ConnectDB() error {

	if DB != nil {
		return nil
	}
	dbUser := os.Getenv("USER_NAME")
	dbPass := os.Getenv("USER_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectDbStr := mysql.Config{
		DBName:               dbName,
		User:                 dbUser,
		Passwd:               dbPass,
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Open("mysql", connectDbStr.FormatDSN())

	if err != nil {
		fmt.Println("Couldn't connect the Database")
		return err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Ping to DB is failed")
		return err
	} else {
		log.Println("DB is healthy")
	}

	DB = db
	return nil
}
