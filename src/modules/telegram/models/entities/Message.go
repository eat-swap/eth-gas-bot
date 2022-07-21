package entities

type Message struct {
	MessageId            int64    `json:"message_id"`
	From                 *User    `json:"from,omitempty"`
	Chat                 *Chat    `json:"chat,omitempty"`
	Date                 int64    `json:"date"`
	SenderChat           *Chat    `json:"sender_chat,omitempty"`
	ForwardFrom          *User    `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat    `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId int64    `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string   `json:"forward_signature,omitempty"`
	ForwardSenderName    string   `json:"forward_sender_name,omitempty"`
	ForwardDate          int64    `json:"forward_date,omitempty"`
	IsAutomaticForward   bool     `json:"is_automatic_forward,omitempty"`
	ReplyToMessage       *Message `json:"reply_to_message,omitempty"`
	EditDate             int64    `json:"edit_date,omitempty"`
	Text                 string   `json:"text,omitempty"`

	// Message types other than text not implemented yet
}
