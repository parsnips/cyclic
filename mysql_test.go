package cyclic

import (
	"database/sql"
	"testing"
)

func TestRunMysqlMigrations(t *testing.T) {
	_, err := RunMysqlMigrations()

	if nil != err {
		t.Error(err)
	}
}

func TestSimpleMigration(t *testing.T) {
	WithDatabase(RunMysqlMigrations, t, func(db *sql.DB, t *testing.T) {
		_, err := db.Query("SELECT * FROM hello")
		if nil != err {
			t.Error(err)
		}
	})
}
