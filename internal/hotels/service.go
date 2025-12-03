package hotels

import (
	"hotel-management/internal/models"
)

type HotelService struct {
	repository *HotelRepository
}

func InitializeHotelService(repository *HotelRepository) *HotelService {
	return &HotelService{
		repository: repository,
	}
}

func (s *HotelService) CreateBusiness(name string) (*models.Businesses, error) {
	return s.repository.CreateBusiness(name)
}

func (s *HotelService) CreateHotel(schemaName string, req models.CreateHotelRequest) (*models.Hotel, error) {
	hotel := &models.Hotel{Name: req.Name, Address: req.Address}
	err := s.repository.CreateHotel(schemaName, hotel)
	return hotel, err
}

func (s *HotelService) CreateRoom(schemaName string, hotelID string, req models.CreateRoomRequest) (*models.Room, error) {
	room := &models.Room{HotelID: hotelID, RoomNumber: req.RoomNumber, RoomType: req.RoomType}
	err := s.repository.CreateRoom(schemaName, room)
	return room, err
}

func (s *HotelService) GetHotelsWithRooms(schemaName string) ([]models.Hotel, error) {
	return s.repository.GetHotelsWithRooms(schemaName)
}

func (s *HotelService) GetBusinessByID(businessID string) (*models.Businesses, error) {
	return s.repository.GetBusinessByID(businessID)
}

func (s *HotelService) CheckBusinessNameExists(name string) (bool, error) {
	return s.repository.CheckBusinessNameExists(name)
}

func (s *HotelService) CheckHotelNameExists(schemaName string, name string) (bool, error) {
	return s.repository.CheckHotelNameExists(schemaName, name)
}

func (s *HotelService) CheckRoomNumberExists(schemaName string, hotelID string, roomNumber string) (bool, error) {
	return s.repository.CheckRoomNumberExists(schemaName, hotelID, roomNumber)
}

func (s *HotelService) GetHotelByID(schemaName string, hotelID string) (*models.Hotel, error) {
	return s.repository.GetHotelByID(schemaName, hotelID)
}
