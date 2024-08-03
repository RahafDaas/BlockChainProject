package handler

import (
	"apartment-rental/internal/model/apartment"
	"apartment-rental/internal/repository/contract"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ApartmentHandler struct {
	// RepoContract memory.RepoMemory
	contract.RepoContract
}

// To get specific apartment
func (h *ApartmentHandler) GetApartment(ctx contractapi.TransactionContextInterface, ID string) (*apartment.Apartment, error) {
	return h.RepoContract.GetApartment(ctx, ID)
}

// To get all apartments
func (h *ApartmentHandler) GetAllApartments(ctx contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	return h.RepoContract.GetAllApartments(ctx)
}

// To register new apartment
func (h *ApartmentHandler) Register(ctx contractapi.TransactionContextInterface, ID string, Name string, StreetName string, DistrictName string, Size int, RentalPrice float64, Color string) error {
	// Check if the aparment is Registered
	_, err := h.GetApartment(ctx, ID)
	if err == nil {
		return fmt.Errorf("apartment with ID %s is registered", ID)
	}
	apartment := apartment.New(ID, Name, StreetName, DistrictName, Size, RentalPrice, Color)
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the name of an apartment
func (h *ApartmentHandler) UpdateName(ctx contractapi.TransactionContextInterface, ID, newName string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Name = newName
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the street name of an apartment
func (h *ApartmentHandler) UpdateStreetName(ctx contractapi.TransactionContextInterface, ID, newStreetName string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.StreetName = newStreetName
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the name of an apartment
func (h *ApartmentHandler) ChangeColor(ctx contractapi.TransactionContextInterface, ID, newColor string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Color = newColor
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the district name of an apartment
func (h *ApartmentHandler) UpdateDistrictName(ctx contractapi.TransactionContextInterface, ID string, newDistrictName string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.DistrictName = newDistrictName
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the size of an apartment
func (h *ApartmentHandler) UpdateSize(ctx contractapi.TransactionContextInterface, ID string, newSize int) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Size = newSize
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the price of an apartment
func (h *ApartmentHandler) Price(ctx contractapi.TransactionContextInterface, ID string, Price float64) error {
	if Price < 0 {
		return fmt.Errorf("%.2f must be positive", Price)
	}
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.RentalPrice = Price
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the ID of an apartment
func (h *ApartmentHandler) UpdateID(ctx contractapi.TransactionContextInterface, ID string, newID string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.ID = newID
	return h.RepoContract.PutApartment(ctx, apartment)
}
