package storiChallenge

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	ID     string
	Date   string
	Amount float64
}

func readTransactions(filePath string) ([]Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []Transaction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		transactionParts := strings.Split(line, ",")
		// Skip no valid entries
		if len(transactionParts) < 3 {
			continue
		}
		transactionAmount, err := strconv.ParseFloat(transactionParts[2], 64)
		if err != nil {
			continue
		}
		transactions = append(transactions, Transaction{
			ID:     transactionParts[0],
			Date:   transactionParts[1],
			Amount: transactionAmount,
		})
	}
	return transactions, scanner.Err()
}

func summarizeTransactions(transactions []Transaction) (debitTotal float64, creditTotal float64) {
	for _, t := range transactions {
		if t.Amount < 0 {
			debitTotal += t.Amount
		} else {
			creditTotal += t.Amount
		}
	}
	return debitTotal, creditTotal
}
