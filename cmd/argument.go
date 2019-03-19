package cmd

type argument struct {
	Token       string       `json:"token"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	AsUser      bool         `json:"as_user"`
	Attachments []attachment `json:"attachments"`
}

func newArgument(token, channel, text string) *argument {
	return &argument{
		Token:   token,
		Channel: channel,
		Text:    text,
		AsUser:  true,
	}
}

func (arg *argument) setAttachments(attachments ...attachment) {
	arg.Attachments = attachments
}
