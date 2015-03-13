package main

import (
	"log"

	oh "github.com/140proof/oss-health"
)

func main() {

	stats := oh.GetRepoStats("rails", "rails")

	log.Println("Stats:", stats)
}
