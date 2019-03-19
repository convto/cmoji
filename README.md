# cmoji
cmoji is custom error manager in slack.

## how to use
```go
c := cmoji.NewCmd(slackToken, channelID)
em, err := c.ListEmoji()
if err != nil {
	// err handling
}

err := c.StampEmoji(":slack:", em)
if err != nil {
	err handling
}
```