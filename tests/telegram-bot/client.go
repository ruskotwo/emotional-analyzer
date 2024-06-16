package main

import (
	"bytes"
	"encoding/json"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

const EaHost = "http://golang:3000"
const EaRegisterUrl = EaHost + "/register"
const EaRefreshUrl = EaHost + "/oauth/token"
const EaAddToAnalysisUrl = EaHost + "/addToAnalysis"

const Secret = "secret"

var token *oauth2.Token

func addToAnalysis(id, message, chatId string) {
	if token == nil {
		res := sendRequest(EaRegisterUrl, BodyRegister{
			CallbackUrl: "http://telegram-bot:4000/hook",
		})

		token = &oauth2.Token{}
		err := json.NewDecoder(res.Body).Decode(token)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !token.Valid() {
		res := sendRequest(EaRefreshUrl, BodyToken{
			GrantType:    "refresh_token",
			RefreshToken: token.RefreshToken,
		})

		token = &oauth2.Token{}
		err := json.NewDecoder(res.Body).Decode(token)
		if err != nil {
			log.Fatal(err)
		}
	}

	_ = sendRequest(EaAddToAnalysisUrl, BodyAddToAnalysis{
		Messages: map[string]string{
			id: message,
		},
		Secret: Secret + " " + chatId,
	})
}

type BodyAddToAnalysis struct {
	Messages map[string]string `json:"messages"`
	Secret   string            `json:"secret"`
}

type BodyRegister struct {
	CallbackUrl string `json:"callback_url"`
}

type BodyToken struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

func sendRequest(url string, body interface{}) *http.Response {
	jsonBody, _ := json.Marshal(body)

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	r.Header.Add("Content-Type", "application/json")

	if token != nil {
		r.Header.Add("Authorization", "Bearer "+token.AccessToken)
	}

	httpClient := http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	return res
}
