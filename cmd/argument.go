package cmd

type argument struct {
	Token       string       `json:"token"`
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	ID          string       `json:"user"`
	AsUser      bool         `json:"as_user"`
	User
	Attachments []attachment `json:"attachments"`
}

type User struct {
	UserName    string       `json:"username"`
	IconURL     string       `json:"icon_url"`
}

func newUser(name, url string) User {
	return User{
		UserName: name,
		IconURL:  url,
	}
}

// public chat argument.
// ユーザーのアイコンを偽装するために投稿ユーザーのIDから各種情報を取ってきて使う
// あくまで見た目の差し替えなので、as_userをtrueにしてはいけない
func newUserArgument(token, channel, text string, user User) *argument {
	return &argument{
		Token:   token,
		Channel: channel,
		Text:    text,
		AsUser:  false,
		User:    user,
	}
}

// private chat argument.
func newPrivateArgument(token, channel, text, userID string) *argument {
	return &argument{
		Token:   token,
		Channel: channel,
		ID:      userID,
		Text:    text,
		AsUser:  false,
	}
}

func (arg *argument) setAttachments(attachments ...attachment) {
	arg.Attachments = attachments
}
