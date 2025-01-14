package internal

import (
	"strconv"
	"unicode"

	"github.com/gabrielfmcoelho/platform-core/domain"
)

// ParseUint parses a number from a string and returns a uint
func ParseUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, domain.ErrInvalidNumberToParse
	}
	return uint(i), nil
}

// Parse domain.User to domain.PublicUser
func ParsePublicUser(user domain.User) domain.PublicUser {
	return domain.PublicUser{
		ID:               user.ID,
		Email:            user.Email,
		FirstName:        user.Bio.FirstName,
		OrganizationID:   user.Organization.ID,
		OrganizationName: user.Organization.Name,
		RoleID:           user.Role.ID,
	}
}

func ParseCreateUser(createUser *domain.CreateUser) *domain.User {
	return &domain.User{
		Email:          createUser.Email,
		Password:       createUser.Password,
		OrganizationID: createUser.OrganizationID,
		RoleID:         createUser.RoleID,
	}
}

func ParsePublicOrganization(org domain.Organization) domain.PublicOrganization {
	// Converte domain.Organization -> domain.PublicOrganization
	// Carrega Users
	var publicUsers []domain.PublicUser
	for _, u := range org.Users {
		publicUsers = append(publicUsers, domain.PublicUser{
			ID:               u.ID,
			Email:            u.Email,
			FirstName:        "", // mapeie se tiver no domain.User
			OrganizationID:   u.OrganizationID,
			OrganizationName: org.Name,
			RoleID:           u.RoleID,
		})
	}

	// Converte SubscribedServices
	var publicServices []domain.PublicService
	for _, srv := range org.SubscribedServices {
		publicServices = append(publicServices, domain.PublicService{
			ID:   srv.ID,
			Name: srv.Name,
		})
	}

	return domain.PublicOrganization{
		ID:                 org.ID,
		Name:               org.Name,
		LogoUrl:            org.LogoUrl,
		Users:              publicUsers,
		SubscribedServices: publicServices,
	}
}

// parsePublicService converte de domain.Service para domain.PublicService
func ParsePublicService(s domain.Service) domain.PublicService {
	return domain.PublicService{
		ID:   s.ID,
		Name: s.Name,
	}
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
