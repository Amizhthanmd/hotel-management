package models

import (
	"time"
)

type Businesses struct {
	ID         string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name       string    `json:"name" gorm:"not null"`
	SchemaName string    `json:"schema_name"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Businesses) TableName() string {
	return "public.businesses"
}

type Hotel struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rooms     []Room    `json:"rooms" gorm:"foreignKey:HotelID"`
}

type Room struct {
	ID         string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	HotelID    string    `json:"hotel_id" gorm:"type:uuid;not null"`
	RoomNumber string    `json:"room_number" gorm:"not null"`
	RoomType   string    `json:"room_type" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateBusinessRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateHotelRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type CreateRoomRequest struct {
	RoomNumber string `json:"room_number" binding:"required"`
	RoomType   string `json:"room_type" binding:"required"`
}
