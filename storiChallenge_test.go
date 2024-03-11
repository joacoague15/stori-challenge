package storiChallenge

import (
	"os"
	"testing"
)

func TestReadTransactions(t *testing.T) {
	testFileName := "test_transactions.csv"
	testData := "1,2022-01-01,100.00\n2,2022-01-02,-50.00\n3,2022-01-03,25.00"
	err := os.WriteFile(testFileName, []byte(testData), 0644)
	if err != nil {
		t.Fatalf("Error writing test file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Error doing cleanup in test: %v", err)
		}
	}(testFileName)

	transactions, _ := readTransactions(testFileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(transactions) != 3 {
		t.Errorf("Expected 3 valid transactions, got %d", len(transactions))
	}
}
