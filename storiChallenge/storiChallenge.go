package storiChallenge

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID     string
	Date   string
	Amount float64
}

func ReadTransactions(filePath string) ([]Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var transactions []Transaction

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue // Skip over any other errors
		}

		// Skip the incomplete transactions
		if len(record) < 3 {
			continue
		}

		amountStr := record[2]
		// Ensure the amount starts with a valid sign and is not empty
		if amountStr == "" || (amountStr[0] != '+' && amountStr[0] != '-') {
			continue
		}

		// Parse the amount string to a float64
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			continue
		}

		transactions = append(transactions, Transaction{
			ID:     record[0],
			Date:   record[1],
			Amount: amount,
		})
	}

	return transactions, nil
}

func SendEmail(subject, bodyHTML string) error {
	from := "storichallenge7@gmail.com"
	pass := "yybhkcalxermxdky"
	to := "storichallenge7@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	headers := []string{
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
		"From: " + from,
		"To: " + to,
		"Subject: " + subject,
	}

	header := strings.Join(headers, "\r\n")

	message := []byte(header + "\r\n\r\n" + bodyHTML)

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, []string{to}, message)
	return err
}

func TotalBalance(transactions []Transaction) float64 {
	total := 0.0
	for _, transaction := range transactions {
		total += transaction.Amount
	}
	return total
}

func AverageDebitAndCredit(transactions []Transaction) (averageDebit, averageCredit float64) {
	var totalDebit, totalCredit float64
	var debitCount, creditCount int

	for _, transaction := range transactions {
		if transaction.Amount < 0 {
			totalDebit += transaction.Amount
			debitCount++
		} else if transaction.Amount > 0 {
			totalCredit += transaction.Amount
			creditCount++
		}
	}

	// This is to avoid division by zero
	if debitCount > 0 {
		averageDebit = totalDebit / float64(debitCount)
	}
	if creditCount > 0 {
		averageCredit = totalCredit / float64(creditCount)
	}

	return averageDebit, averageCredit
}

func TransactionsByMonth(transactions []Transaction) map[int][]Transaction {
	transactionsByMonth := make(map[int][]Transaction)

	for _, transaction := range transactions {
		dateParts := strings.Split(transaction.Date, "/")
		if len(dateParts) < 2 {
			fmt.Printf("Invalid date format for transaction ID %s: %s\n", transaction.ID, transaction.Date)
			continue
		}
		month, err := strconv.Atoi(dateParts[0])
		if err != nil {
			fmt.Printf("Error parsing month for transaction ID %s: %v\n", transaction.ID, err)
			continue
		}

		transactionsByMonth[month] = append(transactionsByMonth[month], transaction)
	}

	return transactionsByMonth
}

func PrepareMonthlyTransactionsCountDisplay(transactionsByMonth map[int][]Transaction) string {
	var sb strings.Builder

	for month, transactions := range transactionsByMonth {
		monthName := time.Month(month).String() // Convert month number to name
		transactionCount := len(transactions)
		sb.WriteString(fmt.Sprintf("<p>Number of transactions in %s: %d</p>", monthName, transactionCount))
	}

	return sb.String()
}
