package memory

import (
	"apartment-rental/internal/model/apartment"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoMemory map[string]*apartment.Apartment

func New() *RepoMemory {
	rm := make(RepoMemory)
	return &rm
}

func (rm *RepoMemory) GetApartment(_ contractapi.TransactionContextInterface, ID string) (*apartment.Apartment, error) {
	apartment, existing := (*rm)[ID]
	if !existing {
		return nil, fmt.Errorf("no apartment with ID: %s", ID)
	}
	return apartment, nil
}

func (rm *RepoMemory) GetAllApartments(_ contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	var Apartments []*apartment.Apartment

	for _, v := range *rm {
		Apartments = append(Apartments, v)
	}

	if len(Apartments) == 0 {
		return nil, fmt.Errorf("no apartments found")
	}

	return Apartments, nil
}

// used to add new apartment and to update
func (rm *RepoMemory) PutApartment(_ contractapi.TransactionContextInterface, apartment *apartment.Apartment) error {
	(*rm)[apartment.ApartmentID] = apartment
	return nil
}
