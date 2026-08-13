package middleware

import (
	"net/http"
	"os"
	"strings"

	"backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "ไม่พบ token",
			})
			return
		}

		// Authorization: Bearer xxxxx
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "รูปแบบ token ไม่ถูกต้อง",
			})
			return
		}

		tokenString := parts[1]

		secret := os.Getenv("JWT_SECRET")

		if secret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "JWT_SECRET is not configured",
			})
			return
		}

		claims := &utils.UserClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			},
			jwt.WithValidMethods([]string{
				jwt.SigningMethodHS256.Alg(),
			}),
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token ไม่ถูกต้องหรือหมดอายุ",
			})
			return
		}

		// เก็บ user id ไว้ให้ handler ตัวถัดไปใช้
		c.Set("userID", claims.UserID)

		c.Next()
	}
}