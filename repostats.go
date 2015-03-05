package main

import (
	"log"
	"os"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
)

type Context struct {
	client *github.Client
}

var context Context

type RepoStats struct {
	Stars        int
	Forks        int
	Contributors int
	Followers    int
}

func getRepoStats(owner, repoName string) *RepoStats {

	// url := strings.Split(repoUrl, "/")

	repo, _, err := context.client.Repositories.Get(owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	c, _, err := context.client.Repositories.ListContributors(owner, repoName, nil)
	if err != nil {
		log.Fatal(err)
	}

	rs := &RepoStats{
		Stars:        *repo.StargazersCount,
		Forks:        *repo.ForksCount,
		Contributors: len(c),
		Followers:    *repo.WatchersCount,
	}

	return rs
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
