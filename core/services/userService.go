package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/tools"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domains.UserRepository
}

func NewUserService(userRepository domains.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, username, password, email string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return tools.PasswordHashErr
	}

	user := domains.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashed),
		Status:       "feel happy",
		OnlineStatus: domains.StateOnline,
	}

	if err := s.repo.AddUser(ctx, user); err != nil {
		return err
	}
	return nil
}

type CustomeClaim struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", tools.LoginErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", tools.LoginErr
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomeClaim{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	secretKey := os.Getenv("SECRET_KEY")
	return token.SignedString([]byte(secretKey))
}

func (s *UserService) ForgotPasswordSendCode(ctx context.Context, email string) error {

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	var code string
	skip := true

	if user.ExpireForgotToken != nil {
		if user.ExpireForgotToken.After(time.Now()) {
			code = *user.ForgotToken
			skip = false
		}
	}
	if skip {
		b := make([]byte, 8)
		rand.Read(b)
		code = fmt.Sprintf("%x", b)
	}
	if err := s.repo.ChangeForgotCodeUser(ctx, user.ID, code, time.Now().Add(time.Hour*1)); err != nil {
		return err
	}

	body := fmt.Sprintf("recovery code: %s", code)
	if err := tools.NotificationSender(user.Email, "Recover Password From Aura", body); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ForgotPasswordRecovery(ctx context.Context, email, newPassword, token string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user.ExpireForgotToken.Before(time.Now()) {
		return tools.TokenExpiredErr
	}
	if user.ForgotToken != &token {
		return tools.TokenNotEqualErr
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(newPassword)); err != nil {
		return tools.NewPasswordSameAsOldErr
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.ChangePasswordUser(ctx, user.ID, string(hashed))
}
