package apiUtils

import "testing"

func TestGetTransactionValue(t *testing.T) {
	const coinId = "cortex"
	const unixTimestamp = 1615030537 //	Sat Mar 06 2021 11:35:37 GMT+0000
	priceStore := make(map[string]string)
	value := GetTransactionValue(coinId, unixTimestamp, priceStore)
	if value != "0.2207229133399098" {
		t.Errorf("Value was %v but was suppose to be 0.2207229133399098", value)
	}
}
