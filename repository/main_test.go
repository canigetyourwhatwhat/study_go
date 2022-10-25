package repository

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"practice_go/database"
	"testing"
)

func setUp() error {
	err := database.ConnectDB()
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
		return err
	}
	return nil
}

func cleanUp() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ../database/cleanupTestDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func tearDown() {
	err := cleanUp()
	if err != nil {
		fmt.Println(err.Error())
	}
	_ = database.DB.Close()
}

func TestMain(m *testing.M) {
	var err error

	if err = setUp(); err != nil {
		log.Println("couldn't setup test DB")
		os.Exit(1)
	}

	m.Run()

	tearDown()
}
