package seeds

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

// SeedUserRoles cria os user roles (Admin, Manager, User, Guest) se não existirem
func SeedUserRoles(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.UserRole{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		roles := []domain.UserRole{
			{RoleName: "Admin"},
			{RoleName: "Manager"},
			{RoleName: "User"},
			{RoleName: "Guest"},
		}

		if err := db.Create(&roles).Error; err != nil {
			return err
		}

		log.Printf("[SeedUserRoles] Criados %d UserRoles (Admin, Manager, User, Guest)\n", len(roles))
	}
	return nil
}

// SeedOrganizationRoles cria os organization roles (Admin, Manager, User, Guest) se não existirem
func SeedOrganizationRoles(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.OrganizationRole{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		roles := []domain.OrganizationRole{
			{RoleName: "Admin"},
			{RoleName: "Hospital"},
			{RoleName: "Guest"},
		}

		if err := db.Create(&roles).Error; err != nil {
			return err
		}

		log.Printf("[SeedOrganizationRoles] Criados %d OrganizationRoles (Admin, Manager, User, Guest)\n", len(roles))
	}
	return nil
}
