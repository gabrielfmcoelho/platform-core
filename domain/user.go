package domain

import (
	"context"

	"gorm.io/gorm"
)

// MANY TO ONE WITH ORGANIZATION

type User struct {
	gorm.Model
	Email          string           `gorm:"size:255;uniqueIndex;not null"`
	Password       string           `gorm:"size:255;not null"`
	OrganizationID uint             `gorm:"not nul;Index"`
	Organization   Organization     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Relationship to Organization
	RoleID         uint             `gorm:"not null;Index"`
	Role           UserRole         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Relationship to UserRole
	Bio            UserBio          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Metrics        UserMetrics      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Configs        UserConfig       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logs           []UserLog        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ServiceLogs    []UserServiceLog `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateUser struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	OrganizationID uint   `json:"organization_id" binding:"required"`
	RoleID         uint   `json:"role" binding:"required"`
}

type PublicUser struct {
	ID               uint   `json:"id"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	OrganizationID   uint   `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	RoleID           uint   `json:"role_id"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, userID uint, user *User) error
	Archive(ctx context.Context, userID uint) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *CreateUser) error
	Fetch(ctx context.Context) ([]PublicUser, error)
	GetByIdentifier(ctx context.Context, identifier string) (PublicUser, error)
	Update(ctx context.Context, userID uint, user *User) error
	Archive(ctx context.Context, userID uint) error
}

// EXEMPLE TIP: To access Bio from a User, use the following:

// var user domain.User
// err := db.Preload("Bio").First(&user, 1).Error
// if err != nil {
// 	return err
// }
