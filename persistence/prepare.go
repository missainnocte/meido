package persistence

import (
	"database/sql"
	"github.com/ivanh/meido/config"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var p *Persistence
var once sync.Once

type Persistence struct {
	db *sql.DB
}

type Executor interface{}

func GetInstance() *Persistence {
	once.Do(Init)
	return p
}

func Init() {
	p = &Persistence{}
	db, err := sql.Open(config.SQL_DRIVER, config.SQL_URL)
	if err != nil {
		panic(err)
	}
	p.db = db
}
func Close() error {
	return p.db.Close()
}

func (ps Persistence) GetDb() *sql.DB {
	return ps.db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
