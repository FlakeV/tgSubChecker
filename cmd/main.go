package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"tgSubChecker/internal/repo"
	"tgSubChecker/internal/repo/csvStorage"

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

func sendMessageToOwner(update *models.Update, BotToken string, ownerID int, wg *sync.WaitGroup) {
	defer wg.Done()
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
	log.Println("Starting app tgSubChecker")
	defer log.Println("Finish app tgSubChecker")
	wg := sync.WaitGroup{}
	defer wg.Wait()

	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	var reader repo.Reader
	var saver repo.Saver

	dbConfig, err := pgxpool.ParseConfig(appConfig.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		log.Println(err)
		saver = csvStorage.NewSaver()
		reader = csvStorage.NewReader()
	} else {
		saver = postgres.NewSaver(dbPool)
		reader = postgres.NewReader(dbPool)
	}
	defer dbPool.Close()

	for {
		resp := getChatMemberUpdates(offset, appConfig.BotToken)
		if len(resp.Updates) != 0 {
			err = saver.NewSub(ctx, &resp.Updates[0])
			if err != nil {
				log.Fatal(err)
			}

			owner, err := reader.GetOwner(ctx, resp.Updates[0].ChatMember.Chat.ID)
			if err != nil {
				log.Fatal(err)
			}

			if owner.Notifications {
				wg.Add(1)
				go sendMessageToOwner(&resp.Updates[0], appConfig.BotToken, owner.OwnerID, &wg)
			}

			offset = resp.Updates[0].UpdateID + 1
		}
		for _, update := range resp.Updates {
			log.Println(update.UpdateID)
		}
	}
}

// getChatMemberUpdates(0)
