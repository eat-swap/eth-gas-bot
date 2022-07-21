package entities

type Update struct {
	UpdateId          int64    `json:"update_id"`
	Message           *Message `json:"message,omitempty"`
	EditedMessage     *Message `json:"edited_message,omitempty"`
	ChannelPost       *Message `json:"channel_post,omitempty"`
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Fancy types Not implemented yet
}

func (u *Update) ExtractMessage() *Message {
	if u.Message != nil {
		return u.Message
	} else if u.EditedMessage != nil {
		return u.EditedMessage
	} else if u.ChannelPost != nil {
		return u.ChannelPost
	}
	return u.EditedChannelPost
}
