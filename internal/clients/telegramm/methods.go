package telegramm

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"tgSubChecker/internal/models"
)

func (c *Client) SendMessage(ctx context.Context, chatId int, text string) error {
	return nil
}

func (c *Client) GetUpdates(ctx context.Context) (*models.Updates, error) {
	params := map[string]string{
		"offset":  strconv.Itoa(c.offset),
		"limit":   "1",
		"timeout": "100",
	}

	// get updates
	var url = c.baseUrl + "getUpdates?" + strings.Trim(strings.Join([]string{"offset=", params["offset"], "&limit=", params["limit"], "&timeout=", params["timeout"]}, ""), " ") + "&allowed_updates=" + c.allowedUpdates
	fmt.Print(url)
	//resp, err := http.Get(url)

	return nil, nil
}
