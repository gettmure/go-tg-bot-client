package client

import (
	"fmt"
	"net/http"

	"github.com/gettmure/go-tg-bot/internal/pkg/request"
	"github.com/gettmure/go-tg-bot/internal/pkg/response"
	"github.com/gettmure/go-tg-bot/internal/pkg/types"
)

type Client interface {
	// https://core.telegram.org/bots/api#getme
	GetMe(token string) (*response.GetMeResult, error)

	// https://core.telegram.org/bots/api#getupdates
	GetUpdates(token string, offset int64) (*response.GetUpdatesResult, error)

	// https://core.telegram.org/bots/api#sendmessage
	SendMessage(token string, chatId types.ChatID, text string) (*response.SendMessageResult, error)

	// https://core.telegram.org/bots/api#sendaudio
	SendAudio(token string, chatId types.ChatID, audioId string, caption string) (*response.SendAudioResult, error)
}

type client struct {
	http http.Client
}

func InitClient() Client {
	return &client{http: http.Client{}}
}

func (c *client) GetMe(token string) (*response.GetMeResult, error) {
	resp, err := fetch[response.GetMeResult](c, token, GetMe, nil)
	if err != nil {
		return nil, err
	}

	return validateResponse[response.GetMeResult](resp)
}

func (c *client) GetUpdates(token string, offset int64) (*response.GetUpdatesResult, error) {
	resp, err := fetch[response.GetUpdatesResult](c, token, GetMe, &Params{
		query: &map[string]string{"offset": fmt.Sprint(offset)},
	})
	if err != nil {
		return nil, err
	}

	return validateResponse[response.GetUpdatesResult](resp)
}

func (c *client) SendMessage(token string, chatId types.ChatID, text string) (*response.SendMessageResult, error) {
	resp, err := fetch[response.SendMessageResult](c, token, GetMe, &Params{
		post: request.CreateSendMessage(chatId, text),
	})
	if err != nil {
		return nil, err
	}

	return validateResponse[response.SendMessageResult](resp)
}

func (c *client) SendAudio(token string, chatId types.ChatID, audioId string, caption string) (*response.SendAudioResult, error) {
	resp, err := fetch[response.SendAudioResult](c, token, GetMe, &Params{
		post: request.CreateSendAudio(chatId, audioId, caption),
	})
	if err != nil {
		return nil, err
	}

	return validateResponse[response.SendAudioResult](resp)
}

func validateResponse[T response.Result](resp *response.Response[T]) (*T, error) {
	if resp.Ok != true {
		return nil, fmt.Errorf("telegram response is not OK\ndescription:%s", *resp.Description)
	}

	return resp.Result, nil
}
