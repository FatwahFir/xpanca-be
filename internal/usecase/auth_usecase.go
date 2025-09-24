package usecase

import (
	"context"
	"errors"

	mysqlrepo "github.com/FatwahFir/xpanca-be/internal/adapter/repository/mysql"
	"github.com/FatwahFir/xpanca-be/internal/domain"
	"github.com/FatwahFir/xpanca-be/pkg/jwtx"
	"github.com/FatwahFir/xpanca-be/pkg/password"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthUsecase struct {
	users mysqlrepo.UserRepository
	jwtm  *jwtx.Manager
}

func NewAuthUsecase(users mysqlrepo.UserRepository, jwtm *jwtx.Manager) *AuthUsecase {
	return &AuthUsecase{users: users, jwtm: jwtm}
}

func (uc *AuthUsecase) Login(ctx context.Context, username, pass string) (string, *domain.User, error) {
	u, err := uc.users.FindByUsername(ctx, username)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}
	if !password.Check(pass, u.Password) {
		return "", nil, ErrInvalidCredentials
	}
	token, err := uc.jwtm.Generate(u.ID, u.Username, u.Role)
	if err != nil {
		return "", nil, err
	}
	return token, u, nil
}
