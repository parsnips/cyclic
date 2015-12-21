package cyclic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
)

func RunMysqlMigrations() (*sql.DB, error) {
	name := "cyclic_" + RandomString(8)

	fmt.Println(name)

	control, err := sql.Open("mysql", "cyclic:cyclic@tcp(127.0.0.1:3306)/cyclic")
	defer control.Close()

	if nil != err {
		return nil, err
	}

	stmt, err := control.Prepare("CREATE DATABASE " + name)

	if nil != err {
		return nil, err
	}

	_, err = stmt.Exec()

	if nil != err {
		return nil, err
	}

	errors, ok := migrate.UpSync("mysql://cyclic:cyclic@tcp(127.0.0.1:3306)/"+name, "./sql")

	if !ok {
		return nil, errors[0]
	}

	db, err := sql.Open("mysql", "cyclic:cyclic@tcp(127.0.0.1:3306)/"+name)

	if nil != err {
		return nil, err
	}

	return db, nil
}
