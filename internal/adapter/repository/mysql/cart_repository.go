package mysqlrepo

import (
	"context"

	"github.com/FatwahFir/xpanca-be/internal/domain"
	"gorm.io/gorm"
)

type CartRepository interface {
	EnsureCart(ctx context.Context, userID uint) (*domain.Cart, error)
	AddOrIncrease(ctx context.Context, userID, productID uint, qty int) error
	Inc(ctx context.Context, userID, productID uint) error
	Dec(ctx context.Context, userID, productID uint) error
	Remove(ctx context.Context, userID, productID uint) error
	GetCart(ctx context.Context, userID uint) (*domain.Cart, error)
}

type cartRepository struct{ db *gorm.DB }

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) EnsureCart(ctx context.Context, userID uint) (*domain.Cart, error) {
	var cart domain.Cart
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		FirstOrCreate(&cart, domain.Cart{UserID: userID}).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) AddOrIncrease(ctx context.Context, userID, productID uint, qty int) error {
	if qty <= 0 {
		qty = 1
	}
	cart, err := r.EnsureCart(ctx, userID)
	if err != nil {
		return err
	}

	var item domain.CartItem
	err = r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cart.ID, productID).
		First(&item).Error
	if err == gorm.ErrRecordNotFound {
		item = domain.CartItem{CartID: cart.ID, ProductID: productID, Qty: qty}
		return r.db.WithContext(ctx).Create(&item).Error
	}
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Model(&item).
		Update("qty", gorm.Expr("qty + ?", qty)).Error
}

func (r *cartRepository) Inc(ctx context.Context, userID, productID uint) error {
	return r.AddOrIncrease(ctx, userID, productID, 1)
}

func (r *cartRepository) Dec(ctx context.Context, userID, productID uint) error {
	cart, err := r.EnsureCart(ctx, userID)
	if err != nil {
		return err
	}
	var item domain.CartItem
	if err := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cart.ID, productID).
		First(&item).Error; err != nil {
		return err
	}
	if item.Qty <= 1 {
		return r.db.WithContext(ctx).Delete(&item).Error
	}
	return r.db.WithContext(ctx).
		Model(&item).
		Update("qty", gorm.Expr("qty - 1")).Error
}

func (r *cartRepository) Remove(ctx context.Context, userID, productID uint) error {
	cart, err := r.EnsureCart(ctx, userID)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cart.ID, productID).
		Delete(&domain.CartItem{}).Error
}

func (r *cartRepository) GetCart(ctx context.Context, userID uint) (*domain.Cart, error) {
	cart, err := r.EnsureCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).
		Preload("Items.Product.Images", "is_thumbnail = ?", true).
		First(cart, cart.ID).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
