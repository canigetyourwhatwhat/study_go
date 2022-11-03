package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"practice_go/customErrors"
	"practice_go/database"
	"practice_go/service"
	"testing"
)

var aCon *ArticleController
var ser *service.MyAppService

func TestMain(m *testing.M) {

	// Set up DB
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Errorf("error: %v", customErrors.FailedConnectingDB.Wrap(err, "Failed to connect DB"))
	}

	ser = service.NewMyAppService(db)
	aCon = NewArticleController(ser)

	m.Run()
}
