package parser

import (
	"github.com/gabrielfmcoelho/platform-core/domain"
)

// Parse User to PublicUser
func ToPublicUser(u domain.User) domain.PublicUser {
	return domain.PublicUser{
		ID:               u.ID,
		Email:            u.Email,
		FirstName:        u.Bio.FirstName,
		OrganizationID:   u.Organization.ID,
		OrganizationName: u.Organization.Name,
		RoleID:           u.Role.ID,
	}
}

// Parse CreateUser to User
func ToUser(cu *domain.CreateUser) *domain.User {
	return &domain.User{
		Email:          cu.Email,
		Password:       cu.Password,
		OrganizationID: cu.OrganizationID,
		RoleID:         cu.RoleID,
	}
}
