package main

import (
	"fmt"
	"simple-poa-blockchain-peer/orchestrator/client"
)

func main() {
	res, err := client.GetAllPeers("localhost", ":8080")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}

