package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	httpadapter "github.com/FatwahFir/xpanca-be/internal/adapter/http"
	mysqlrepo "github.com/FatwahFir/xpanca-be/internal/adapter/repository/mysql"
	"github.com/FatwahFir/xpanca-be/internal/config"
	"github.com/FatwahFir/xpanca-be/internal/middleware"
	"github.com/FatwahFir/xpanca-be/internal/usecase"
	"github.com/FatwahFir/xpanca-be/pkg/jwtx"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	jwtm := jwtx.New(cfg.JWTSecret, cfg.JWTTTLMinutes)

	v1 := r.Group("api/v1")

	//REPO
	userRepo := mysqlrepo.NewUserRepo(db)
	productRepo := mysqlrepo.NewProductRepo(db)

	//USECASE
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtm)
	ProductUsecase := usecase.NewProductUsecase(productRepo)

	httpadapter.NewAuthHandler(v1, authUsecase)

	v1.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now()}) })

	authGroup := v1.Group("/")

	authGroup.Use(middleware.JWTAuth(jwtm))
	httpadapter.NewProductHandler(authGroup, ProductUsecase)

	log.Printf("listening on :%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
