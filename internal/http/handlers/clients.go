package handlers

import (
	"encoding/json"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/gorm/clients"
	"net/http"
	"strconv"
)

type ClientsHandler struct {
	config            *config.Config
	oAuth2            *server.Server
	clientsRepository *clients.Repository
}

func NewClientsHandler(
	cfg *config.Config,
	oAuth2 *server.Server,
	clientsRepository *clients.Repository,
) *ClientsHandler {
	return &ClientsHandler{
		config:            cfg,
		oAuth2:            oAuth2,
		clientsRepository: clientsRepository,
	}
}

type BodyRegister struct {
	CallbackUrl string `json:"callback_url"`
}

func (h ClientsHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := &BodyRegister{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.CallbackUrl == "" {
		http.Error(w, "Parameter callback_url is required", http.StatusBadRequest)
		return
	}

	client, err := h.clientsRepository.Create(body.CallbackUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := h.generateToken(r, client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(h.oAuth2.GetTokenData(token))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h ClientsHandler) generateToken(request *http.Request, client clients.Client) (oauth2.TokenInfo, error) {
	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     h.config.OAuth2.Client.ID,
		ClientSecret: h.config.OAuth2.Client.Secret,
		UserID:       strconv.Itoa(int(client.ID)),
		Request:      request,
	}

	return h.oAuth2.Manager.GenerateAccessToken(request.Context(), oauth2.ClientCredentials, tgr)
}
