package mysqlrepo

import (
	"context"

	"github.com/FatwahFir/xpanca-be/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
}

type userRepositoryImpl struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) UserRepository { return &userRepositoryImpl{db: db} }

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var u domain.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
