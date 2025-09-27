package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

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

	// seed users
	users := []struct {
		Username string
		Password string
		Role     string
	}{
		{"cole", "password", "user"},
		{"brian", "password", "user"},
		{"alice", "password", "user"},
	}
	for _, u := range users {
		hash, _ := password.Hash(u.Password)
		user := domain.User{Username: u.Username, Password: hash, Role: u.Role}
		db.Where("username = ?", user.Username).FirstOrCreate(&user)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	categories := []string{"men", "women", "bag", "electronics", "accessories"}

	// seed products
	for i := 1; i <= 100; i++ {
		name := fmt.Sprintf("Product %03d", i)
		category := categories[r.Intn(len(categories))]
		price := int64(100000 + r.Intn(900000))
		desc := fmt.Sprintf("Description for %s in category %s", name, category)

		p := domain.Product{
			Name:        name,
			Category:    category,
			Price:       price,
			Description: desc,
		}
		db.Where("name = ?", p.Name).FirstOrCreate(&p)

		// Thumbnail + 2 detail images
		thumb := domain.Image{
			ProductID:   p.ID,
			URL:         fmt.Sprintf("https://picsum.photos/seed/thumb-%d/600/400", i),
			IsThumbnail: true,
		}
		db.Where("product_id = ? AND url = ?", thumb.ProductID, thumb.URL).
			FirstOrCreate(&thumb)

		for j := 1; j <= 2; j++ {
			img := domain.Image{
				ProductID:   p.ID,
				URL:         fmt.Sprintf("https://picsum.photos/seed/detail-%d-%d/600/400", i, j),
				IsThumbnail: false,
			}
			db.Where("product_id = ? AND url = ?", img.ProductID, img.URL).
				FirstOrCreate(&img)
		}
	}

	// summary
	var userCount, productCount int64
	db.Model(&domain.User{}).Count(&userCount)
	db.Model(&domain.Product{}).Count(&productCount)

	out, _ := json.MarshalIndent(map[string]any{
		"users":    userCount,
		"products": productCount,
	}, "", "  ")
	fmt.Println(string(out))
}
