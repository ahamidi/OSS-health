package main

import "log"

type CommitStats struct {
	Day     int
	Week    int
	Month   int
	Quarter int
	Year    int
}

func GetCommitStats(owner, repoName string) *CommitStats {

	// url := strings.Split(repoUrl, "/")

	repo, _, err := context.client.Repositories.ListCommitActivity(owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(repo[49])
	// c, _, err := context.client.Repositories.ListContributors(owner, repoName, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//has wiki

	//has issues
	//commit activity
	//participation
	// rs := &RepoStats{
	// 	Stars:        *repo.StargazersCount,
	// 	Forks:        *repo.ForksCount,
	// 	Contributors: len(c),
	// 	Followers:    *repo.WatchersCount,
	// 	Wiki:         *repo.HasWiki,
	// 	Issues:       *repo.HasIssues,
	// }
	//
	// return rs
	return nil
}
