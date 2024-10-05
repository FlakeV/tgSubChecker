package main

import (
	"context"
	"encoding/json"
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

func getChatMemberUpdates(offset int, BOT_TOKEN string) *models.Updates {

	var url = API_URL + BOT_TOKEN + "/getUpdates?offset=" + strconv.Itoa(offset) + "&limit=1&allowed_updates=%5B%22chat_member%22%5D&timeout=100"
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

	for {
		resp := getChatMemberUpdates(offset, appConfig.BotToken)
		if len(resp.Updates) != 0 {
			saver.NewSub(ctx, &resp.Updates[0])
			offset = resp.Updates[0].UpdateID + 1
		}
		for _, update := range resp.Updates {
			log.Println(update.UpdateID)
		}
	}
}

// getChatMemberUpdates(0)
