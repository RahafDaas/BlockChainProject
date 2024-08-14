package main

import (
	"apartment-rental/internal/handler"
	//"apartment-rental/internal/repository/memory"
	"fmt"
	//"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func mainChaincode() {
	chaincode, err := contractapi.NewChaincode(new(handler.ApartmentHandler))

	if err != nil {
		fmt.Printf("error creating Apartment chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("error starting Apartment chaincode: %v", err)
	}
}

func main() {

	mainChaincode()

	/*handler := handler.ApartmentHandler{
		RepoContract: memory.New(), // Pass the pointer directly
	}

	var ctx contractapi.TransactionContextInterface

	// Register a new apartment
	err := handler.RegisterApartment(ctx, "Modern Apartment", 3, 1200, 2, 5, 2019, "123 Main St", true, true, true, 2500.0)
	if err != nil {
		panic(err)
	}

	// For demonstration purposes, retrieve and print the registered apartments
	apts, err := handler.GetAllApartments(ctx)
	if err != nil {
		panic(err)
	}

	for _, apt := range apts {
		fmt.Printf("Apartment ID: %s, Title: %s, Address: %s, Rent: %.2f\n", apt.ApartmentID, apt.Title, apt.Address, apt.MonthlyRent)
	}

	// Rent out the first apartment
	if len(apts) > 0 {
		leaseStartDate := time.Now()
		leaseEndDate := leaseStartDate.AddDate(1, 0, 0) // 1-year lease
		err = handler.RentApartment(ctx, apts[0].ApartmentID, "John Doe", "123-456-7890", leaseStartDate, leaseEndDate)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Apartment rented to John Doe for Apartment ID: %s\n", apts[0].ApartmentID)
	}*/
}
