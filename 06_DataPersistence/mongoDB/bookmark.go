package main

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Bookmark represents a bookmarks metadata
type Bookmark struct {
	ID                          bson.ObjectId `bson:"_id,omitempty"`
	Name, Description, Location string
	Priority                    int
	CreatedOn                   time.Time
	Tags                        []string
}

// BookmarkStore provides CRUD application agains the collection bookmarks
type BookmarkStore struct {
	C *mgo.Collection
}

// Create inserts the value of Bookmark b into the collection
func (store BookmarkStore) Create(b *Bookmark) error {
	// Assign a new bson.ObjectID
	b.ID = bson.NewObjectId()
	err := store.C.Insert(b)
	return err
}

// Update modifies the value of Bookmark B in the collection
func (store BookmarkStore) Update(b Bookmark) error {
	err := store.C.Update(
		bson.M{"_id": b.ID},
		bson.M{"$set": bson.M{
			"name":        b.Name,
			"description": b.Description,
			"location":    b.Location,
			"priority":    b.Priority,
			"tags":        b.Tags,
		}})
	return err
}

// Delete removes an existing value from the collection
func (store BookmarkStore) Delete(id string) error {
	err := store.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetAll returns all documents from the collection
func (store BookmarkStore) GetAll() []Bookmark {
	var b []Bookmark
	iter := store.C.Find(nil).Sort("priority", "-createdon").Iter()
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

// GetByID returns a single document from the collection
func (store BookmarkStore) GetByID(id string) (Bookmark, error) {
	var b Bookmark
	err := store.C.FindId(bson.ObjectIdHex(id)).One(&b)
	return b, err
}

// GetByTag returns all documents from the collection filtering by tags
func (store BookmarkStore) GetByTag(tags []string) []Bookmark {
	var b []Bookmark
	iter := store.C.Find(bson.M{"tags": bson.M{"$in": tags}}).Sort("priority", "-createdon").Iter()
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}
