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
	if transactionvalues["2021-10-21"] != "0.28344685601454134" {
		t.Errorf("Value was %v but was suppose to be 0.28344685601454134", transactionvalues["2021-10-21"])
	}
}
