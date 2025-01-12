package bootstrap

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		// add domains like: &domain.User{},
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
}
