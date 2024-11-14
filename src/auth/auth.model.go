package auth

import (
	"time"
	"url-shortner/src/url"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	URLs []url.URL `gorm:"foreignKey:UserId"`
}
