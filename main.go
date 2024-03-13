package main

import (
	"fmt"
	"storiChallenge/storiChallenge"
)

func main() {
	filePath := "./test1.csv"
	transactions, err := storiChallenge.ReadTransactions(filePath)
	if err != nil {
		fmt.Println("Error reading transactions:", err)
		return
	}

	debitTotal, creditTotal := storiChallenge.SummarizeTransactions(transactions)
	fmt.Printf("Total Debit: %.2f, Total Credit: %.2f\n", debitTotal, creditTotal)

	subject := "Transaction Summary"
	body := fmt.Sprintf("Transaction summary:\nTotal Debit: %.2f\nTotal Credit: %.2f", debitTotal, creditTotal)
	err = storiChallenge.SendEmail(subject, body)
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}
