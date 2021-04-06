package apiUtils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetTransactionValue(coinId string, timestamp int64, priceStore map[string]string) string {
	date := time.Unix(timestamp, 0)

	// If price for date already found no need to query the API again.
	price, ok := priceStore[date.Format("02-01-2006")]
	if ok {
		return price
	}

	url := "https://api.coingecko.com/api/v3/coins/" + coinId + "/history?date=" + date.Format("02-01-2006") + "&localization=false"

	bodyBytes := CallJsonAPI(url)

	var responseObject *interface{}
	json.Unmarshal(bodyBytes, &responseObject)

	dataMap := (*responseObject).(map[string]interface{})
	marketData := dataMap["market_data"].(map[string]interface{})
	currentPrices := marketData["current_price"].(map[string]interface{})
	value := fmt.Sprintf("%v", currentPrices["cad"])

	priceStore[date.Format("02-01-2006")] = value

	return value
}
