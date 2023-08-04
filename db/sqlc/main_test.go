package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://billybills:secret@localhost:5432/billy_bills?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Error conectandose ala base de datos: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())

}
