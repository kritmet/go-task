package domain

import "net/http"

// Message message model
type Message struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// NewSuccessMessage new message success
func NewSuccessMessage() Message {
	return Message{
		Code:        http.StatusOK,
		Description: http.StatusText(http.StatusOK),
	}
}
