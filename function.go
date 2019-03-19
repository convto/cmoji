package cmoji

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/srttk/cmoji/cmd"
)

func Cmoji(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	query, _ := url.ParseQuery(string(b))
	t := query.Get("text")
	channelID := query.Get("channel_id")
	userID := query.Get("user_id")
	token := os.Getenv("SLACK_OAUTH_TOKEN")

	c := cmd.NewCmd(token, channelID, userID)
	em, err := c.ListEmoji()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	switch {
	case t == "list":
		err := c.SendEmojiMap(em)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	case strings.Contains(t, "stamp"):
		fields := strings.Fields(t)
		if len(fields) != 2 {
			err := errors.New("required :custom_emoji: parameter. and support only one emoji")
			fmt.Fprintln(w, err)
			return
		}
		err := c.StampEmoji(strings.Fields(t)[1], em)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	default:
		err := c.HelpMessage()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	}
}
