package message

type MessagesRequest struct {
	AnthropicVersion string    `json:"anthropic_version"`
	System           string    `json:"system"`
	Messages         []Message `json:"messages"`
	MaxTokens        int       `json:"max_tokens"`
	Temperature      *float32  `json:"temperature,omitempty"`
	TopP             *float32  `json:"top_p,omitempty"`
	TopK             *int      `json:"top_k,omitempty"`
}

type MessagesResponse struct {
	ID         string               `json:"id"`
	Type       MessagesResponseType `json:"type"`
	Role       string               `json:"role"`
	Content    []MessageContent     `json:"content"`
	StopReason MessagesStopReason   `json:"stop_reason"`
}

type Message struct {
	Role    string           `json:"role"`
	Content []MessageContent `json:"content"`
}

func NewUserTextMessage(text string) Message {
	return Message{
		Role:    RoleUser,
		Content: []MessageContent{NewTextMessageContent(text)},
	}
}
