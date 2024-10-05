package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"tgSubChecker/internal/repo"
)

type Reader struct {
	pool *pgxpool.Pool
}

func NewReader(pool *pgxpool.Pool) repo.Reader {
	return &Reader{pool: pool}
}

func (r *Reader) GetOwner(ctx context.Context, chatId int) (int, error) {
	var ownerID int
	err := r.pool.QueryRow(ctx, "SELECT owner_id FROM tgSubChecker.channels WHERE id = $1", chatId).Scan(&ownerID)
	if err != nil {
		return 0, err
	}
	return ownerID, nil
}
