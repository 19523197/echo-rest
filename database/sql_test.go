package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSetupDatabaseSQL(t *testing.T) {
	_, err := SetupSQLDatabase()

	if err != nil {
		t.Error(err.Error())
	}
	t.Log(err)
}
