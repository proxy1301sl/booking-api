package router

import (
	"booking-api/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *auth.Handler) *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/auth/register", h.Register)
		users.POST("/auth/login", h.Login)
		users.GET("/users/me", h.Me)
	}
	return r

}
