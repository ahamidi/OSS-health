package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func ossNameHandler(w http.ResponseWriter, r *http.Request) {
	// req := r.URL.Path[1:]
	data := mux.Vars(r)

	log.Println(data)
	// log.Println(data["owner"], data["repo"])

	log.Println(context)

	// repoStats := getRepoStats(data["owner"], data["repo"])
	//
	// log.Println(repoStats)
}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{owner}/{repo}", ossNameHandler).Methods("GET")
	return r
}
func startServer() {
	configureClient()
	log.Println("serving...")
	if os.Getenv("PORT") != "" {
		http.ListenAndServe(strings.Join([]string{":", os.Getenv("PORT")}, ""), handlers())
	} else {
		http.ListenAndServe(":8080", handlers())
	}

}

func main() {
	startServer()
}
