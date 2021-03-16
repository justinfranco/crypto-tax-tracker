package cortex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type accountInfo struct {
	Addressid string  `json:"addressId"`
	Ctxcprice float64 `json:"ctxcprice"`
	Balance   string  `json:"balance"`
	Timestamp int64   `json:"timestamp"`
	Type      int     `json:"type"`
	Code      string  `json:"code"`
}

type accountTransactions struct {
	Txcount int `json:"txCount"`
	Length  int `json:"length"`
	Result  []struct {
		Value       string `json:"value"`
		Type        string `json:"type"`
		Gasprice    string `json:"gasPrice"`
		Gasused     int    `json:"gasUsed"`
		Status      string `json:"status"`
		Timestamp   int    `json:"timestamp"`
		From        string `json:"from"`
		To          string `json:"to"`
		Frombalance string `json:"fromBalance"`
		Tobalance   string `json:"toBalance"`
		Blocknumber int    `json:"blockNumber"`
		Hash        string `json:"hash"`
		Nonce       int    `json:"nonce"`
	} `json:"result"`
}

func GetAccountInfo(walletAddress string) *accountInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://cerebro.cortexlabs.ai/mysql?addressId="+walletAddress+"&type=accountInfo", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var responseObject accountInfo
	json.Unmarshal(bodyBytes, &responseObject)

	return &responseObject
}

func GetAccountBalance(accountInfo *accountInfo) float64 {
	balance, err := strconv.ParseFloat(accountInfo.Balance, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	balance = balance / float64(1000000000000000000)

	return balance
}
