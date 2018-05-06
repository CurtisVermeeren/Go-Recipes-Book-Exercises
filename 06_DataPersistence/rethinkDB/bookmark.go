package main

import (
	"time"

	r "github.com/dancannon/gorethink"
)

// Bookmark represents the metadata of a bookmark
type Bookmark struct {
	ID                          string `gorethink:"id,omitempty" json:"id"`
	Name, Description, Location string
	Priority                    int // (1-5)
	CreatedOn                   time.Time
	Tags                        []string
}

// BookmarkStore provides CRUD operations against the table "bookmarks".
type BookmarkStore struct {
	Session *r.Session
}

// Create inserts bookmark b into the table
func (store BookmarkStore) Create(b *Bookmark) error {
	// Insert on RethinkDB
	resp, err := r.Table("bookmarks").Insert(b).RunWrite(store.Session)
	if err == nil {
		b.ID = resp.GeneratedKeys[0]
	}
	return err
}

// Update modifies an exsiting value of the table
func (store BookmarkStore) Update(b *Bookmark) error {
	var data = map[string]interface{}{
		"name":        b.Name,
		"description": b.Description,
		"location":    b.Location,
		"priority":    b.Priority,
		"tags":        b.Tags,
	}
	// Partial update on RethinkDB
	_, err := r.Table("bookmarks").Get(b.ID).Update(data).RunWrite(store.Session)
	return err
}

// Delete removes a value from the table
func (store BookmarkStore) Delete(id string) error {
	_, err := r.Table("bookmarks").Get(id).Delete().RunWrite(store.Session)
	return err
}

// GetAll returns all documents from the table
func (store BookmarkStore) GetAll() ([]Bookmark, error) {
	bookmarks := []Bookmark{}
	res, err := r.Table("bookmarks").OrderBy("priority", r.Desc("date")).Run(store.Session)
	err = res.All(&bookmarks)
	return bookmarks, err
}

// GetByID returns a singlge element from the table
func (store BookmarkStore) GetByID(id string) (Bookmark, error) {
	var b Bookmark
	res, err := r.Table("bookmarks").Get(id).Run(store.Session)
	res.One(&b)
	return b, err
}
