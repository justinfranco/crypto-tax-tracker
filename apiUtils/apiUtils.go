package apiUtils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func CallJsonAPI(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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

	return bodyBytes
}

func GetTransactionValues(coinId string, currency string, startTimestamp int64, endTimestamp int64) map[string]string {
	priceStore := make(map[string]string)

	// Add/subtract 86,400 seconds to the timestamps to get the full price data for the entire day
	url := "https://api.coingecko.com/api/v3/coins/" + coinId + "/market_chart/range?vs_currency=" + currency + "&from=" + strconv.FormatInt((startTimestamp - 86400), 10) + "&to=" + strconv.FormatInt((endTimestamp + 86400), 10)

	bodyBytes := CallJsonAPI(url)

	var responseObject map[string][][]float64
	err := json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Currently records the highest prices found for each day
	for _, dailyData := range responseObject["prices"] {
		transactionDate := time.UnixMilli(int64(dailyData[0])).Format("2006-01-02")
		transactionValuefloat := dailyData[1]
		transactionValueString := strconv.FormatFloat(dailyData[1], 'f', -1, 64)

		if currentValueString, exists := priceStore[transactionDate]; exists {
			currentValueFloat, err := strconv.ParseFloat(currentValueString, 64)
			if err != nil {
				fmt.Println(err.Error())
			}
			if transactionValuefloat > currentValueFloat {
					priceStore[transactionDate] = transactionValueString
			}
		} else {
			priceStore[transactionDate] = transactionValueString
		}
	}

	return priceStore
}
