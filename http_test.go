package main

import (
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	//server = httptest.NewServer(handler) //Creating new server with the user handlers

	//usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint
}

func TestValidPath(t *testing.T) {

	in := "140proof/OSS-Health"
	out := string(`owner:140proof project:OSS-Health grade:%100`)
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/"+in, nil)
	if err != nil {
		t.Fatal(err)
	}

	m := handlers()
	m.ServeHTTP(resp, req)

	if o, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		assert.Equal(t, out, string(o))
	}
}

//func TestInvalidPath(t *testing.T) {
//test := "140proof/OSS-Health"
//resp := httptest.NewRecorder()

//req, err := http.NewRequest("GET", "localhost:8080/"+test, nil)
//if err != nil {
//t.Fatal(err)
//}

////t.Fail()
//}
