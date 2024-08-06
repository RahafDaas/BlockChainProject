package handler

import (
	"apartment-rental/internal/model/apartment"
	"apartment-rental/internal/repository/memory"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ApartmentHandler struct {
	RepoContract memory.RepoMemory
	//contract.RepoContract
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
func (h *ApartmentHandler) Register(ctx contractapi.TransactionContextInterface, ID string, Name string, NumberOfRooms int, Size int, rentalPrice float64, Kitchen bool, LivingRoom bool, RestroomNumber int, FloorNumber int, BuildYear int, Location string, WiFi bool) error {
	// Check if the aparment is Registered
	_, err := h.GetApartment(ctx, ID)
	if err == nil {
		return fmt.Errorf("apartment with ID %s is registered", ID)
	}
	apartment := apartment.New(ID, Name, NumberOfRooms, Size, Kitchen, LivingRoom, RestroomNumber, FloorNumber, BuildYear, Location, WiFi, rentalPrice)
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
func (h *ApartmentHandler) UpdateNumberOfRooms(ctx contractapi.TransactionContextInterface, ID string, newNumberOfRooms int) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.NumberOfRooms = newNumberOfRooms
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the name of an apartment
func (h *ApartmentHandler) ChangeKitchen(ctx contractapi.TransactionContextInterface, ID string, newKitchen bool) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Kitchen = newKitchen
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the district name of an apartment
func (h *ApartmentHandler) UpdateLivingRoom(ctx contractapi.TransactionContextInterface, ID string, newLivingRoom bool) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.LivingRoom = newLivingRoom
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

// To update the Restroom Number of an apartment
func (h *ApartmentHandler) UpdateRestroomNumber(ctx contractapi.TransactionContextInterface, ID string, newRestroomNumber int) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.RestroomNumber = newRestroomNumber
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the Floor Number of an apartment
func (h *ApartmentHandler) UpdateFloorNumber(ctx contractapi.TransactionContextInterface, ID string, newFloorNumber int) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.FloorNumber = newFloorNumber
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the BuildYear of an apartment
func (h *ApartmentHandler) UpdateBuildYear(ctx contractapi.TransactionContextInterface, ID string, newBuildYear int) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.BuildYear = newBuildYear
	return h.RepoContract.PutApartment(ctx, apartment)
}

// To update the location of an apartment
func (h *ApartmentHandler) Updatelocation(ctx contractapi.TransactionContextInterface, ID, newlocation string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Location = newlocation
	return h.RepoContract.PutApartment(ctx, apartment)
}

// Rent
func (h *ApartmentHandler) RentApartment(ctx contractapi.TransactionContextInterface, ID string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Rented = true
	fmt.Printf("Rented with %v", apartment.RentalPrice)
	return h.RepoContract.PutApartment(ctx, apartment)

}

// UnRent
func (h *ApartmentHandler) UnRentApartment(ctx contractapi.TransactionContextInterface, ID string) error {
	// check if the apartment exists
	apartment, err := h.GetApartment(ctx, ID)
	if err != nil {
		return err
	}
	apartment.Rented = false
	//fmt.Printf("Rental price is %v", apartment.RentalPrice)
	return h.RepoContract.PutApartment(ctx, apartment)

}
