package presenters

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

const templ = `{{ .Number | printf "%d" }} | {{ .Title }}
Labels: {{ .AllLabels }}
`

var report = template.Must(template.New("issueList").Parse(templ))

type Result struct {
	Number           int
	Title, AllLabels string
}

func PrintPullRequestData(pullRequests []*github.Issue) {
	for _, pr := range pullRequests {
		title := *pr.Title
		number := *pr.Number
		labels := []string{}
		for _, l := range pr.Labels {
			labels = append(labels, *l.Name)
		}
		allLabels := strings.Join(labels, ",")
		result := &Result{number, title, allLabels}
		if err := report.Execute(os.Stdout, result); err != nil {
			log.Fatal(err)
		}
	}
}
