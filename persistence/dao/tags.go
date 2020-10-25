package dao

import "database/sql"

type Tag struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	CreateDate int64  `json:"create_date"`
	UpdateDate int64  `json:"update_date"`
}

/**
获取标签列表
*/
func GetTags() ([]Tag, error) {
	return nil, nil
}

/**
添加标签
*/
func CreateTag() error {
	return nil
}

/**
获取item对应的tags
*/
func GetTagsByItem(db *sql.DB, itemId int64) ([]Tag, error) {
	query := buildGetTagsByItem()
	rows, err := db.Query(query, itemId)
	if err != nil {
		return nil, err
	}
	var ret []Tag
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.Id, &tag.Title); err != nil {
			return nil, err
		}
		ret = append(ret, tag)
	}
	return nil, nil
}

func buildCreateTag() string {
	return ""
}

func buildGetTagsByItem() string {
	return `
SELECT
	tags.*
FROM
	tags
	INNER JOIN res_tags_items ON res_tags_items.tag_id = tags.id 
	AND res_tags_items.item_id = ?;
`
}
