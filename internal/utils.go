package internal

import (
	"strconv"

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
