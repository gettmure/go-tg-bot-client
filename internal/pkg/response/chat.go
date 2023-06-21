package response

import (
	"github.com/gettmure/go-tg-bot/internal/pkg/types"
)

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id types.ChatID `json:"id"`
}
