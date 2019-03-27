package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	postMessageAPI   = "https://slack.com/api/chat.postMessage"
	postEphemeralAPI = "https://slack.com/api/chat.postEphemeral"
	listEmojiAPI     = "https://slack.com/api/emoji.list"
	listUserAPI      = "https://slack.com/api/users.list"
)

func callChatAPI(token string, arg *argument, url string) error {
	b, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)

	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-type", "application/json; charset=UTF-8")

	c := http.DefaultClient
	res, err := c.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}

func listEmoji(token string) (map[string]string, error) {
	url := listEmojiAPI + "?token=" + token
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
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
	type result struct {
		Ok bool `json:"ok"`
		EmojiMap map[string]string `json:"emoji"`
	}

	var r result
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	fmt.Println(r)

	if !r.Ok {
		return nil, errors.New("failed request")
	}
	return r.EmojiMap, nil
}

type profile struct {
	name    string
	iconURL string
}

func listProfile(token string) (map[string]profile, error) {
	url := listUserAPI + "?token=" + token
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type result struct {
		Ok      bool `json:"ok"`
		Members []struct {
			ID       string `json:"id"`
			Profile  struct {
				DisplayName string `json:"display_name"`
				Image512    string `json:"image_512"`
			} `json:"profile"`
		} `json:"members"`
	}

	var r result
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	fmt.Println(r)

	if !r.Ok {
		return nil, errors.New("failed request")
	}

	profiles := make(map[string]profile)
	for _, m := range r.Members {
		u := profile{
			name:    m.Profile.DisplayName,
			iconURL: m.Profile.Image512,
		}
		profiles[m.ID] = u
	}

	return profiles, nil
}
