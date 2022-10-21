package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func ConnectDB() error {

	if DB != nil {
		return nil
	}
	dbUser := "docker"
	dbPass := "docker"
	dbName := "sampledb"

	// bad example
	connectDbStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)
	db, err := sql.Open("mysql", connectDbStr)

	// good example
	//connectDbStr := mysql.Config{
	//	DBName:    dbName,
	//	User:      dbUser,
	//	Passwd:    dbPass,
	//	Addr:      "localhost:3306",
	//	Net:       "tcp",
	//	ParseTime: true,
	//}
	//db, err := sql.Open("mysql", connectDbStr.FormatDSN())

	//defer db.Close()

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
