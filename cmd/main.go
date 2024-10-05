package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"

	"tgSubChecker/internal/config"
	"tgSubChecker/internal/models"
	"tgSubChecker/internal/repo/postgres"
)

const (
	API_URL = "https://api.telegram.org/bot"
)

var (
	offset int
)

func getChatMemberUpdates(offset int, BotToken string) *models.Updates {

	var url = API_URL + BotToken + "/getUpdates?offset=" + strconv.Itoa(offset) + "&limit=1&allowed_updates=%5B%22chat_member%22%5D&timeout=100"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Get member updates error: ", err)
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read data member updates error: ", err)
	}
	var updates models.Updates
	err = json.Unmarshal(jsonData, &updates)
	if err != nil {
		log.Fatal("Unmarshal member updates error: ", err)
	}
	return &updates
}

func sendMessageToOwner(update *models.Update, BotToken string, ownerID int) {
	var msgText string
	if update.ChatMember.NewChatMember.Status == "left" {
		msgText = fmt.Sprintf(
			"Пользователь %s %s (%d) покинул чат",
			update.ChatMember.NewChatMember.User.FirstName,
			update.ChatMember.NewChatMember.User.LastName,
			update.ChatMember.NewChatMember.User.ID,
		)
	} else if update.ChatMember.NewChatMember.Status == "member" {
		msgText = fmt.Sprintf(
			"Пользователь %s %s (%d) присоединился к чату",
			update.ChatMember.NewChatMember.User.FirstName,
			update.ChatMember.NewChatMember.User.LastName,
			update.ChatMember.NewChatMember.User.ID,
		)
	}
	var url = API_URL + BotToken + "/sendMessage?chat_id=" + strconv.Itoa(ownerID) + "&text=" + msgText
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Send message error: ", err)
	}
	defer resp.Body.Close()
}

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(appConfig)
	ctx := context.Background()

	dbConfig, err := pgxpool.ParseConfig(appConfig.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	saver := postgres.NewSaver(dbPool)
	reader := postgres.NewReader(dbPool)

	for {
		resp := getChatMemberUpdates(offset, appConfig.BotToken)
		if len(resp.Updates) != 0 {
			err = saver.NewSub(ctx, &resp.Updates[0])
			if err != nil {
				log.Fatal(err)
			}

			ownerId, err := reader.GetOwner(ctx, resp.Updates[0].ChatMember.Chat.ID)
			if err != nil {
				log.Fatal(err)
			}

			sendMessageToOwner(&resp.Updates[0], appConfig.BotToken, ownerId)
			offset = resp.Updates[0].UpdateID + 1
		}
		for _, update := range resp.Updates {
			log.Println(update.UpdateID)
		}
	}
}

// getChatMemberUpdates(0)
