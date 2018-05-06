package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetUsers test HTTP Get to "/users" using ResponseRecorder
func TestGetUsers(t *testing.T) {
	// Setup mux for testing
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	// create a GET request at /users
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	//  A ResponseRecorder object is created using httptest.NewRecorder to record the returned HTTP responses for later inspection in tests.
	w := httptest.NewRecorder()
	// Method ServeHTTP of the HTTP handler is called by providing ResponseRecorder and Request objects to invoke the HTTP Get request on "/users"
	r.ServeHTTP(w, req)
	// Check for status code 200
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

// TestGetUsersWithServer test HTTP Get to "/users" using Server
func TestGetUsersWithServer(t *testing.T) {
	// Setup mux for testing
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	// Create a test server
	server := httptest.NewServer(r)
	defer server.Close()
	//  Create the url and GET request
	usersURL := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersURL, nil)
	// Use the DefaultClient to make the GET request on the test server
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	// Check that code 200 is returned
	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}
