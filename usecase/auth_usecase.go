package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"github.com/gabrielfmcoelho/platform-core/internal/password"
	"github.com/gabrielfmcoelho/platform-core/internal/tokenutil"
)

type AuthUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(userRepository domain.UserRepository, timeout time.Duration) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (au *AuthUsecase) LoginUserByEmail(c context.Context, email string, rawPassword string, accessSecret string, accessExpiry int, refreshSecret string, refreshExpiry int) (loginResponse *domain.LoginResponse, err error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout) // This creates a new context with a timeout and a cancel function, which should be called at the end of the function to release resources
	defer cancel()

	user, err := au.userRepository.GetByEmail(ctx, email)
	// if the user is not found, return an error with a message
	if err != nil {
		if !errors.Is(err, domain.ErrUserEmailNotFound) {
			return nil, domain.ErrInternalServerError
		}
		return nil, err
	}

	// verify if the password is match
	err = password.VerifyPassword(user.Password, rawPassword)
	if err != nil {
		return nil, err
	}

	// create access token
	accessToken, err := au.CreateAccessToken(&user, accessSecret, accessExpiry)
	if err != nil {
		return nil, err
	}

	// create refresh token
	refreshToken, err := au.CreateRefreshToken(&user, refreshSecret, refreshExpiry)
	if err != nil {
		return nil, err
	}

	// return the login response
	return &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (au *AuthUsecase) LoginGuestUser(c context.Context, accessSecret string, accessExpiry int, refreshSecret string, refreshExpiry int) (loginResponse *domain.LoginResponse, err error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout) // This creates a new context with a timeout and a cancel function, which should be called at the end of the function to release resources
	defer cancel()

	incomingIP := ctx.Value("ip").(string)
	log.Println("Incoming IP:", incomingIP)

	user, err := au.userRepository.GetByID(ctx, 1)
	if err != nil {
		return nil, err
	}

	// create access token
	accessToken, err := au.CreateAccessToken(&user, "accessSecret", 1)
	if err != nil {
		return nil, err
	}

	// create refresh token
	refreshToken, err := au.CreateRefreshToken(&user, "refreshSecret", 1)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (au *AuthUsecase) CreateAccessToken(user *domain.User, accessSecret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, accessSecret, expiry)
}

func (au *AuthUsecase) CreateRefreshToken(user *domain.User, refreshSecret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, refreshSecret, expiry)
}

func (au *AuthUsecase) ForgotPassword(c context.Context, email string) (err error) {
	//ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	//defer cancel()

	//user, err := au.userRepository.GetByEmail(ctx, email)
	//if err != nil {
	//	if !errors.Is(err, domain.ErrUserEmailNotFound) {
	//		return domain.ErrInternalServerError
	//	}
	//	return err
	//}

	// send email with the reset password link
	return nil
}

func (au *AuthUsecase) ResetPassword(c context.Context, email string, newRawPassword string) (err error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	user, err := au.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, domain.ErrUserEmailNotFound) {
			return domain.ErrInternalServerError
		}
		return err
	}

	// update user password
	user.Password, err = password.HashPassword(newRawPassword)
	if err != nil {
		return err
	}

	err = au.userRepository.Update(ctx, user.ID, &user)
	if err != nil {
		return err
	}

	// send email with the password reset confirmation
	return nil
}
