package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Cmd struct {
	token     string
	channelID string
	userID    string
}

func NewCmd(token, channelID, userID string) Cmd {
	return Cmd{
		token:     token,
		channelID: channelID,
		userID:    userID,
	}
}

func (c Cmd) ListEmoji() (map[string]string, error) {
	return listEmoji(c.token)
}

func (c Cmd) StampEmoji(emoji string, emojiMap map[string]string) error {
	text := fmt.Sprintf("stamp `%s`", emoji)
	imgURL := emojiMap[strings.Trim(emoji, ":")]
	a := newAttachment(text, imgURL, "#FFAACC")
	arg := newPublicArgument(c.token, c.channelID, "")
	arg.setAttachments(a)

	return callChatAPI(c.token, arg, postMessageAPI)
}

func (c Cmd) SendEmojiMap(emojiMap map[string]string) error {
	var keys []string
	for k := range emojiMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	buf := new(bytes.Buffer)
	for _, v := range keys {
		fmt.Fprintf(buf, "%s - :%s: | ", v, v)
	}
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	arg := newPrivateArgument(c.token, c.channelID, string(b), c.userID)

	return callChatAPI(c.token, arg, postEphemeralAPI)
}

func (c Cmd) HelpMessage() error {
	text := "cmoji is custom emoji manager, useage.\n\n*get emoji list*\n```cmoji list```\n\n*stamp emoji*\n```cmoji stamp :custom_emoji:```"
	arg := newPrivateArgument(c.token, c.channelID, text, c.userID)

	return callChatAPI(c.token, arg, postEphemeralAPI)
}
