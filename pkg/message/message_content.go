package message

type MessageContent struct {
	Type MessagesContentType `json:"type"`
	Text *string             `json:"text,omitempty"`
}

func NewTextMessageContent(text string) MessageContent {
	return MessageContent{
		Type: MessagesContentTypeText,
		Text: &text,
	}
}
