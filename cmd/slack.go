package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postMessage(token string, arg *argument) error {
	b, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)

	url := "https://slack.com/api/chat.postMessage"
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-type", "application/json; charset=UTF-8")

	c := http.DefaultClient
	res, err := c.Do(req)
	res.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func listEmoji(token string) (map[string]string, error) {
	url := fmt.Sprintf("https://slack.com/api/emoji.list?token=%s", token)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// emojiJSON is json parse struct.
	// sample response:
	/*
		{
			"ok": true,
			"emoji": {
				"bowtie": "https:\/\/my.slack.com\/emoji\/bowtie\/46ec6f2bb0.png",
				"squirrel": "https:\/\/my.slack.com\/emoji\/squirrel\/f35f40c0e0.png",
				"shipit": "alias:squirrel",
				…
			}
		}
	*/
	type emojiJSON struct {
		OK bool `json:"ok"`
		// errorレスポンスの時はemojiフィールドは存在しない
		// そのため、型はpointerで定義
		EmojiMap map[string]string `json:"emoji"`
	}

	var e emojiJSON
	if err := json.Unmarshal(b, &e); err != nil {
		return nil, err
	}
	fmt.Println(e)

	if !e.OK {
		return nil, errors.New("failed request")
	}
	return e.EmojiMap, nil
}
