package main

import (
	"cryptoTracker/cortex"
	"fmt"
	"os"
)

func main() {
	walletAddress := os.Args[1]
	blockchain := os.Args[2]

	switch blockchain {
	case "cortex":
		fmt.Println("Calling cortex labs API...")
		accountInfo := cortex.GetAccountInfo(walletAddress)
		balance := cortex.GetAccountBalance(accountInfo)
		fmt.Printf("%v wallet balance: %.5f", blockchain, balance)
	}

	fmt.Println("")
}
