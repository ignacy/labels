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
)

var url = os.Getenv("SLACK_WEBHOOK_URL")

// Message - represents message to be send to Slack
type Message struct {
    Text      string `json:"text"`
    Username  string `json:"username"`
    IconEmoji string `json:"icon_emoji"`
}

// NewMessage builds a new message with text and default username and
// emoji
func NewMessage(text string) *Message {
    return &Message{
        text,
        defaultBotName,
        defaultEmoji,
    }
}

// SendPullRequestDataToSlack sends formated message to slack channel
// with status of all open pull requests
func SendPullRequestDataToSlack(pullRequests []*github.Issue, owner string, repo string) {
    log.Println("URL:>", url)

    m := NewMessage(format(pullRequests, owner, repo))

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

func format(pullRequests []*github.Issue, owner string, repo string) string {
    result := "Pull request status \n"
    result += fmt.Sprintf("_There are %d open PR's _\n", len(pullRequests))

    for _, pr := range pullRequests {
        number := *pr.Number
        result += fmt.Sprintf("(%4d) %s :label: %s \n", number, prLink(owner, repo, *pr.Title, number), labelsList(pr.Labels))
    }
    return result
}

func prLink(owner string, repo string, prTitle string, prNumber int) string {
    return fmt.Sprintf("<http://github.com/%s/%s/pull/%d|%s>", owner, repo, prNumber, prTitle)
}

func labelsList(listOfLabels []github.Label) string {
    labels := []string{}
    for _, l := range listOfLabels {
        labels = append(labels, "*"+*l.Name+"*")
    }
    return strings.Join(labels, ",")
}
