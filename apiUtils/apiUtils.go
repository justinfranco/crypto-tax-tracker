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

	url := "https://api.coingecko.com/api/v3/coins/" + coinId + "/market_chart/range?vs_currency=" + currency + "&from=" + strconv.FormatInt(startTimestamp, 10) + "&to=" + strconv.FormatInt(endTimestamp, 10)

	bodyBytes := CallJsonAPI(url)

	var responseObject map[string][][]float64
	err := json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO: This currently takes the last price of the day. It should probably average all the reported prices throughout a day.
	// However for now this will be accurate enough.
	for _, dailyData := range responseObject["prices"] {
		priceStore[time.UnixMilli(int64(dailyData[0])).Format("2006-01-02")] = strconv.FormatFloat(dailyData[1], 'f', -1, 64) 
	}

	return priceStore
}
