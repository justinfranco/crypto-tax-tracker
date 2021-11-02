package apiUtils

import (
	"testing"
)

func TestGetTransactionValues(t *testing.T) {
	const coinId = "cortex"
	const currency = "cad"
	const unixTimestampStart = 1634734226 // Wed Oct 20 2021 12:50:26 GMT+0000
	const unixTimestampEnd = 1635743460 // Sun Oct 31 2021 23:11:00 GMT-0600
	transactionvalues := GetTransactionValues(coinId, currency, unixTimestampStart, unixTimestampEnd)
	// 0.2884155723678864 is the largest value recorded on coingecko on 2021-10-21 (between 1634796000 - 1634882399)
	if transactionvalues["2021-10-21"] != "0.2884155723678864" {
		t.Errorf("Value was %v but was suppose to be 0.2884155723678864", transactionvalues["2021-10-21"])
	}
}
