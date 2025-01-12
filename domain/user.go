package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	IsArchived   bool         `gorm:"not null;default:false"` // Soft delete only removes from the system but keeps the data
	Email        string       `gorm:"size:255;uniqueIndex;not null"`
	Password     string       `gorm:"size:255;not null"`
	Organization Organization `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleID       uint
	Role         UserRole
	Bio          UserBio          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Metrics      UserMetrics      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Configs      UserConfig       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logs         []UserLog        `gorm:"foreignKey:UserID"`
	ServiceLogs  []UserServiceLog `gorm:"foreignKey:UserID"`
}

type CreateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Access   int    `json:"access" binding:"required"`
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
	Create(ctx context.Context, user *User) error
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
