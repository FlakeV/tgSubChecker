package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"tgSubChecker/internal/models"
	"tgSubChecker/internal/repo"
)

type Saver struct {
	pool *pgxpool.Pool
}

func NewSaver(pool *pgxpool.Pool) repo.Saver {
	return &Saver{pool: pool}
}

func (s *Saver) NewSub(ctx context.Context, update *models.Update) error {
	log.Println("new sub", update.UpdateID)
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	querySub := `
	INSERT INTO tgSubChecker.subscribers
		(
			id, 
			username, 
			first_name, 
			last_name,
			invate_link,
			is_bot,
			is_premium
		)
	VALUES
		($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (id) DO NOTHING
	`

	_, err = tx.Exec(ctx, querySub,
		update.ChatMember.NewChatMember.User.ID,
		update.ChatMember.NewChatMember.User.Username,
		update.ChatMember.NewChatMember.User.FirstName,
		update.ChatMember.NewChatMember.User.LastName,
		update.ChatMember.InviteLink.InviteLink,
		update.ChatMember.NewChatMember.User.IsBot,
		update.ChatMember.NewChatMember.User.IsPremium,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}

	queryEvent := `
	INSERT INTO tgSubChecker.sub_events
		(
			subscriber_id, 
			event_type, 
			event_time,
			chat_id
		)
	VALUES
		($1, $2, $3, $4)
	`

	_, err = tx.Exec(ctx, queryEvent,
		update.ChatMember.NewChatMember.User.ID,
		repo.EventType[update.ChatMember.NewChatMember.Status],
		time.Unix(int64(update.ChatMember.Date), 0),
		update.ChatMember.Chat.ID,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
