package factory

import (
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	http2 "github.com/ruskotwo/emotional-analyzer/internal/http"
	"github.com/ruskotwo/emotional-analyzer/internal/http/handlers"
	"net/http"
	"time"
)

var httpSet = wire.NewSet(
	provideOAuth2,
	http2.NewServer,
	handlers.NewAnalysisHandler,
	handlers.NewClientsHandler,
	handlers.NewOAuthHandler,
)

func provideOAuth2(cfg *config.Config) *server.Server {
	manager := manage.NewDefaultManager()

	// use mysql token store
	tokenStore := mysql.NewDefaultStore(
		mysql.NewConfig(cfg.MysqlDSN),
	)

	manager.MapTokenStorage(tokenStore)

	// client memory store
	clientStore := store.NewClientStore()
	_ = clientStore.Set(cfg.OAuth2.Client.ID, &models.Client{
		ID:     cfg.OAuth2.Client.ID,
		Secret: cfg.OAuth2.Client.Secret,
		Domain: cfg.OAuth2.Client.Domain,
	})
	manager.MapClientStorage(clientStore)

	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    2 * time.Hour,
		RefreshTokenExp:   48 * time.Hour,
		IsGenerateRefresh: true,
	})

	serv := server.NewServer(server.NewConfig(), manager)

	serv.SetClientInfoHandler(func(r *http.Request) (string, string, error) {
		return cfg.OAuth2.Client.ID, cfg.OAuth2.Client.Secret, nil
	})

	return serv
}
