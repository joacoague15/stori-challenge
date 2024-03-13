package storiChallenge

import (
	"encoding/csv"
	"io"
	"net/smtp"
	"os"
	"strconv"
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

func SummarizeTransactions(transactions []Transaction) (debitTotal float64, creditTotal float64) {
	for _, t := range transactions {
		if t.Amount < 0 {
			debitTotal += t.Amount
		} else {
			creditTotal += t.Amount
		}
	}
	return debitTotal, creditTotal
}

func SendEmail(subject, body string) error {
	from := "storichallenge7@gmail.com"
	pass := "yybhkcalxermxdky"
	to := "storichallenge7@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body)

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, []string{to}, message)
	return err
}
