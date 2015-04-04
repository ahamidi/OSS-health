package main

import (
	"log"

	"github.com/google/go-github/github"
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

// FIXME can we delete? Output subsumed by GetParticipationStats
func GetCommitStats(owner, repoName string) *CommitStats {

	commitStats := &CommitStats{}

	activity, _, err := context.client.Repositories.ListCommitActivity(owner, repoName)
	if err != nil {
		log.Fatal(err)
	}

	l := len(activity) - 1
	curSumTotal := 0
	for i := range activity {
		curSumTotal += *activity[l-i].Total

		if i == 1 {
			commitStats.Week = curSumTotal
		}

		if i == 4 {
			commitStats.Month = curSumTotal
		}

		if i == 12 {
			commitStats.Quarter = curSumTotal
		}
	}
	commitStats.Year = curSumTotal

	return commitStats
}

func GetParticipationStats(owner, repoName string) *ParticipationStats {
	parStats := &ParticipationStats{
		All:   &CommitStats{},
		Owner: &CommitStats{},
	}

	par, _, err := context.client.Repositories.ListParticipation(owner, repoName)
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
