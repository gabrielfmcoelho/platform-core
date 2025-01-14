package bootstrap

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/bootstrap/seeds"
	"gorm.io/gorm"
)

// RunSeeds é a função principal que chama cada seeder específico
func RunSeeds(db *gorm.DB) {
	// Podemos rodar dentro de uma transação, caso deseje atomicidade
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := seeds.SeedServices(tx); err != nil {
			return err
		}

		if err := seeds.SeedOrganizationRoles(tx); err != nil {
			return err
		}

		if err := seeds.SeedOrganizations(tx); err != nil {
			return err
		}

		if err := seeds.SeedUserRoles(tx); err != nil {
			return err
		}

		if err := seeds.SeedUsers(tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Printf("Erro ao rodar seeds: %v", err)
	} else {
		log.Printf("Seeds executados com sucesso!")
	}
}
