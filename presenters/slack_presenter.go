package presenters

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/google/go-github/github"
)

const (
    defaultBotName = "github-update"
    defaultEmoji   = ":ghost:"
    defaultColor   = "#7CD197"
)

var url = os.Getenv("SLACK_WEBHOOK_URL")

type attachments struct {
    Attachments []*Attachment `json:"attachments"`
}

// Attachment type used to wrap PR data
type Attachment struct {
    Fallback  string `json:"fallback"`
    Pretext   string `json:"pretext"`
    Title     string `json:"title"`
    TitleLink string `json:"title_link"`
    Text      string `json:"text"`
    Color     string `json:"color"`
}

// Builds new attachment with prefiled data
func NewAttachment(title, titleLink, text string) *Attachment {
    fallback := fmt.Sprintf("%s - %s", title, titleLink)
    return &Attachment{
        fallback,
        "",
        title,
        titleLink,
        text,
        defaultColor,
    }
}

// SendPullRequestDataToSlack sends formated message to slack channel
// with status of all open pull requests
func SendPullRequestDataToSlack(pullRequests []*github.Issue, owner string, repo string) {
    log.Println("URL:>", url)

    m := &attachments{format(pullRequests, owner, repo)}

    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(m)

    req, err := http.NewRequest("POST", url, b)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    log.Println("response Body:", string(body))
}

func format(pullRequests []*github.Issue, owner string, repo string) []*Attachment {
    attachments := []*Attachment{}

    for _, pr := range pullRequests {
        number := *pr.Number
        newAttachment := NewAttachment(
            fmt.Sprintf("(%d) %s", number, *pr.Title),
            prLink(owner, repo, number),
            fmt.Sprintf("%s %s", labelsList(pr.Labels), listAssignees(pr.Assignees)),
        )
        attachments = append(attachments, newAttachment)
        //result += fmt.Sprintf("%4d %s %s %s \n", number, prLink(owner, repo, *pr.Title, number), , listAssignees(pr.Assignees))
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
