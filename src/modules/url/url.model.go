package url

import (
	auth "url-shortner/src/modules/auth"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	UserID      uint32
	User        auth.User `gorm:"foreignKey:UserID"`
	OriginalUrl string    `gorm:"type:varchar(120); not null;"`
	Code        string    `gorm:"type:varchar(10); not null;index"`
}
