package main

func getBadge(grade string) string {

	var color string

	switch grade {
	case "A":
		color = "brightgreen"
	case "B":
		color = "green"
	case "C":
		color = "yellowgreen"
	case "D":
		color = "yellow"
	case "E":
		color = "orange"
	case "F":
		color = "red"
	default:
		color = "grey"
	}

	url := "https://img.shields.io/badge/Project_Health-" + grade + "-" + color + ".svg?style=flat&link=http://github.com/140proof.com/oss-health"

	return url
}
