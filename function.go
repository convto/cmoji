package cmoji

import (
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
	token := os.Getenv("SLACK_OAUTH_TOKEN")

	c := cmd.NewCmd(token, channelID)
	em, err := c.ListEmoji()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	switch {
	case t == "list":
		err := c.SendEmojiMap(em)
		if err != nil {
			fmt.Fprintln(w, err)
		}
	case strings.Contains(t, "stamp"):
		err := c.StampEmoji(strings.Fields(t)[1], em)
		if err != nil {
			fmt.Fprintln(w, err)
		}
	default:
		str := "*get emoji list*\n```cmoji list```\n\n*stamp emoji*\n```cmoji stamp :custom_emoji:```"
		fmt.Fprintln(w, str)
	}
}
