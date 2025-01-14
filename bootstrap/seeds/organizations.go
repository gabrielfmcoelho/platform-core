package seeds

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

// SeedOrganizations cria registros iniciais de organização, se não existirem
func SeedOrganizations(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.Organization{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		// Exemplo: suposição de que já exista um "OrganizationRole" com ID=1 e ID=2

		var resistracker domain.Service
		if err := db.Where("name = ?", "Resistracker").First(&resistracker).Error; err != nil {
			return err
		}

		orgs := []domain.Organization{
			{
				Name:               "Solude",
				LogoUrl:            "https://example.com/logo-acme.png",
				RoleID:             1, // por exemplo, se 1 = "Admin"
				SubscribedServices: []domain.Service{resistracker},
			},
			{
				Name:               "Hospital São Marcos",
				LogoUrl:            "https://example.com/logo-beta.png",
				RoleID:             2, // por exemplo, se 2 = "Manager"
				SubscribedServices: []domain.Service{resistracker},
			},
			{
				Name:    "Hospital Universitário - UFPI",
				LogoUrl: "https://example.com/logo-beta.png",
				RoleID:  2,
			},
			{
				Name:    "Guest Organization",
				LogoUrl: "https://example.com/logo-beta.png",
				RoleID:  3, // por exemplo, se 3 = "Guest"
			},
		}

		if err := db.Create(&orgs).Error; err != nil {
			return err
		}

		log.Printf("[SeedOrganizations] Criadas %d organizações (Acme, Beta)\n", len(orgs))
	}
	return nil
}
