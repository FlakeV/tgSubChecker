package csvStorage

import (
	"context"
	"tgSubChecker/internal/models"
)

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) GetOwner(ctx context.Context, chatId int) (*models.ChannelOwner, error) {
	return &models.ChannelOwner{
		OwnerID:       415195483,
		Notifications: true,
	}, nil
}
