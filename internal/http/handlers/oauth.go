package handlers

import (
	"encoding/json"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/gorm/clients"
	"net/http"
	"net/url"
	"strings"
)

type OAuthHandler struct {
	config            *config.Config
	oAuth2            *server.Server
	clientsRepository *clients.Repository
}

func NewOAuthHandler(
	cfg *config.Config,
	oAuth2 *server.Server,
) *OAuthHandler {
	return &OAuthHandler{
		config: cfg,
		oAuth2: oAuth2,
	}
}

type BodyToken struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

func (h OAuthHandler) HandleToken(w http.ResponseWriter, r *http.Request) {
	if strings.Index(r.Header.Get("Content-Type"), "application/json") != -1 {
		body := BodyToken{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		r.Form = url.Values{}
		r.Form.Set("grant_type", body.GrantType)
		r.Form.Set("refresh_token", body.RefreshToken)
	}

	err := h.oAuth2.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
