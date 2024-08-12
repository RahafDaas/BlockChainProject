package handler

import (
	"apartment-rental/internal/model/apartment"
	"apartment-rental/internal/repository/contract"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ApartmentHandler struct {
	//RepoContract memory.RepoMemory
	contract.RepoContract
}

// GetApartment retrieves a specific apartment by ID
func (h *ApartmentHandler) GetApartment(ctx contractapi.TransactionContextInterface, apartmentID string) (*apartment.Apartment, error) {
	return h.RepoContract.GetApartment(ctx, apartmentID)
}

// GetAllApartments retrieves all apartments
func (h *ApartmentHandler) GetAllApartments(ctx contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	return h.RepoContract.GetAllApartments(ctx)
}

// RegisterApartment registers a new apartment
func (h *ApartmentHandler) RegisterApartment(ctx contractapi.TransactionContextInterface, title string, roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear int, address string, hasWiFi, hasKitchen, hasLivingRoom bool, monthlyRent float64) error {
	// Generate a new UUID for apartmentID
	apartmentID := uuid.New().String()

	// Create a new apartment instance
	apt, err := apartment.New(apartmentID, title, address, roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear, hasKitchen, hasLivingRoom, hasWiFi, monthlyRent)
	if err != nil {
		return err
	}
	fmt.Print(apartmentID)
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentTitle updates the title of an apartment
func (h *ApartmentHandler) UpdateApartmentTitle(ctx contractapi.TransactionContextInterface, apartmentID, newTitle string) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.Title = newTitle
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentRooms updates the number of rooms in an apartment
func (h *ApartmentHandler) UpdateApartmentRooms(ctx contractapi.TransactionContextInterface, apartmentID string, newRoomsCount int) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.RoomsCount = newRoomsCount
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentKitchen updates the kitchen status of an apartment
func (h *ApartmentHandler) UpdateApartmentKitchen(ctx contractapi.TransactionContextInterface, apartmentID string, hasKitchen bool) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.HasKitchen = hasKitchen
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentLivingRoom updates the living room status of an apartment
func (h *ApartmentHandler) UpdateApartmentLivingRoom(ctx contractapi.TransactionContextInterface, apartmentID string, hasLivingRoom bool) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.HasLivingRoom = hasLivingRoom
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentSize updates the square footage of an apartment
func (h *ApartmentHandler) UpdateApartmentSize(ctx contractapi.TransactionContextInterface, apartmentID string, newSquareFootage int) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.SquareFootage = newSquareFootage
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentRent updates the rental price of an apartment
func (h *ApartmentHandler) UpdateApartmentRent(ctx contractapi.TransactionContextInterface, apartmentID string, newRent float64) error {
	if newRent < 0 {
		return fmt.Errorf("rental price must be positive")
	}
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.MonthlyRent = newRent
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentID updates the ID of an apartment
func (h *ApartmentHandler) UpdateApartmentID(ctx contractapi.TransactionContextInterface, apartmentID, newID string) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.ApartmentID = newID
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentRestrooms updates the number of restrooms in an apartment
func (h *ApartmentHandler) UpdateApartmentRestrooms(ctx contractapi.TransactionContextInterface, apartmentID string, newRestroomsCount int) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.RestroomsCount = newRestroomsCount
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentFloor updates the floor level of an apartment
func (h *ApartmentHandler) UpdateApartmentFloor(ctx contractapi.TransactionContextInterface, apartmentID string, newFloorLevel int) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.FloorLevel = newFloorLevel
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentConstructionYear updates the construction year of an apartment
func (h *ApartmentHandler) UpdateApartmentConstructionYear(ctx contractapi.TransactionContextInterface, apartmentID string, newConstructionYear int) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.ConstructionYear = newConstructionYear
	return h.RepoContract.PutApartment(ctx, apt)
}

// UpdateApartmentAddress updates the address of an apartment
func (h *ApartmentHandler) UpdateApartmentAddress(ctx contractapi.TransactionContextInterface, apartmentID, newAddress string) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	apt.Address = newAddress
	return h.RepoContract.PutApartment(ctx, apt)
}

// RentApartment rents out an apartment to a client
func (h *ApartmentHandler) RentApartment(ctx contractapi.TransactionContextInterface, apartmentID, clientName, contactDetails string, leaseStartDate, leaseEndDate time.Time) error {
	if leaseEndDate.Before(leaseStartDate) {
		return fmt.Errorf("lease end date must be after the lease start date")
	}

	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	if apt.IsRented {
		return fmt.Errorf("apartment is already rented")
	}

	// Generate a new UUID for clientID
	clientID := uuid.New().String()

	client := &apartment.Client{
		ClientID:       clientID,
		FullName:       clientName,
		ContactDetails: contactDetails,
		LeaseStartDate: leaseStartDate,
		LeaseEndDate:   leaseEndDate,
	}

	apt.IsRented = true
	apt.CurrentTenant = client

	return h.RepoContract.PutApartment(ctx, apt)
}

// UnrentApartment marks an apartment as unrented
func (h *ApartmentHandler) UnrentApartment(ctx contractapi.TransactionContextInterface, apartmentID string) error {
	apt, err := h.GetApartment(ctx, apartmentID)
	if err != nil {
		return err
	}
	if !apt.IsRented {
		return fmt.Errorf("apartment is not currently rented")
	}

	apt.IsRented = false
	apt.CurrentTenant = nil

	return h.RepoContract.PutApartment(ctx, apt)
}
