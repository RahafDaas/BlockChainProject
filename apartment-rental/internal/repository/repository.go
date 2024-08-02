package repository

import (
	"apartment-rental/internal/model/apartment"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Repository interface {
	GetApartment(contractapi.TransactionContextInterface, string) (*apartment.Apartment, error)
	GetAllApartments(contractapi.TransactionContextInterface) ([]*apartment.Apartment, error)
	PutApartment(contractapi.TransactionContextInterface, *apartment.Apartment) error
}
