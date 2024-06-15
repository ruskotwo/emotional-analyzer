package clients

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	CallbackUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
