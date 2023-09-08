package models

import (
	"fmt"
)

type Bookmark struct {
	ID          int `json:"id"`
	link_text   string
	href        string
	description string
	created_at  string
	updated_at  string
}

func allBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark

	rows, err := db.Query("select * from bookmark")

	if err != nil {
		return nil, fmt.Errorf("allBookmarks: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var b Bookmark

		if err := rows.Scan(&b.ID, &b.link_text, &b.href, &b.description, &b.created_at, &b.updated_at); err != nil {
			return nil, fmt.Errorf("allBookmarks: %v", err)
		}
		bookmarks = append(bookmarks, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allBookmarks: %v", err)
	}

	return bookmarks, nil
}
