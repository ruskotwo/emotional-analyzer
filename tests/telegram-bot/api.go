package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func startServer() {
	http.HandleFunc("/hook", handleHook)

	log.Printf("Listening on :4000...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":4000"), nil))
}

type BodyHook struct {
	Messages map[string]string `json:"messages"`
	Secret   string            `json:"secret"`
}

func handleHook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := &BodyHook{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var secret, chatId string

	words := strings.Fields(body.Secret)
	for idx, word := range words {
		switch idx {
		case 0:
			secret = word
		case 1:
			chatId = word
		}
	}

	if secret != Secret {
		http.Error(w, "Secret not ok", http.StatusUnauthorized)
		return
	}

	chatID, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		http.Error(w, "Secret is invalid", http.StatusBadRequest)
		return
	}

	for key, sentiment := range body.Messages {
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Printf("Invalid message key %s\n", key)
			continue
		}

		msg := tgbotapi.NewMessage(chatID, sentiment)
		msg.ReplyToMessageID = id
		_, err = bot.Send(msg)
		if err != nil {
			http.Error(w, "Cant send sentiment", http.StatusInternalServerError)
			return
		}
	}
}
