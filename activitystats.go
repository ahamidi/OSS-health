package main

import "log"

type CommitStats struct {
	//Day     int
	Week    int
	Month   int
	Quarter int
	Year    int
}

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
