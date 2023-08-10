package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/9bany/task/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config := util.LoadConfig()

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database:", err)
	}
	testDb = conn
	testQueries = New(conn)

	os.Exit(m.Run())
}
