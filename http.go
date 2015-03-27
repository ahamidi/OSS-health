package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Context struct {
	client *github.Client
}

var context Context

func ossNameHandler(w http.ResponseWriter, r *http.Request) {
	// req := r.URL.Path[1:]
	data := mux.Vars(r)

	log.Println(data)
	// log.Println(data["owner"], data["repo"])

	repoStats := GetRepoStats(data["owner"], data["repo"])
	//
	log.Printf("%+v\n", repoStats)
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

func configureClient() {
	log.Println("Starting OSS Health App")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("OAUTH_TOKEN")},
	}
	client := github.NewClient(t.Client())

	context.client = client

	// rs := getRepoStats("140proof/OSS-Health")

	// log.Printf("rs: %#v", rs)

	//log.Println("Repo: ", repo)
	//log.Println("Contributors: ", c)

}
