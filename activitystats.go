package main

import (
	"log"
	//"github.com/google/go-github/github"
)

// CommitStats is a summary of the last year's worth of commit activity.
type CommitStats struct {
	//Day     int
	Week    int
	Month   int
	Quarter int
	Year    int
}

// ParticipationStats contains a summary of commit activity for the repo
// by everyone and by the owner only.
type ParticipationStats struct {
	All   *CommitStats
	Owner *CommitStats
}

func GetParticipationStats(owner, repoName string) *ParticipationStats {
	parStats := &ParticipationStats{
		All:   &CommitStats{},
		Owner: &CommitStats{},
	}

	par, _, err := cfg.client.Repositories.ListParticipation(cfg.ctx, owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	l := len(par.All) - 1 // to traverse backwards
	allSum := 0           // running commit count for everyone
	ownerSum := 0         // running commit count for owner
	for i := range par.All {
		allSum += par.All[l-i]
		ownerSum += par.Owner[l-i]

		switch i {
		case 1:
			parStats.All.Week = allSum
			parStats.Owner.Week = ownerSum
		case 4:
			parStats.All.Month = allSum
			parStats.Owner.Month = ownerSum
		case 12:
			parStats.All.Quarter = allSum
			parStats.Owner.Quarter = ownerSum
		}
	}

	parStats.All.Year = allSum
	parStats.Owner.Year = ownerSum

	return parStats
}
