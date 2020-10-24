package dao

import (
	"database/sql"
	"fmt"
)

type Item struct {
	Id         int64
	Path       string
	Alias      string
	Type       int
	UpdateDate int64
	CreateDate int64
}

func CreateItem(tx sql.Tx, item Item) error {
	query := buildCreateItemSQL(item)
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func GetItems() error {
	return nil
}

func buildCreateItemSQL(item Item) string {
	sqlFmt := `
INSERT INTO items ( path, alias, type, create_date, update_date )
VALUES
	( '%v', '%v', %v, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP );
`
	sql := fmt.Sprintf(sqlFmt, item.Path, item.Alias, item.Type)
	return sql
}

func GetItemSQL() string {
	return `SELECT items.* FROM items`
}
