package hotels

import (
	"net/http"

	"hotel-management/internal/models"
	"hotel-management/internal/tenant"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Handlers struct {
	Logger       *log.Logger
	TenantDB     *gorm.DB
	HotelService *HotelService
}

func InitializeHandler(
	Logger *log.Logger,
	TenantDB *gorm.DB,
	HotelService *HotelService,
) *Handlers {
	return &Handlers{
		Logger:       Logger,
		TenantDB:     TenantDB,
		HotelService: HotelService,
	}
}

func (h *Handlers) CreateBusiness(ctx *gin.Context) {
	var req models.CreateBusinessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error in bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	exists, err := h.HotelService.CheckBusinessNameExists(req.Name)
	if err != nil {
		h.Logger.Error("Error checking business name existence")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to check business name",
		})
		return
	}
	if exists {
		h.Logger.Error("Business name already exists")
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "Business name already exists",
		})
		return
	}

	business, err := h.HotelService.CreateBusiness(req.Name)
	if err != nil {
		h.Logger.Error("Failed to create business")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create business",
		})
		return
	}

	h.Logger.Info("Business created successfully")
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"data":    business,
		"message": "Business created successfully",
	})
}

func (h *Handlers) CreateHotels(ctx *gin.Context) {
	var req models.CreateHotelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error in bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	businessID := ctx.Request.Context().Value(tenant.BusinessIDKey).(string)
	business, err := h.HotelService.GetBusinessByID(businessID)
	if err != nil {
		h.Logger.Error("Error in fetching business")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid business",
		})
		return
	}

	exists, err := h.HotelService.CheckHotelNameExists(business.SchemaName, req.Name)
	if err != nil {
		h.Logger.Error("Error checking hotel name existence")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to check hotel name",
		})
		return
	}
	if exists {
		h.Logger.Error("Hotel name already exists")
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "Hotel name already exists",
		})
		return
	}

	hotel, err := h.HotelService.CreateHotel(business.SchemaName, req)
	if err != nil {
		h.Logger.Error("failed to create hotel")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "failed to create hotel",
		})
		return
	}

	h.Logger.Info("Hotel created successfully")
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Hotel created successfully",
		"data":    hotel,
	})
}

func (h *Handlers) CreateHotelRooms(ctx *gin.Context) {
	var req models.CreateRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error in bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request",
		})
		return
	}

	hotelID := ctx.Param("hotel_id")
	if hotelID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Hotel id required",
		})
		return
	}

	businessID := ctx.Request.Context().Value(tenant.BusinessIDKey).(string)
	business, err := h.HotelService.GetBusinessByID(businessID)
	if err != nil {
		h.Logger.Error("Error in fetching business")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to fetch business",
		})
		return
	}

	// Validate that the hotel belongs to the business
	_, err = h.HotelService.GetHotelByID(business.SchemaName, hotelID)
	if err != nil {
		h.Logger.Error("Hotel not found or does not belong to the business")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid hotel ID for this business",
		})
		return
	}

	exists, err := h.HotelService.CheckRoomNumberExists(business.SchemaName, hotelID, req.RoomNumber)
	if err != nil {
		h.Logger.Error("Error checking room number")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to check room number",
		})
		return
	}
	if exists {
		h.Logger.Error("Room number already exists")
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "Room number already exists",
		})
		return
	}

	room, err := h.HotelService.CreateRoom(business.SchemaName, hotelID, req)
	if err != nil {
		h.Logger.Error("Error in create rooms")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create rooms",
		})
		return
	}

	h.Logger.Info("Room created successfully")
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Room created successfully",
		"data":    room,
	})
}

func (h *Handlers) GetHotelsWithRooms(ctx *gin.Context) {
	businessID := ctx.Request.Context().Value(tenant.BusinessIDKey).(string)
	business, err := h.HotelService.GetBusinessByID(businessID)
	if err != nil {
		h.Logger.Error("Error in fetching business")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Failed to fetch business",
		})
		return
	}

	hotels, err := h.HotelService.GetHotelsWithRooms(business.SchemaName)
	if err != nil {
		h.Logger.Error("Error in fetching hotels and rooms")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to get hotels and rooms",
		})
		return
	}

	h.Logger.Info("Hotels and rooms fetched successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Hotels and rooms fetched successfully",
		"data":    hotels,
	})
}
