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

	repo, _, err := cfg.client.Repositories.Get(cfg.ctx, owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	c, _, err := cfg.client.Repositories.ListContributors(cfg.ctx, owner, repoName, nil)
	if err != nil {
		log.Fatal(err)
	}

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
