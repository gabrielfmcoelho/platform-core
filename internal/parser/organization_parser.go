package parser

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
)

// Parse Organization to PublicOrganization
func ToPublicOrganization(org domain.Organization) domain.PublicOrganization {
	return domain.PublicOrganization{
		ID:       org.ID,
		Name:     org.Name,
		Nickname: org.Nickname,
		LogoUrl:  org.LogoUrl,
	}
}
