package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var store BookmarkStore
var id string

// init will invoke before the function main.
func init() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{"127.0.0.1"},
		Timeout: 60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[MongoDB Session]: %s\n", err)
	}
	collection := session.DB("bookmarkdb").C("bookmarks")
	store = BookmarkStore{
		C: collection,
	}
}

// createUpdate creates and updates documents in the database
func createUpdate() {
	bookmark := Bookmark{
		Name:        "mgo",
		Description: "Go driver for MongoDB",
		Location:    "https://github.com/go-mgo/mgo",
		Priority:    2,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "mongodb"},
	}
	// Insert a new document
	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}

	id = bookmark.ID.Hex()
	fmt.Printf("New bookmark has been instered with ID: %s\n", id)
	// Update an existing document
	bookmark.Priority = 1
	if err := store.Update(bookmark); err != nil {
		log.Fatalf("[Update]: %s\n", err)
	}
	fmt.Println("The value after update:")
	// Retreive the updated document
	getByID(id)

	// Add another bookmark
	bookmark = Bookmark{
		Name:        "gorethink",
		Description: "Go driver for RethinkDB",
		Location:    "https://github.com/dancannon/gorethink",
		Priority:    3,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "rethinkdb"},
	}
	// Insert the bookmark
	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}
	id = bookmark.ID.Hex()
	fmt.Printf("New bookmark has been inserted with ID: %s\n", id)
}

// getByID gets a document by given ID and print it's details
func getByID(id string) {
	bookmark, err := store.GetByID(id)
	if err != nil {
		log.Fatalf("[GetByID]: %s\n", err)
	}
	fmt.Printf("Name:%s, Description:%s, Priority:%d\n", bookmark.Name, bookmark.Description, bookmark.Priority)
}

// getAll gets all documents and prints their details
func getAll() {
	// Layout for formatting
	layout := "2006-01-02 15:04:05"
	// Retreive all documents
	bookmarks := store.GetAll()
	fmt.Println("Read all documents")
	for _, v := range bookmarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n", v.Name, v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
}

// getByTags gets all documents with matching tags and prints their details
func getByTags() {
	layout := "2006-01-02 15:04:05"
	fmt.Println("Query with Tags - 'go, nosql'")
	bookmarks := store.GetByTag([]string{"go", "nosql"})
	for _, v := range bookmarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n",
			v.Name, v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
	fmt.Println("Query with Tags - 'mongodb'")
	bookmarks = store.GetByTag([]string{"mongodb"})
	for _, v := range bookmarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n",
			v.Name, v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
}

// delete removes an existing document from the collection.
func delete() {
	if err := store.Delete(id); err != nil {
		log.Fatalf("[Delete]: %s\n", err)
	}
	bookmarks := store.GetAll()
	fmt.Printf("Number of documents in the collection after delete:%d\n",
		len(bookmarks))
}

func main() {
	createUpdate()
	getAll()
	getByTags()
	delete()
}
