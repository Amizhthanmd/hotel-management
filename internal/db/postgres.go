package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeTenantDB() *gorm.DB {
	dsn := fmt.Sprintf("%s%s", os.Getenv("POSTGRES_DB_URL"), os.Getenv("DB_NAME"))
	tenantDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error initializing tenant database")
	}
	return tenantDB
}
