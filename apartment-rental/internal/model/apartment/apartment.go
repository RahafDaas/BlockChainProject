package apartment

import (
	"errors"
	"time"
)

// Client represents a tenant in an apartment
type Client struct {
	ClientID       string    `json:"client_id"`        // Unique identifier for the client
	FullName       string    `json:"full_name"`        // Full name of the client
	TotalPayments  float64   `json:"total_payments"`   // Total amount of payments made by the client
	ContactDetails string    `json:"contact_details"`  // Contact information for the client (email or phone)
	LeaseStartDate time.Time `json:"lease_start_date"` // Start date of the lease agreement
	LeaseEndDate   time.Time `json:"lease_end_date"`   // End date of the lease agreement
}

// Apartment represents a rental apartment
type Apartment struct {
	ApartmentID      string  `json:"apartment_id"`             // Unique identifier for the apartment
	Title            string  `json:"title"`                    // Descriptive title of the apartment
	RoomsCount       int     `json:"rooms_count"`              // Number of rooms in the apartment
	SquareFootage    int     `json:"square_footage"`           // Size of the apartment in square feet or meters
	MonthlyRent      float64 `json:"monthly_rent"`             // Rental price per month
	HasKitchen       bool    `json:"has_kitchen"`              // Indicates if the apartment has a kitchen
	HasLivingRoom    bool    `json:"has_living_room"`          // Indicates if the apartment has a living room
	RestroomsCount   int     `json:"restrooms_count"`          // Number of restrooms in the apartment
	FloorLevel       int     `json:"floor_level"`              // Floor number where the apartment is located
	ConstructionYear int     `json:"construction_year"`        // Year the building was constructed
	Address          string  `json:"address"`                  // Physical address of the apartment
	HasWiFi          bool    `json:"has_wifi"`                 // Indicates if the apartment has WiFi access
	IsRented         bool    `json:"is_rented"`                // Rental status of the apartment
	CurrentTenant    *Client `json:"current_tenant,omitempty"` // Details of the current tenant (if rented)
}

// New creates a new Apartment instance
func New(apartmentID, title, address string, roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear int, hasKitchen, hasLivingRoom, hasWiFi bool, monthlyRent float64) (*Apartment, error) {
	if apartmentID == "" || title == "" || address == "" {
		return nil, errors.New("apartmentID, title, and address cannot be empty")
	}
	if roomsCount < 0 || squareFootage < 0 || restroomsCount < 0 || floorLevel < 0 || constructionYear < 0 || monthlyRent < 0 {
		return nil, errors.New("roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear, and monthlyRent must be non-negative")
	}
	return &Apartment{
		ApartmentID:      apartmentID,
		Title:            title,
		RoomsCount:       roomsCount,
		SquareFootage:    squareFootage,
		HasKitchen:       hasKitchen,
		HasLivingRoom:    hasLivingRoom,
		RestroomsCount:   restroomsCount,
		FloorLevel:       floorLevel,
		ConstructionYear: constructionYear,
		Address:          address,
		HasWiFi:          hasWiFi,
		MonthlyRent:      monthlyRent,
		IsRented:         false,
		CurrentTenant:    nil,
	}, nil
}
