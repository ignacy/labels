package presenters

import (
	"github.com/google/go-github/github"
	"log"
	"strings"
)

func PrintPullRequestData(pullRequests []*github.Issue) {
	for _, pr := range pullRequests {
		title := *pr.Title
		number := *pr.Number
		labels := []string{}
		for _, l := range pr.Labels {
			labels = append(labels, *l.Name)
		}
		all_labels := strings.Join(labels, ",")
		log.Printf("%4d | %s", number, title)
		log.Printf("Labels: %s", all_labels)
	}
}
