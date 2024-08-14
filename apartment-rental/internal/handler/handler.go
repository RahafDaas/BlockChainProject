package handler

import (
	"apartment-rental/internal/model/apartment"
	"apartment-rental/internal/repository/contract"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ApartmentHandler struct {
	contract.RepoContract
}

func (h *ApartmentHandler) GetApartment(ctx contractapi.TransactionContextInterface, apartmentID string) (*apartment.Apartment, error) {
	return h.RepoContract.GetApartment(ctx, apartmentID)
}

func (h *ApartmentHandler) GetAllApartments(ctx contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	return h.RepoContract.GetAllApartments(ctx)
}

// RegisterApartment registers a new apartment
func (h *ApartmentHandler) RegisterApartment(ctx contractapi.TransactionContextInterface, title string, roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear int, address string, hasWiFi, hasKitchen, hasLivingRoom bool, monthlyRent float64) (string, error) {
	apartmentID := generateDeterministicID(title, address, roomsCount, squareFootage, floorLevel)

	existingApt, _ := h.RepoContract.GetApartment(ctx, apartmentID)
	if existingApt != nil {
		return "", fmt.Errorf("apartment with ID %s already exists", apartmentID)
	}

	apt, err := apartment.New(apartmentID, title, address, roomsCount, squareFootage, restroomsCount, floorLevel, constructionYear, hasKitchen, hasLivingRoom, hasWiFi, monthlyRent)
	if err != nil {
		return "", err
	}

	err = h.RepoContract.PutApartment(ctx, apt)
	if err != nil {
		return "", err
	}

	return apartmentID, nil
}

func generateDeterministicID(title, address string, roomsCount, squareFootage, floorLevel int) string {
	data := fmt.Sprintf("%s|%s|%d|%d|%d", title, address, roomsCount, squareFootage, floorLevel)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
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

	apt, err := h.RepoContract.GetApartment(ctx, apartmentID)
	if err != nil {
		return fmt.Errorf("failed to get apartment: %v", err)
	}

	if apt.IsRented {
		return fmt.Errorf("apartment with ID %s is already rented", apartmentID)
	}

	clientID := generateDeterministicClientID(apartmentID, clientName, leaseStartDate)

	client := &apartment.Client{
		ClientID:       clientID,
		FullName:       clientName,
		ContactDetails: contactDetails,
		LeaseStartDate: leaseStartDate,
		LeaseEndDate:   leaseEndDate,
	}

	apt.IsRented = true
	apt.CurrentTenant = client

	if err := h.RepoContract.PutApartment(ctx, apt); err != nil {
		return fmt.Errorf("failed to update apartment rental status: %v", err)
	}

	return nil
}

func generateDeterministicClientID(apartmentID, clientName string, leaseStartDate time.Time) string {
	data := fmt.Sprintf("%s|%s|%s", apartmentID, clientName, leaseStartDate.Format(time.RFC3339))
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// ReturnApartment returns an apartment to the available list
func (h *ApartmentHandler) ReturnApartment(ctx contractapi.TransactionContextInterface, apartmentID string) error {
	apt, err := h.RepoContract.GetApartment(ctx, apartmentID)
	if err != nil {
		return fmt.Errorf("failed to get apartment: %v", err)
	}

	if !apt.IsRented {
		return fmt.Errorf("apartment with ID %s is not currently rented", apartmentID)
	}

	apt.IsRented = false
	apt.CurrentTenant = nil

	if err := h.RepoContract.PutApartment(ctx, apt); err != nil {
		return fmt.Errorf("failed to update apartment rental status: %v", err)
	}

	return nil
}
