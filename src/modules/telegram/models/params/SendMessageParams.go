package params

type SendMessageParams struct {
	ChatId                   int64            `json:"chat_id"`
	Text                     string           `json:"text"`
	ParseMode                MessageParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool             `json:"disable_notification,omitempty"`
	ProtectedContent         bool             `json:"protected_content,omitempty"`
	ReplyToMessageId         int64            `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"`

	// Markup is not implemented yet
}

type MessageParseMode string

const (
	MessageParseModeMarkdown MessageParseMode = "MarkdownV2"
	MessageParseModeHTML     MessageParseMode = "HTML"
)
