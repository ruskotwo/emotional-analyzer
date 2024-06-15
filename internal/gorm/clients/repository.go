package clients

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetByID(id uint) (client *Client, err error) {
	err = r.db.Take(&client, id).Error
	return
}

func (r Repository) Create(callbackUrl string) (client Client, err error) {
	client.CallbackUrl = callbackUrl

	err = r.db.Create(&client).Error

	return
}
