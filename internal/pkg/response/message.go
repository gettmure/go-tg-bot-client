package response

import (
	"github.com/gettmure/go-tg-bot/internal/pkg/types"
)

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId types.MessageID `json:"message_id"`
	Text      string          `json:"text"`
	Chat      Chat            `json:"chat"`
	Audio     *Audio          `json:"audio"`
	From      *From           `json:"from"`
}

type SendMessageResult Message
