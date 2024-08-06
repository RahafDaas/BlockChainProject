package contract

import (
	"apartment-rental/internal/model/apartment"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoContract struct {
	contractapi.Contract
}

func (r *RepoContract) GetApartment(ctx contractapi.TransactionContextInterface, ID string) (*apartment.Apartment, error) {
	apartmentJSON, err := ctx.GetStub().GetState(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	if apartmentJSON == nil {
		return nil, fmt.Errorf("Apartment %s does not exist", ID)
	}

	var apartment apartment.Apartment
	err = json.Unmarshal(apartmentJSON, &apartment)
	if err != nil {
		return nil, err
	}

	return &apartment, nil
}

func (r *RepoContract) GetAllApartments(ctx contractapi.TransactionContextInterface) ([]*apartment.Apartment, error) {
	apartmentIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read form world state: error %v", err)
	}
	defer apartmentIterator.Close()

	var Apartments []*apartment.Apartment

	for apartmentIterator.HasNext() {
		result, err := apartmentIterator.Next()
		if err != nil {
			return nil, err
		}
		var apartment apartment.Apartment
		err = json.Unmarshal(result.Value, &apartment)
		if err != nil {
			return nil, err
		}

		Apartments = append(Apartments, &apartment)
	}

	return Apartments, nil
}

func (r *RepoContract) PutApartment(ctx contractapi.TransactionContextInterface, apartment *apartment.Apartment) error {
	apartmentJSON, err := json.Marshal(apartment)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(apartment.ID, apartmentJSON)
}
