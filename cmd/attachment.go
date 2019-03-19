package cmd

type attachment struct {
	Text     string `json:"text"`
	ImageUrl string `json:"image_url"`
	Color    string `json:"color"`
	Hidden   bool   `json:"hidden"`
	User     string `json:"user"`
}

func newAttachment(text, img, color string) attachment {
	return attachment{
		Text:     text,
		ImageUrl: img,
		Color:    color,
	}
}
