package cmd

type argument struct {
	Token       string       `json:"token"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	*User
	AsUser      bool         `json:"as_user"`
	Attachments []attachment `json:"attachments"`
}

type User struct {
	ID string `json:"user"`
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
	user := &User{userID}
	return &argument{
		Token:   token,
		Channel: channel,
		User:    user,
		Text:    text,
		AsUser:  false,
	}
}

func (arg *argument) setAttachments(attachments ...attachment) {
	arg.Attachments = attachments
}
