package request

import (
	"github.com/gettmure/go-tg-bot/internal/pkg/types"
)

// https://core.telegram.org/bots/api#sendmessage
type SendMessage struct {
	ChatId types.ChatID `json:"chat_id"`
	Text   string       `json:"text"`
}

func CreateSendMessage(chatId types.ChatID, text string) *SendMessage {
	return &SendMessage{
		ChatId: chatId,
		Text:   text,
	}
}
