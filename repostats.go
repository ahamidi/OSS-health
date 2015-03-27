package main

import "log"

type RepoStats struct {
	Stars        int  `json:"stars"`
	Forks        int  `json:"forks"`
	Contributors int  `json:"contributors"`
	Followers    int  `json:"followers"`
	Wiki         bool `json:"wiki"`
	Issues       bool `json:"issues"`
}

func GetRepoStats(owner, repoName string) *RepoStats {

	// url := strings.Split(repoUrl, "/")

	repo, _, err := context.client.Repositories.Get(owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	c, _, err := context.client.Repositories.ListContributors(owner, repoName, nil)
	if err != nil {
		log.Fatal(err)
	}
	//has wiki

	//has issues
	//commit activity
	//participation
	rs := &RepoStats{
		Stars:        *repo.StargazersCount,
		Forks:        *repo.ForksCount,
		Contributors: len(c),
		Followers:    *repo.WatchersCount,
		Wiki:         *repo.HasWiki,
		Issues:       *repo.HasIssues,
	}

	return rs
}
