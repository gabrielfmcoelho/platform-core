package seeds

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

// SeedServices cria os serviços padrão (Resistracker, ...) se não existirem
func SeedServices(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.Service{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		services := []domain.Service{
			{
				MarketingName: "Resistracker",
				Name:          "Resistracker",
				Description:   "Acompanhamento de registros",
				AppUrl:        "https://resistracker.solude.tech",
				IconUrl:       "https://resistracker.solude.tech/favicon.ico",
				ScreenshotUrl: "https://resistracker.solude.tech/screenshot.png",
				TagLine:       "Gestão Inteligente de Resistência Bacteriana",
				Benefits:      "Redução de 40% no tempo de identificação de padrões de resistência;Aumento de 60% na eficácia do tratamento inicial;Economia de 30% nos custos com antibióticos",
				Features:      "Monitoramento em tempo real;Análise preditiva de resistência;Suporte à decisão clínica;Relatórios personalizados",
				Tags:          "IA;Microbiologia;Antibióticos",
				LastUpdate:    "2021-09-01",
				Status:        "Online",
				Price:         0.00,
			},
		}

		if err := db.Create(&services).Error; err != nil {
			return err
		}

		log.Printf("[SeedServices] Criados %d serviços (Resistracker, ...)\n", len(services))
	}
	return nil
}
