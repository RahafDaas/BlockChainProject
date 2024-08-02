package main

import (
	"apartment-rental/internal/handler"
	"fmt"

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
	// handler := handler.ApartmentHandler{
	// 	RepoContract: *memory.New(),
	// }

}
