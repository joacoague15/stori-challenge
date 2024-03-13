package storiChallenge

import (
	"os"
	"testing"
)

func TestReadTransactions(t *testing.T) {
	testFileName := "test_transactions.csv"
	testData := "1,7/15,+100.75\n2,1/20,-50.50\n3,10/2,+25"
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

func TestTotalBalance(t *testing.T) {
	transactions := []Transaction{
		{ID: "1", Date: "1/1", Amount: 100.25},
		{ID: "2", Date: "11/23", Amount: -50.25},
		{ID: "3", Date: "5/25", Amount: 25},
	}
	expectedTotal := 75.0
	total := TotalBalance(transactions)
	if total != expectedTotal {
		t.Errorf("Expected total balance = %.2f, got %.2f", expectedTotal, total)
	}
}

func TestCalculateAverageDebitAndCredit(t *testing.T) {
	transactions := []Transaction{
		{ID: "1", Date: "1/1", Amount: 100.25},
		{ID: "2", Date: "11/23", Amount: -50.25},
		{ID: "3", Date: "5/25", Amount: 25},
	}
	averageDebit, averageCredit := AverageDebitAndCredit(transactions)

	expectedDebitAverage := -50.25
	expectedCreditAverage := 62.625

	if averageDebit != expectedDebitAverage {
		t.Errorf("Expected a debit average of %.2f, got %.2f", expectedDebitAverage, averageDebit)
	}

	if averageCredit != expectedCreditAverage {
		t.Errorf("Expected a credit average of %.2f, got %.2f", expectedCreditAverage, averageCredit)
	}
}

func TestTransactionsByMonth(t *testing.T) {
	transactions := []Transaction{
		{ID: "1", Date: "1/1", Amount: 100.25},
		{ID: "2", Date: "11/23", Amount: -50.25},
		{ID: "3", Date: "5/25", Amount: 25},
	}

	expectedMonthDisplay := map[int]int{
		1:  1,
		5:  1,
		11: 1,
	}

	result := TransactionsByMonth(transactions)

	// Check if the number of transactions per month matches the expected number
	for month, expectedCount := range expectedMonthDisplay {
		transactions := result[month]

		if len(transactions) != expectedCount {
			t.Errorf("Expected %d transactions for month %d, but got %d", expectedCount, month, len(transactions))
		}
	}
}
