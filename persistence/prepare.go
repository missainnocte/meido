package persistence

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Prepare() {
	db, err := sql.Open("sqlite3", "../foo.db")
	checkErr(err)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
