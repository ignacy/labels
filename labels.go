package main

import (
	"log"

	"github.com/google/go-github/github"
	"github.com/ignacy/labels/presenters"
	"golang.org/x/oauth2"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	owner       = kingpin.Arg("owner", "Id (username) of github repository owner").Required().String()
	repo        = kingpin.Arg("repo", "Id (name) of the repository").Required().String()
	skipSlack   = kingpin.Flag("skipSlack", "With Slack notification").Bool()
	accessToken = kingpin.Arg("accessToken", "Access token from github").Envar("GITHUB_ACCESS_TOKEN").String()
)

func main() {
	kingpin.Parse()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *accessToken})
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	pullRequests, _, err := client.Issues.ListByRepo(*owner, *repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	presenters.PrintPullRequestData(pullRequests)
	if !*skipSlack {
		presenters.SendPullRequestDataToSlack(pullRequests, *owner, *repo)
	}
}
