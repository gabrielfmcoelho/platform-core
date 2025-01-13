// bootstrap/seeds/user_seed.go
package seeds

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

// SeedUsers popula a tabela de usuários com dados iniciais
func SeedUsers(db *gorm.DB) error {
	// Exemplo: Verificar se já existe algum usuário
	var count int64
	if err := db.Model(&domain.User{}).Count(&count).Error; err != nil {
		return err
	}

	// Se não tiver usuário nenhum, cria alguns
	if count == 0 {
		users := []domain.User{
			{
				Email:          "admin@example.com",
				Password:       "admin123",
				OrganizationID: 1,
				RoleID:         1, // Ex: Role Admin
			},
			{
				Email:          "user@example.com",
				Password:       "user123",
				OrganizationID: 2,
				RoleID:         2, // Ex: Role User
			},
		}

		if err := db.Create(&users).Error; err != nil {
			return err
		}

		log.Printf("Users seed executado: criados %d usuários\n", len(users))
	}
	return nil
}
