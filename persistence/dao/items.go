package dao

import (
	"database/sql"
	"fmt"
)

type Item struct {
	Id         int64  `json:"id"`
	Path       string `json:"path"`
	Alias      string `json:"alias"`
	Type       int64  `json:"type"`
	UpdateDate int64  `json:"update_date"`
	CreateDate int64  `json:"create_date"`
	Tags       []Tag  `json:"tags"`
}

/**
添加item
*/
func CreateItem(tx *sql.DB, item Item) error {
	query := buildCreateItemSQL(item)
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

/**
获取item列表
*/
func GetItems(db *sql.DB) ([]Item, error) {
	//list := []Item{}
	query := buildGetItemSQL()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Path, &item.Alias, &item.Type, &item.CreateDate, &item.UpdateDate); err != nil {
			return nil, err
		}
		ret = append(ret, item)
	}
	for _, item := range ret {
		tags, err := GetTagsByItem(db, item.Id)
		if err != nil {
			continue
		}
		item.Tags = tags
	}
	return []Item{}, nil
}

/**
为item设置tag
*/
func SetItemTag(id int64) error {
	return nil
}

/**
根据tag获取item
*/
func GetItemsByTags(id []int64) ([]Item, error) {
	return nil, nil
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

func buildGetItemSQL() string {
	return `SELECT items.* FROM items`
}
