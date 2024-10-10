package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"tgSubChecker/internal/models"
	"tgSubChecker/internal/repo"
)

type Reader struct {
	pool *pgxpool.Pool
}

func NewReader(pool *pgxpool.Pool) repo.Reader {
	return &Reader{pool: pool}
}

func (r *Reader) GetOwner(ctx context.Context, chatId int) (*models.ChannelOwner, error) {
	var owner models.ChannelOwner
	err := r.pool.QueryRow(
		ctx,
		`
			SELECT 
			    owner_id,
			    notifications
			FROM 
			    tgSubChecker.channels 
			JOIN
			    tgSubChecker.users 
			ON
			    tgSubChecker.channels.owner_id = tgSubChecker.users.id
			WHERE 
			    tgSubChecker.channels.id = $1
			`,
		chatId,
	).Scan(&owner.OwnerID, &owner.Notifications)
	if err != nil {
		return nil, err
	}
	return &owner, nil
}
