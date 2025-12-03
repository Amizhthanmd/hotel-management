package main

import (
	log "hotel-management/internal"
	"hotel-management/internal/db"
	"hotel-management/internal/hotels"
	"hotel-management/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	logger := log.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	tenantDB := db.InitializeTenantDB()
	if !tenantDB.Migrator().HasTable(&models.Businesses{}) {
		tenantDB.AutoMigrate(&models.Businesses{})
	}
	repository := hotels.NewHotelRepository(tenantDB)
	hotelService := hotels.InitializeHotelService(repository)

	handler := hotels.InitializeHandler(logger, tenantDB, hotelService)
	hotels.StartRouter(handler)
}
