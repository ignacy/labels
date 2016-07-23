## Labels

Github doesn't send notifications when labels get changed on a
PR. Sometimes this is a problem. If your workflow depends on it you
can use labels to send periodic notifications to Slack channel with
the information about current PR status.

## Usage

```
> ./labels GITHUB_REPO_OWNER GITHUB_REPO_NAME

# for example for this repo:

> ./labels ignacy lables

# More info:

> ./labels --help
usage: labels [<flags>] <owner> <repo> [<accessToken>]

Flags:
  --help       Show context-sensitive help (also try --help-long and --help-man).
  --skipSlack  With Slack notification

Args:
  <owner>          Id (username) of github repository owner
  <repo>           Id (name) of the repository
  [<accessToken>]  Access token from github

```

**Note** If you don't want to pass accessToken on every call you can
  set it in GITHUB_ACCESS_TOKEN

To make the Slack integration work you will need to setup a webhook on
Slack and then set it in the SLACK_WEBHOOK_URL environment variable.

## Example


![example.png](https://github.com/ignacy/labels/raw/master/resources/example.png)


## License

This code is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
