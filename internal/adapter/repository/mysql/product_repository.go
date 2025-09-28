package mysqlrepo

import (
	"context"
	"strings"

	"github.com/FatwahFir/xpanca-be/internal/domain"
	"github.com/FatwahFir/xpanca-be/internal/dto"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Find(ctx context.Context, q dto.ProductQuery) ([]domain.Product, int64, error)
	GetByID(ctx context.Context, id uint) (*domain.Product, error)
}

type productRepositoryImpl struct{ db *gorm.DB }

func NewProductRepo(db *gorm.DB) ProductRepository { return &productRepositoryImpl{db: db} }

func (r *productRepositoryImpl) Find(ctx context.Context, q dto.ProductQuery) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Product{}).Preload("Images")

	// filters
	if q.Name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(q.Name)+"%")
	}
	if q.Category != "" {
		query = query.Where("LOWER(category) = ?", strings.ToLower(q.Category))
	}
	if q.Search != "" {
		like := "%" + strings.ToLower(q.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", like, like)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// pagination guard
	page := q.Page
	if page <= 0 {
		page = 1
	}
	size := q.PageSize
	if size <= 0 || size > 100 {
		size = 10
	}
	offset := (page - 1) * size

	// sorting whitelist
	order := strings.ToLower(strings.TrimSpace(q.Sort))
	switch order {
	case "name_asc":
		query = query.Order("name ASC")
	case "name_desc":
		query = query.Order("name DESC")
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	default:
		query = query.Order("id DESC")
	}

	if err := query.Limit(size).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.WithContext(ctx).Preload("Images").First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
