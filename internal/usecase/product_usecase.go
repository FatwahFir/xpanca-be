package usecase

import (
	"context"

	mysqlrepo "github.com/FatwahFir/xpanca-be/internal/adapter/repository/mysql"
	"github.com/FatwahFir/xpanca-be/internal/domain"
	"github.com/FatwahFir/xpanca-be/internal/dto"
)

type ProductUsecase struct {
	repo mysqlrepo.ProductRepository
}

func NewProductUsecase(repo mysqlrepo.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) Find(ctx context.Context, q dto.ProductQuery) ([]domain.Product, int64, error) {
	return uc.repo.Find(ctx, q)
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}
