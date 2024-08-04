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
	err := handler.Register(nil, "101", "Aljoharah", "Banban", "Alolaya", 300, 40.000, "Brown")
	if err != nil {
		panic(err)
	}

	a, err := handler.GetApartment(nil, "101")

	fmt.Printf("%#v", a)

}
