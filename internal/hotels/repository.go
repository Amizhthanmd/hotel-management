package hotels

import (
	"fmt"
	"hotel-management/internal/models"
	"strings"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{db: db}
}

func (r *HotelRepository) GetBusinessByID(businessID string) (*models.Businesses, error) {
	var business models.Businesses
	err := r.db.Where("id = ?", businessID).First(&business).Error
	return &business, err
}

func (r *HotelRepository) CheckBusinessNameExists(name string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Businesses{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (r *HotelRepository) CheckHotelNameExists(schemaName string, name string) (bool, error) {
	var count int64
	err := r.db.Table(schemaName+".hotels").Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (r *HotelRepository) CheckRoomNumberExists(schemaName string, hotelID string, roomNumber string) (bool, error) {
	var count int64
	err := r.db.Table(schemaName+".rooms").Where("hotel_id = ? AND room_number = ?", hotelID, roomNumber).Count(&count).Error
	return count > 0, err
}

func (r *HotelRepository) CreateBusiness(name string) (*models.Businesses, error) {
	schemaName := "tenant_" + strings.ReplaceAll(strings.TrimSpace(strings.ToLower(name)), " ", "_")
	business := &models.Businesses{Name: name, SchemaName: schemaName}
	err := r.db.Create(business).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)).Error
	if err != nil {
		return nil, err
	}

	if err := r.db.Table(schemaName + ".hotels").AutoMigrate(&models.Hotel{}); err != nil {
		return nil, err
	}

	if err := r.db.Table(schemaName + ".rooms").AutoMigrate(&models.Room{}); err != nil {
		return nil, err
	}

	return business, nil
}

func (r *HotelRepository) CreateHotel(schemaName string, hotel *models.Hotel) error {
	return r.db.Table(schemaName + ".hotels").Create(hotel).Error
}

func (r *HotelRepository) CreateRoom(schemaName string, room *models.Room) error {
	return r.db.Table(schemaName + ".rooms").Create(room).Error
}

func UseSchema(schema string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Exec(fmt.Sprintf(`SET search_path = "%s"`, schema))
	}
}

func (r *HotelRepository) GetHotelsWithRooms(schemaName string) ([]models.Hotel, error) {
	var hotels []models.Hotel

	err := r.db.Scopes(UseSchema(schemaName)).
		Preload("Rooms").
		Find(&hotels).Error

	return hotels, err
}

func (r *HotelRepository) GetHotelByID(schemaName string, hotelID string) (*models.Hotel, error) {
	var hotel models.Hotel
	err := r.db.Table(schemaName+".hotels").Where("id = ?", hotelID).First(&hotel).Error
	return &hotel, err
}
