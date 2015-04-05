package main

import (
	"encoding/json"
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

type RepoHealth struct {
	RepoStats   *RepoStats          `json:"repo,omitempty"`
	CommitStats *ParticipationStats `json:"participation,omitempty"`
}

func ossNameHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)

	log.Println(data)

	repoStats := GetRepoStats(data["owner"], data["repo"])

	partStats := GetParticipationStats(data["owner"], data["repo"])

	rh := &RepoHealth{
		RepoStats:   repoStats,
		CommitStats: partStats,
	}

	jsonResponse, err := json.Marshal(rh)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResponse)
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

}
