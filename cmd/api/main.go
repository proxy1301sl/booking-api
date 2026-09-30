package api

import (
	"booking-api/internal/config"
	db2 "booking-api/internal/db"
	"booking-api/internal/domain/auth"
	"booking-api/internal/domain/user"
	"context"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	db, err := db2.NewDB(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := user.New(db)
	manager := auth.NewManager(cfg)
	authService := auth.NewService(userRepo, manager)
	authHandler := auth.NewHandler(authService)

}
