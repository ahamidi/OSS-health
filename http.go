package main

import (
	"encoding/json"
	"io/ioutil"
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

func getRepoHealth(owner string, repo string) *RepoHealth {
	repoStats := GetRepoStats(owner, repo)

	partStats := GetParticipationStats(owner, repo)

	rh := &RepoHealth{
		RepoStats:   repoStats,
		CommitStats: partStats,
	}

	return rh
}

func ossNameHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)

	log.Println(data)

	rh := getRepoHealth(data["owner"], data["repo"])

	jsonResponse, err := json.Marshal(rh)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func ossBadgeHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)

	log.Println(data)

	rh := getRepoHealth(data["owner"], data["repo"])
	log.Println("RH:", &rh)

	// TODO: Get actual score for Repo
	grade := "A"

	badgeUrl := getBadge(grade)

	// Get badge image
	// TODO:
	// - This is a very stupid hack
	// - We should be generating the badge ourselves, not proxying the badge
	//	 from shields.io
	resp, err := http.Get(badgeUrl)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.Header().Add("Content-Type", "image/svg+xml;charset=utf-8")
	w.WriteHeader(200)
	w.Write(body)

}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{owner}/{repo}", ossNameHandler).Methods("GET")
	r.HandleFunc("/{owner}/{repo}/badge.svg", ossBadgeHandler).Methods("GET")
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
