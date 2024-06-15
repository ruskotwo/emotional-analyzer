package factory

import (
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/gorm/clients"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var gormSet = wire.NewSet(
	provideGormDB,
	clients.NewRepository,
)

func provideGormDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                     cfg.MysqlDSN,
		DefaultStringSize:       256,
		DontSupportRenameColumn: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cant connect to DB: %s", err)
	}

	err = db.AutoMigrate(&clients.Client{})
	if err != nil {
		log.Fatalf("Auto migrate failed: %s", err)
	}

	return db
}
