package client

import (
	"fmt"
	"net/url"
)

type Method string

// https://core.telegram.org/bots/api#available-methods
const (
	GetMe       Method = "getMe"
	GetUpdates  Method = "getUpdates"
	SendMessage Method = "sendMessage"
	SendAudio   Method = "sendAudio"
)

func buildMethodUrl(token string, method Method, params *map[string]string) (*string, error) {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if params == nil {
		result := url.String()

		return &result, nil
	}

	query := url.Query()
	for key, value := range *params {
		query.Set(key, value)
	}

	url.RawQuery = query.Encode()
	result := url.String()

	return &result, nil
}
