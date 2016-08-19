package presenters

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/google/go-github/github"
    "github.com/ignacy/slackpost"
)

const (
    defaultBotName = "github-update"
    defaultEmoji   = ":ghost:"
    defaultColor   = "#e5e5e5"
)

var url = os.Getenv("SLACK_WEBHOOK_URL")

// SendPullRequestDataToSlack sends formated message to slack channel
// with status of all open pull requests
func SendPullRequestDataToSlack(pullRequests []*github.Issue, owner string, repo string) {
    log.Println("URL:>", url)

    m := &slackpost.Attachments{format(pullRequests, owner, repo)}
    body, err := m.Send(url)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("response Body:", body)
}

func format(pullRequests []*github.Issue, owner string, repo string) []*slackpost.Attachment {
    attachments := []*slackpost.Attachment{}

    for _, pr := range pullRequests {
        number := *pr.Number
        newAttachment := slackpost.NewAttachment(
            "",
            "",
            fmt.Sprintf("(%d) %s", number, *pr.Title),
            prLink(owner, repo, number),
            fmt.Sprintf("%s %s", labelsList(pr.Labels), listAssignees(pr.Assignees)),
            labelColorOrDefault(pr.Labels),
            fmt.Sprintf("Opened by %s", *pr.User.Login),
        )
        attachments = append(attachments, newAttachment)
    }
    return attachments
}

func prLink(owner string, repo string, prNumber int) string {
    return fmt.Sprintf("http://github.com/%s/%s/pull/%d", owner, repo, prNumber)
}

func labelsList(listOfLabels []github.Label) string {
    labels := []string{}
    for _, l := range listOfLabels {
        labels = append(labels, *l.Name)
    }

    if len(labels) == 0 {
        return ""
    }
    return fmt.Sprintf(":label: %s", strings.Join(labels, ","))
}

func labelColorOrDefault(listOfLabels []github.Label) string {
    if len(listOfLabels) == 0 {
        return defaultColor
    }
    return *listOfLabels[0].Color
}

func listAssignees(listOfAssignees []*github.User) string {
    users := []string{}
    for _, u := range listOfAssignees {
        users = append(users, *u.Login)
    }

    if len(users) == 0 {
        return ""
    }
    return fmt.Sprintf(":bust_in_silhouette: %s", strings.Join(users, ","))
}
