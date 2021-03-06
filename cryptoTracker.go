package main

import (
	"cryptoTracker/apiUtils"
	"cryptoTracker/cortex"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	walletAddress := os.Args[1]
	blockchain := os.Args[2]

	data := [][]string{{"Date", "Wallet Address", "Sender Address", "Value (CTXC)", "Value (CAD)", "Transaction Hash"}}

	switch blockchain {
	case "cortex":
		fmt.Println("Calling cortex labs API...")
		accountInfo := cortex.GetAccountInfo(walletAddress)
		balance := cortex.GetAccountBalance(accountInfo)
		fmt.Printf("%v wallet balance: %.5f \n", blockchain, balance)
		accountTransactions := cortex.GetAccountTransactions(walletAddress)
		fmt.Printf("Your wallet has %v transactions \n", accountTransactions.Length)

		priceStore := make(map[string]string)
		for _, transaction := range accountTransactions.Result {
			transactionPrice := apiUtils.GetTransactionValue(blockchain, transaction.Timestamp, priceStore)
			dataRow := []string{time.Unix(transaction.Timestamp, 0).String(),
				transaction.To,
				transaction.From,
				transaction.Value,
				transactionPrice,
				transaction.Hash}
			data = append(data, dataRow)
		}

		file, err := os.Create("cortex-records.csv")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, value := range data {
			err := writer.Write(value)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
