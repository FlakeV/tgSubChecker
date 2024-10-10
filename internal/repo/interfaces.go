package repo

import (
	"context"

	"tgSubChecker/internal/models"
)

var EventType = map[string]string{
	"left":   "unsubscribed",
	"member": "subscribed",
}

type Saver interface {
	NewSub(ctx context.Context, update *models.Update) error
}

type Reader interface {
	GetOwner(ctx context.Context, chatId int) (*models.ChannelOwner, error)
}
