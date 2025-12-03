package tenant

import (
	"context"
	"net/http"

	"hotel-management/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type contextKey string

const BusinessIDKey contextKey = "business_id"

func Middleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		businessID := c.GetHeader("X-Business-ID")
		if businessID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "business id required"})
			c.Abort()
			return
		}

		var business models.Businesses
		if err := db.Where("id = ?", businessID).First(&business).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid business id"})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), BusinessIDKey, businessID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
