package hotels

import (
	"fmt"
	"os"

	"hotel-management/internal/tenant"

	"github.com/gin-gonic/gin"
)

func StartRouter(handler *Handlers) {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Route not found"})
		c.Abort()
	})
	router.POST("/api/v1/business", handler.CreateBusiness)

	hotelApi := router.Group("/api/v1")
	hotelApi.Use(tenant.Middleware(handler.TenantDB))
	{
		hotelApi.POST("/hotels", handler.CreateHotels)
		hotelApi.POST("/hotels/:hotel_id/rooms", handler.CreateHotelRooms)
		hotelApi.GET("/hotels", handler.GetHotelsWithRooms)
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
		port = ":8000"
	}
	router.Run(port)
}
