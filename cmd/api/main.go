package api

import (
	"booking-api/internal/config"
	db2 "booking-api/internal/db"
	"booking-api/internal/domain/auth"
	"booking-api/internal/domain/model"
	"booking-api/internal/domain/resource"
	validator2 "booking-api/internal/domain/validator"
	"booking-api/internal/middleware"
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	var db *db2.Postgres
	var userRepo *model.Repository
	var resourceRepo *resource.Repository
	var resourceService *resource.Service
	var valid *validator2.Validatorr

	if cfg.DB.DatabaseURL == "" {
		db, err = db2.NewDB(cfg.DB)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}
	userRepo = model.NewUserRepo(db)
	jwtManager := auth.NewManager(cfg)
	authService := auth.NewService(*userRepo, *jwtManager)
	middleware := middleware.NewAuthMiddleware(jwtManager)
	resourceRepo = resource.New(db)
	resourceService = resource.NewService(resourceRepo)
	valid = validator2.New()
	authHandler := auth.NewHandler(authService, valid)
	resourceHandler := resource.NewHandler(resourceService)

	r := gin.Default()
	protected := r.Group("/")
	{
		protected.Use(middleware.Auth())
		protected.GET("/me", authHandler.Me)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
	resource := r.Group("/resource")
	{
		resource.GET("/resource", resourceHandler.Resource)
		resource.GET("/resource/:id", resourceHandler.Resource)
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Println("server started on", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		log.Println("server error:", err)
	case <-ctx.Done():
		log.Println("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("failed to shutdown server:", err)
		_ = srv.Close()
	}

	log.Println("server stopped")
}
