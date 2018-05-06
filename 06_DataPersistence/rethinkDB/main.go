package main

import (
	"fmt"
	"log"
	"time"

	r "github.com/dancannon/gorethink"
)

var store BookmarkStore
var id string

// initDB creates a new database and table
func initDB(session *r.Session) {
	// Create database
	_, err := r.DBCreate("bookmarkdb").RunWrite(session)
	if err != nil {
		log.Fatalf("[initDB]: %s\n", err)
	}
	// Create table
	_, err = r.DB("bookmarkdb").TableCreate("bookmarks").RunWrite(session)
	if err != nil {
		log.Fatalf("[initDB]: %s\n", err)
	}
}

// changeFeeds subscribes real-time changes on table bookmarks
func changeFeeds(session *r.Session) {
	// "new_val" is null on a delete change, has a value when insert or update used
	// "old_val" is the opposite
	bookmarks, err := r.Table("bookmarks").Changes().Field("new_val").Run(session)
	if err != nil {
		log.Fatalf("[changeFeeds]: %s\n", err)
	}
	// Launch a goroutine to print real-time updates.
	go func() {
		var bookmark Bookmark
		for bookmarks.Next(&bookmark) {
			if bookmark.ID == "" { // for delete, new_val will be null.
				fmt.Println("Real-time update: Document has been deleted")
			} else {
				fmt.Printf("Real-time update: Name:%s, Description:%s, Priority:%d\n", bookmark.Name, bookmark.Description, bookmark.Priority)
			}
		}
	}()
}

// init connects to rethink and creates the database
func init() {
	// connect to rethink and get a session
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "bookmarkdb",
		MaxIdle:  10,
		MaxOpen:  10,
	})

	if err != nil {
		log.Fatalf("[RethinkDB Session]: %s\n", err)
	}

	// Create Database and Table.
	initDB(session)
	store = BookmarkStore{
		Session: session,
	}
	// Subscribe real-time changes
	changeFeeds(session)
}

// createUpdate creates and updates some example documents
func createUpdate() {
	// Create a new document
	bookmark := &Bookmark{
		Name:        "mgo",
		Description: "Go driver for MongoDB",
		Location:    "https://github.com/go-mgo/mgo",
		Priority:    1,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "mongodb"},
	}
	// Insert a new document.
	if err := store.Create(bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}
	id = bookmark.ID
	fmt.Printf("New bookmark has been inserted with ID: %s\n", id)
	// Update an existing document.
	bookmark.Priority = 2
	if err := store.Update(bookmark); err != nil {
		log.Fatalf("[Update]: %s\n", err)
	}
	fmt.Println("The value after update:")
	// Retrieve the updated document.
	getByID(id)
	// Create a new document
	bookmark = &Bookmark{
		Name:        "gorethink",
		Description: "Go driver for RethinkDB",
		Location:    "https://github.com/dancannon/gorethink",
		Priority:    1,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "rethinkdb"},
	}
	// Insert a new document.
	if err := store.Create(bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}
	id = bookmark.ID
	fmt.Printf("New bookmark has been inserted with ID: %s\n", id)
}

// getByID prints the details of the document with id
func getByID(id string) {
	bookmark, err := store.GetByID(id)
	if err != nil {
		log.Fatalf("[GetByID]: %s\n", err)
	}
	fmt.Printf("Name:%s, Description:%s, Priority:%d\n", bookmark.Name, bookmark.
		Description, bookmark.Priority)
}

// getAll returns and prints the details for all documents
func getAll() {
	// Layout for formatting dates.
	layout := "2006-01-02 15:04:05"
	// Retrieve all documents.
	bookmarks, err := store.GetAll()
	if err != nil {
		log.Fatalf("[GetAll]: %s\n", err)
	}
	fmt.Println("Read all documents")
	for _, v := range bookmarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n", v.Name,
			v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
}

// delete removes an exsiting focument from the table
func delete() {
	if err := store.Delete(id); err != nil {
		log.Fatalf("[Delete]: %s\n", err)
	}
	bookmarks, err := store.GetAll()
	if err != nil {
		log.Fatalf("[GetAll]: %s\n", err)
	}
	fmt.Printf("Number of documents in the table after delete:%d\n", len(bookmarks))
}

func main() {
	createUpdate()
	getAll()
	delete()
}
