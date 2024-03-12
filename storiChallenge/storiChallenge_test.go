package storiChallenge

import (
	"os"
	"testing"
)

func TestReadTransactions(t *testing.T) {
	testFileName := "test_transactions.csv"
	testData := "1,7/15,100.75\n2,1-20,-50.50\n3,10-2,25"
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

	transactions, _ := ReadTransactions(testFileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(transactions) != 3 {
		t.Errorf("Expected 3 valid transactions, got %d", len(transactions))
	}
}

func TestSummarizeTransactions(t *testing.T) {
	transactions := []Transaction{
		{ID: "1", Date: "1-1", Amount: 100.25},
		{ID: "2", Date: "11-23", Amount: -50.25},
		{ID: "3", Date: "5-25", Amount: 25},
	}
	debitTotal, creditTotal := SummarizeTransactions(transactions)
	if debitTotal != -50.25 || creditTotal != 125.25 {
		t.Errorf("Expected debitTotal = 50.25 and creditTotal = 125.25, got debitTotal = %.2f and creditTotal = %.2f", debitTotal, creditTotal)
	}
}
