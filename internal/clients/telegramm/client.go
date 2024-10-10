package telegramm

import (
	"context"
	"encoding/json"
)

const (
	API_URL = "https://api.telegram.org/bot"
)

type Client struct {
	ctx            context.Context
	baseUrl        string
	offset         int
	allowedUpdates string
}

func NewClient(ctx context.Context, botToken string, offset int, allowedUpdates []string) (*Client, error) {
	allowedUpdatesJson, err := json.Marshal(allowedUpdates)
	if err != nil {
		return nil, err
	}
	allowedUpdatesStr := string(allowedUpdatesJson)

	return &Client{
		ctx:            ctx,
		baseUrl:        API_URL + botToken + "/",
		offset:         offset,
		allowedUpdates: allowedUpdatesStr,
	}, nil
}
