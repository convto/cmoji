package cmd

type argument struct {
	Token       string       `json:"token"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	User        string       `json:"user"`
	AsUser      bool         `json:"as_user"`
	Attachments []attachment `json:"attachments"`
}

// public chat argument.
func newPublicArgument(token, channel, text string) *argument {
	return &argument{
		Token:   token,
		Channel: channel,
		Text:    text,
		AsUser:  true,
	}
}

// private chat argument.
func newPrivateArgument(token, channel, text, userID string) *argument {
	return &argument{
		Token:   token,
		Channel: channel,
		Text:    text,
		User:    userID,
		AsUser:  false,
	}
}

func (arg *argument) setAttachments(attachments ...attachment) {
	arg.Attachments = attachments
}
