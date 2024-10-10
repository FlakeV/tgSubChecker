package csvStorage

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tgSubChecker/internal/models"
	"time"
)

type Saver struct {
	eventsFile *os.File
	subFile    *os.File
}

func NewSaver() *Saver {
	eventsFile, err := os.OpenFile("temp/events.csv", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}
	_, err = eventsFile.Seek(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	subFile, err := os.OpenFile("temp/subscribers.csv", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}
	_, err = subFile.Seek(0, 2)
	if err != nil {
		log.Fatal(err)
	}

	return &Saver{
		eventsFile: eventsFile,
		subFile:    subFile,
	}
}

func (s *Saver) writeSub(ctx context.Context, update *models.Update) error {
	wr := csv.NewWriter(s.subFile)
	defer wr.Flush()

	unixDate := time.Unix(int64(update.ChatMember.Date), 0).Format("2006-01-02 15:04:05")

	data := []string{
		strconv.Itoa(update.UpdateID),
		unixDate,
		strconv.Itoa(update.ChatMember.NewChatMember.User.ID),
		update.ChatMember.NewChatMember.Status,
		update.ChatMember.NewChatMember.User.Username,
		update.ChatMember.NewChatMember.User.FirstName,
		update.ChatMember.NewChatMember.User.LastName,
		update.ChatMember.InviteLink.InviteLink,
		strconv.FormatBool(update.ChatMember.NewChatMember.User.IsBot),
		strconv.FormatBool(update.ChatMember.NewChatMember.User.IsPremium),
	}

	err := wr.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *Saver) NewSub(ctx context.Context, update *models.Update) error {
	return s.writeSub(ctx, update)
}
