package request

import (
	"github.com/gettmure/go-tg-bot/internal/pkg/types"
)

// https://core.telegram.org/bots/api#sendaudio
type SendAudio struct {
	ChatId  types.ChatID `json:"chat_id"`
	AudioId string       `json:"audio"`
	Caption string       `json:"caption"`
}

func CreateSendAudio(chatId types.ChatID, audioId, caption string) *SendAudio {
	return &SendAudio{
		ChatId:  chatId,
		AudioId: audioId,
		Caption: caption,
	}
}
