package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gettmure/go-tg-bot/internal/pkg/response"
)

type Params struct {
	query *map[string]string
	post  interface{}
}

func fetch[T response.Result](
	c *client,
	token string,
	method Method,
	params *Params,
) (*response.Response[T], error) {
	req, err := buildRequest(token, method, params)
	if err != nil {
		return nil, err
	}

	response, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return decodeBody[T](response)
}

func buildRequest(token string, method Method, params *Params) (*http.Request, error) {
	httpMethod := http.MethodGet
	queryParams := map[string]string{}
	postParams := []byte{}

	if params != nil && params.query != nil {
		queryParams = *params.query
	}

	if params != nil && params.post != nil {
		bytes, err := json.Marshal(params.post)
		if err != nil {
			return nil, err
		}

		httpMethod = http.MethodPost
		postParams = bytes
	}

	url, err := buildMethodUrl(token, method, &queryParams)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(httpMethod, *url, bytes.NewBuffer(postParams))
}

func decodeBody[T response.Result](resp *http.Response) (*response.Response[T], error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch failed with code %d, body: %s", resp.StatusCode, body)
	}

	var data response.Response[T]
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
