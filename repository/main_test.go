package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"os/exec"
	"practice_go/database"
	"testing"
)

var db *sqlx.DB

func setUp() error {
	var err error
	db, err = database.ConnectDB()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = cleanUp()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = setUpTestData()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func setUpTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ../database/init.sql")
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func cleanUp() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ../database/cleanupTestDB.sql")
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func tearDown() {
	err := cleanUp()
	if err != nil {
		log.Println(err.Error())
	}
	_ = db.Close()
}

func TestMain(m *testing.M) {

	err := setUp()
	if err != nil {
		log.Println("couldn't setup test DB")
		os.Exit(1)
	}

	m.Run()

	tearDown()
}
