package main

import (
	"apartment-rental/internal/handler"
	"apartment-rental/internal/repository/memory"
	"fmt"
)

/*
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
*/
func main() {

	//mainChaincode()
	handler := handler.ApartmentHandler{
		RepoContract: *memory.New(),
	}
	err := handler.Register(nil, "101", "Aljoharah", 3, 500, 30.000, true, true, 2, 1, 2015, "Alolaya", true)
	if err != nil {
		panic(err)
	}

	a, err := handler.GetApartment(nil, "101")
	handler.RentApartment(nil, "101")
	fmt.Printf("%#v", a)

}
