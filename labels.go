package main

import (
	"log"
	"strings"

	"github.com/google/go-github/github"
)
import "golang.org/x/oauth2"

const (
	//	advanonRepoId  = "61127437"
	//	advanonOwnerId = "10156154"
	advanonRepoId  = "Advanon-app"
	advanonOwnerId = "Advanon"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "87970a254801738fedf7e8f03a1fab9a2d515c17"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	prs, _, err := client.Issues.ListByRepo(advanonOwnerId, advanonRepoId, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, pr := range prs {
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
