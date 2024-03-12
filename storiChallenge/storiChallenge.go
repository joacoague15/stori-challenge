package storiChallenge

import (
	"encoding/csv"
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
			if err.Error() == "EOF" {
				break
			}
			continue
		}

		// Skip the incomplete transactions
		if len(record) < 3 {
			continue
		}

		// Skip records with invalid debit or credit amounts
		amount, err := strconv.ParseFloat(record[2], 64)
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

//func sendEmail(subject, body string) error {
//	from := "your-email@example.com"
//	pass := "yourpassword"
//	to := "recipient@example.com"
//	smtpHost := "smtp.example.com"
//	smtpPort := "587"
//
//	auth := smtp.PlainAuth("", from, pass, smtpHost)
//
//	message := []byte("To: " + to + "\r\n" +
//		"Subject: " + subject + "\r\n\r\n" +
//		body)
//
//	addr := smtpHost + ":" + smtpPort
//	err := smtp.SendMail(addr, auth, from, []string{to}, message)
//	return err
//}
