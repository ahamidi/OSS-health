package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ossNameHandler(w http.ResponseWriter, r *http.Request) {
	// req := r.URL.Path[1:]
	data := mux.Vars(r)
	log.Println(data)
	if data["owner"] == "140proof" {
		fmt.Fprintf(w, "owner:%s project:%s grade:%s", data["owner"], data["repo"], "%100")
	} else {
		fmt.Fprintf(w, "owner:%s project:%s grade:%s", data["owner"], data["repo"], "%20")
	}
}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{owner}/{repo}", ossNameHandler).Methods("GET")
	return r
}
func startServer() {

	log.Println("serving...")
	http.ListenAndServe(":8080", handlers())
}

func main() {
	startServer()
}
