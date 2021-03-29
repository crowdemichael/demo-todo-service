package schema

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
	DeletedAt gorm.DeletedAt
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		User{},
	)
}
