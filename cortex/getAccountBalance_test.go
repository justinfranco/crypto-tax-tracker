package cortex

import "testing"

func TestGetAccountBalance(t *testing.T) {
	var accInfo accountInfo
	accInfo.Balance = "12345600000000000000"
	balance := GetAccountBalance(&accInfo)
	if balance != 12.3456 {
		t.Errorf("Balance is %v but was suppose to be 12.3456", balance)
	}
}
