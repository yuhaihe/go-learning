package service

import (
	"context"
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"go-learning/projects/notes-api/internal/model"
	"go-learning/projects/notes-api/internal/repo"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict     = errors.New("conflict")
)

type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	FindByUsername(ctx context.Context, username string) (model.User, error)
}

type Auth struct {
	users  UserStore
	secret string
}

func NewAuth(users UserStore, secret string) *Auth {
	return &Auth{users: users, secret: secret}
}

func (a *Auth) Register(ctx context.Context, username, password string) error {
	if username == "" || utf8.RuneCountInString(password) < 6 {
		return ErrInvalidInput
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = a.users.Create(ctx, &model.User{Username: username, PasswordHash: string(hash)})
	if err != nil && (errors.Is(err, gorm.ErrDuplicatedKey) || isDuplicate(err)) {
		return ErrConflict
	}
	return err
}

func isDuplicate(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "Duplicate") || strings.Contains(msg, "UNIQUE")
}

func (a *Auth) Login(ctx context.Context, username, password string) (string, error) {
	user, err := a.users.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return "", ErrUnauthorized
		}
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return "", ErrUnauthorized
	}
	claims := jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(a.secret))
}
