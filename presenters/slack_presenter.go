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

type Message struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

var url = os.Getenv("SLACK_WEBHOOK_URL")

func SendPullRequestDataToSlack(pullRequests []*github.Issue) {
	log.Println("URL:>", url)

	m := Message{format(pullRequests), "github-update", ":ghost:"}

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

func format(pullRequests []*github.Issue) string {
	result := "Pull request status \n"
	result += fmt.Sprintf("_There are %d open PR's _\n", len(pullRequests))

	for _, pr := range pullRequests {
		title := *pr.Title
		number := *pr.Number
		labels := []string{}
		for _, l := range pr.Labels {
			labels = append(labels, "*"+*l.Name+"*")
		}
		all_labels := strings.Join(labels, ",")
		result += fmt.Sprintf("(%4d) <http://github.com/Advanon/Advanon-app/pull/%d|%s>  :label: %s", number, number, title, all_labels)
		result += "\n"
	}
	return result
}
