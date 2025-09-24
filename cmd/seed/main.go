package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/FatwahFir/xpanca-be/internal/config"
	"github.com/FatwahFir/xpanca-be/internal/domain"
	"github.com/FatwahFir/xpanca-be/pkg/password"
)

func main() {
	cfg := config.Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	// seed user
	hash, _ := password.Hash("secret123")
	admin := domain.User{Username: "admin", Password: hash, Role: "admin"}
	db.Where("username = ?", admin.Username).FirstOrCreate(&admin)

	// seed products
	products := []domain.Product{
		{Name: "Sneaker Alpha", Category: "men", Price: 799000, Description: "Lightweight men sneaker"},
		{Name: "Sneaker Beta", Category: "women", Price: 899000, Description: "Comfortable women sneaker"},
		{Name: "Backpack Pro", Category: "bag", Price: 499000, Description: "Durable backpack"},
	}
	imageData := [][]domain.Image{
		{{URL: "https://picsum.photos/seed/alpha/600/400", IsThumbnail: true}, {URL: "https://picsum.photos/seed/alpha2/600/400"}},
		{{URL: "https://picsum.photos/seed/beta/600/400", IsThumbnail: true}, {URL: "https://picsum.photos/seed/beta2/600/400"}},
		{{URL: "https://picsum.photos/seed/bag/600/400", IsThumbnail: true}, {URL: "https://picsum.photos/seed/bag2/600/400"}},
	}

	for i := range products {
		db.Where("name = ?", products[i].Name).FirstOrCreate(&products[i])
		for _, img := range imageData[i] {
			img.ProductID = products[i].ID
			db.Where("product_id=? AND url=?", img.ProductID, img.URL).FirstOrCreate(&domain.Image{ProductID: img.ProductID, URL: img.URL, IsThumbnail: img.IsThumbnail})
		}
	}

	var count int64
	db.Model(&domain.Product{}).Count(&count)
	out, _ := json.MarshalIndent(map[string]any{"users": 1, "products": count}, "", "  ")
	fmt.Println(string(out))
}
