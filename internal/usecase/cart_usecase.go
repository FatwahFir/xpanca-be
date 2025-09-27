package usecase

import (
	"context"

	mysqlrepo "github.com/FatwahFir/xpanca-be/internal/adapter/repository/mysql"
	"github.com/FatwahFir/xpanca-be/internal/domain"
)

type CartUsecase struct{ repo mysqlrepo.CartRepository }

func NewCartUsecase(r mysqlrepo.CartRepository) *CartUsecase { return &CartUsecase{repo: r} }

func (uc *CartUsecase) Add(ctx context.Context, userID, productID uint, qty int) error {
	return uc.repo.AddOrIncrease(ctx, userID, productID, qty)
}

func (uc *CartUsecase) Inc(ctx context.Context, userID, productID uint) error {
	return uc.repo.Inc(ctx, userID, productID)
}

func (uc *CartUsecase) Dec(ctx context.Context, userID, productID uint) error {
	return uc.repo.Dec(ctx, userID, productID)
}

func (uc *CartUsecase) Remove(ctx context.Context, userID, productID uint) error {
	return uc.repo.Remove(ctx, userID, productID)
}

func (uc *CartUsecase) Get(ctx context.Context, userID uint) (*domain.Cart, error) {
	return uc.repo.GetCart(ctx, userID)
}
