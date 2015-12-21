package cyclic

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"
)

type RunMigrations func() (*sql.DB, error)
type TestWithDB func(*sql.DB, *testing.T)

/*
Given a `RunMigrations` and a test function with test environment
Create database with migrations
Run TestWithDB
Close database connections
*/
func WithDatabase(migrate RunMigrations, t *testing.T, test TestWithDB) {
	db, err := migrate()
	defer db.Close()

	if nil != err {
		t.Error(err)
	}

	test(db, t)
}

/*
Convenience method to run with mysql migrations.
*/
func WithMysqlDatabase(t *testing.T, test TestWithDB) {
	WithDatabase(RunMysqlMigrations, t, test)
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
