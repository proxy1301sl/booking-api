package api

import (
	"booking-api/internal/config"
	"booking-api/internal/db"
	"context"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	db, err := db.Connect(cfg.DB, ctx)
	if err != nil {
		panic(err)
	}




}
