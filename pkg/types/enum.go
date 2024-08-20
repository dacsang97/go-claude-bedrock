package types

// enumer -type=MessagesContentType -linecomment -json=true -text=true
type MessagesContentType int

const (
	MessagesContentTypeText MessagesContentType = iota // text
)

// enumer -type=MessagesStopReason -linecomment -json=true -text=true
type MessagesStopReason int

const (
	MessageStopReasonEndTurn       MessagesStopReason = iota //  end_turn
	MessageStopReasonMaxTokens                               //  max_tokens
	MessagesStopReasonStopSequence                           //  stop_sequence
)

// enumer -type=MessagesResponseType -linecomment -json=true -text=true
type MessagesResponseType int

const (
	MessagesResponseTypeMessage MessagesResponseType = iota // message
	MessagesResponseTypeError                               // error
)

const (
	RoleUser      = "user"
	RoleAssistant = "assistant"
)
