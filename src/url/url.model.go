package url

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID          uint32 `gorm:"primaryKey"`
	UserId      uint32 `gorm:"index"`
	OriginalUrl string `gorm:"type:varchar(120); not null;"`
	Code        string `gorm:"type:varchar(10); unique; not null;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
