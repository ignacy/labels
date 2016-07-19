package main

import (
	"log"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	ownerId     = kingpin.Arg("ownerId", "Id (username) of github repository owner").Required().String()
	repoId      = kingpin.Arg("repoId", "Id (name) of the repository").Required().String()
	accessToken = kingpin.Arg("accessToken", "Access token from github").Envar("GITHUB_ACCESS_TOKEN").String()
)

func main() {
	kingpin.Parse()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *accessToken})
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	prs, _, err := client.Issues.ListByRepo(*ownerId, *repoId, nil)
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
