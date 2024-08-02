package handler

import (
	"apartment-rental/internal/model/apartment"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ApartmentHandler struct {
}

func (h *ApartmentHandler) GetApartment(ctx contractapi.TransactionContextInterface, ID string) (*apartment.Apartment, error) {
	return
}

func (h *ApartmentHandler) GetAllApartments(ctx contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	return
}

func (h *ApartmentHandler) Register(ctx contractapi.TransactionContextInterface, ID int, Name string, StreetName string, DistrictName string, Size int, RentalPrice int, Color string) error {
	// Check if the student is Registered
	return
}

func (h *ApartmentHandler) UpdateName(ctx contractapi.TransactionContextInterface, ID, newName string) error {
	// check if the student exists
	return
}

func (h *ApartmentHandler) UpdateStreetName(ctx contractapi.TransactionContextInterface, ID, newStreetName string) error {
	// check if the student exists
	return
}

func (h *ApartmentHandler) ChangeColor(ctx contractapi.TransactionContextInterface, ID, newColor string) error {
	// check if the student exists
	return
}

func (h *ApartmentHandler) UpdateDistrictName(ctx contractapi.TransactionContextInterface, ID string, newDistrictName string) error {
	return
}
func (h *ApartmentHandler) UpdateSize(ctx contractapi.TransactionContextInterface, ID string, newSize int) error {
	return
}
func (h *ApartmentHandler) UpdatePrice(ctx contractapi.TransactionContextInterface, ID string, newPrice float64) error {
	return
}
func (h *ApartmentHandler) UpdateID(ctx contractapi.TransactionContextInterface, ID string, newID string) error {
	return
}
