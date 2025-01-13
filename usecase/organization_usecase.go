package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/gabrielfmcoelho/platform-core/domain"
)

type organizationUsecase struct {
	repo domain.OrganizationRepository
}

// NewOrganizationUsecase retorna uma instância que implementa a interface OrganizationUsecase
func NewOrganizationUsecase(repo domain.OrganizationRepository) domain.OrganizationUsecase {
	return &organizationUsecase{
		repo: repo,
	}
}

// Create cria uma nova organização
func (uc *organizationUsecase) Create(ctx context.Context, organization *domain.Organization) error {
	// Poderia validar se a Role existe, etc.
	if organization.Name == "" {
		return errors.New("organization name is required")
	}

	// Criar no repositório
	if err := uc.repo.Create(ctx, organization); err != nil {
		return err
	}
	return nil
}

// Fetch retorna todas as organizações, convertendo para PublicOrganization
func (uc *organizationUsecase) Fetch(ctx context.Context) ([]domain.PublicOrganization, error) {
	orgs, err := uc.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.PublicOrganization
	for _, org := range orgs {
		result = append(result, uc.toPublicOrganization(org))
	}
	return result, nil
}

// GetByIdentifier busca organização por ID ou Nome (exemplo)
func (uc *organizationUsecase) GetByIdentifier(ctx context.Context, identifier string) (domain.PublicOrganization, error) {
	// tentar converter identifier para uint
	var org domain.Organization
	var err error

	// se for ID numérico, busca por ID; senão, busca por Nome
	var id uint
	_, scanErr := fmt.Sscanf(identifier, "%d", &id)
	if scanErr == nil {
		org, err = uc.repo.GetByID(ctx, id)
	} else {
		org, err = uc.repo.GetByName(ctx, identifier)
	}
	if err != nil {
		return domain.PublicOrganization{}, err
	}
	return uc.toPublicOrganization(org), nil
}

// GetUsers retorna a lista de usuários da organização, convertendo para PublicUser
func (uc *organizationUsecase) GetUsers(ctx context.Context, id uint) ([]domain.PublicUser, error) {
	users, err := uc.repo.GetUsers(ctx, id)
	if err != nil {
		return nil, err
	}

	var publicUsers []domain.PublicUser
	for _, u := range users {
		publicUsers = append(publicUsers, domain.PublicUser{
			ID:               u.ID,
			Email:            u.Email,
			FirstName:        "", // se tiver esse campo no domain.User, pode mapear
			OrganizationID:   u.OrganizationID,
			OrganizationName: "", // se quiser, busque org e retorne
			RoleID:           u.RoleID,
		})
	}
	return publicUsers, nil
}

// GetSubscribedServices retorna os serviços que a org está inscrita
func (uc *organizationUsecase) GetSubscribedServices(ctx context.Context, id uint) ([]domain.PublicService, error) {
	return uc.repo.GetSubscribedServices(ctx, id)
}

// Update atualiza uma organização
func (uc *organizationUsecase) Update(ctx context.Context, organizationID uint, organization *domain.Organization) error {
	return uc.repo.Update(ctx, organizationID, organization)
}

// Delete remove a organização
func (uc *organizationUsecase) Delete(ctx context.Context, organizationID uint) error {
	return uc.repo.Delete(ctx, organizationID)
}

// ==================================================
// Helpers
// ==================================================

func (uc *organizationUsecase) toPublicOrganization(org domain.Organization) domain.PublicOrganization {
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
