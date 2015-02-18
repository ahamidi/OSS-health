package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	server = httptest.NewServer(handler) //Creating new server with the user handlers

	usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint
}

func TestValidPath(t *testing.T) {
	test := "140proof/OSS-Health"
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "localhost:8080/"+test, nil)
	if err != nil {
		t.Fatal(err)
	}
	//t.Fail()
}

func TestInvalidPath(t *testing.T) {
	test := "140proof/OSS-Health"
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "localhost:8080/"+test, nil)
	if err != nil {
		t.Fatal(err)
	}

	//t.Fail()
}
