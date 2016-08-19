package slackpost

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// Wrapper for a list of attachments
// https://api.slack.com/docs/message-attachmentss
type Attachments struct {
    Attachments []*Attachment `json:"attachments"`
}

// Attachment type. Parameter documentation:
// https://api.slack.com/docs/message-attachments
type Attachment struct {
    Fallback  string `json:"fallback"`
    Pretext   string `json:"pretext"`
    Title     string `json:"title"`
    TitleLink string `json:"title_link"`
    Text      string `json:"text"`
    Color     string `json:"color"`
    Footer    string `json:"footer"`
}

// Builds new attachment with prefiled data
func NewAttachment(fallback, pretext, title, titleLink, text, color, footer string) *Attachment {
    return &Attachment{
        fallback,
        pretext,
        title,
        titleLink,
        text,
        color,
        footer,
    }
}

// This method delivers formatted attachment list to slack webhook
// Example:
//
// attachments := []*Attachments{}
// for _, el := range data {
//     newAttachment := NewAttachment(
//         el.fallback,
//         el.pretext, el.title, el.titleLink,
//         el.text, el.color, el.footer,
//     )
//     attachments = append(attachments, newAttachment)
// }
//
// message := &Attachments{attachments}
// responseBody, err := message.Send("https://slack.webhook.url")
func (attachments *Attachments) Send(webhookURL string) (string, error) {
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(attachments)

    req, err := http.NewRequest("POST", webhookURL, b)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}
