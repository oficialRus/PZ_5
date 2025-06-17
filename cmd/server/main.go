// cmd/server/main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"CADET_PROJECT/internal/domain"
	"CADET_PROJECT/internal/repository"
	"CADET_PROJECT/internal/service"
	httpHandlers "CADET_PROJECT/internal/transport/http"
)

func main() {
	// 1. Открываем (или создаём) файл-БД SQLite
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("db open: %v", err)
	}

	// 2. Авто-миграции двух таблиц
	if err := db.AutoMigrate(&domain.User{}, &domain.Product{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// 3. Wire — репозитории → сервисы → HTTP-хендлеры
	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)

	userSvc := service.NewUserService(userRepo)
	productSvc := service.NewProductService(productRepo)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	httpHandlers.RegisterUserRoutes(r, userSvc)
	httpHandlers.RegisterProductRoutes(r, productSvc)

	// 4. Запускаем сервер
	addr := ":8080"
	log.Printf("⇨ start server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("gin: %v", err)
	}
}
