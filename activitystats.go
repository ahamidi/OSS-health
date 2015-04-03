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

	weekIdx := 0
	curSumTotal := 0
	for i := len(activity) - 1; i >= 0; i-- {
		weekIdx++
		curSumTotal += *activity[i].Total

		if weekIdx == 1 {
			commitStats.Week = curSumTotal
		}

		if weekIdx == 4 {
			commitStats.Month = curSumTotal
		}

		if weekIdx == 12 {
			commitStats.Quarter = curSumTotal
		}
	}
	commitStats.Year = curSumTotal

	return commitStats
}
