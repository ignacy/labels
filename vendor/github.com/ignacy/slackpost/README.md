# Slackpost

Post to Slack webhook using Golang

## Description

Right now slackpost makes it easy to post **Attachments** to Slack's webhooks.
More info on how attachments look and work can be found here: [Attaching content and links to messages](https://api.slack.com/docs/message-attachments)

## Usage

```go
attachments := []*Attachments{}
for _, el := range data {
    newAttachment := NewAttachment(
        el.fallback,
        el.pretext,
        el.title,
        el.titleLink,
        el.text,
        el.color,
        el.footer,
    )
    attachments = append(attachments, newAttachment)
}

message := &Attachments{attachments}
responseBody, err := message.Send("https://slack.webhook.url")
```

## License

This code is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).