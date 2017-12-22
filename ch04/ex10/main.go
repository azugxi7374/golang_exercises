package main

import (
	"fmt"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	res, _ := github.SearchIssues(os.Args[1:])

	fmt.Printf("%d issues:\n", res.TotalCount)

	var months []*github.Issue
	var years []*github.Issue
	var others []*github.Issue
	now := time.Now()

	for _, item := range res.Items {
		if item.CreatedAt.After(now.AddDate(0, -1, 0)) {
			months = append(months, item)
		} else if item.CreatedAt.After(now.AddDate(-1, 0, 0)) {
			years = append(years, item)
		} else {
			others = append(others, item)
		}
	}

    fmt.Println("LAST MONTH:")
	showIssues(months)
    fmt.Println()
    fmt.Println("LAST YEAR:")
    showIssues(years)
    fmt.Println()
    fmt.Println("OTHERS:")
    showIssues(others)
}

func showIssues(issues []*github.Issue) {
    for _, item := range issues {
        fmt.Printf("#%-5d %9.9s %.55s\n",
        item.Number, item.User.Login, item.Title)
    }
}
