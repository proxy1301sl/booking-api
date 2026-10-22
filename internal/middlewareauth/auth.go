package middlewareauth

import (
	"booking-api/internal/domain/auth"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	manager *auth.Manager
}

func NewAuthMiddleware(manager *auth.Manager) *AuthMiddleware {
	return &AuthMiddleware{
		manager: manager,
	}
}

func (m *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"code":    http.StatusUnauthorized,
				"massage": "Authorization header is required",
			})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		pars, err := m.manager.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		c.Set("user_id", pars.UserID)
		c.Next()

	}
}
