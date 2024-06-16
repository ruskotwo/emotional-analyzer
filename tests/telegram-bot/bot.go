package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

var bot *tgbotapi.BotAPI

func startBot() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message != nil && update.Message.Text != "" {
				addToAnalysis(
					strconv.Itoa(update.Message.MessageID),
					update.Message.Text,
					strconv.FormatInt(update.Message.Chat.ID, 10),
				)
			}
		}
	}()
}
