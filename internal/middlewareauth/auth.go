package middlewareauth

import (
	"booking-api/internal/config"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, err := config.Load()
		if err != nil {
			http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		req := c.GetHeader("Authorization")
		if req == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"code":    http.StatusUnauthorized,
				"massage": "Authorization header is required",
			})
			return
		}
		tokenf := strings.TrimPrefix(req, "Bearer ")
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenf, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !tkn.Valid {
			http.Error(c.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()

	}
}
