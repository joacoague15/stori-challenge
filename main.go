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

	totalBalance := storiChallenge.TotalBalance(transactions)
	averageDebit, averageCredit := storiChallenge.AverageDebitAndCredit(transactions)

	transactionsByMonth := storiChallenge.TransactionsByMonth(transactions)
	monthlyTransactionsContent := storiChallenge.PrepareMonthlyTransactionsCountDisplay(transactionsByMonth)

	subject := "Transaction Summary"

	bodyHTML := fmt.Sprintf(`
		<html>
			<body>
				<h1>Transaction Summary</h1>
				<p>Total balance is: <strong>$%.2f</strong></p>
				<p>Average debit amount: <strong>$%.2f</strong></p>
				<p>Average credit amount: <strong>$%.2f</strong></p>
				%s
				<img width="200" height="200" src="https://yt3.googleusercontent.com/rn71w21Rhir526-DhwaarxhBSHpsdltK1R4Ym4XYDb9wFSBQWRhUR_ATehcAQYjsaoPFl_Hn2g=s900-c-k-c0x00ffffff-no-rj" />
			</body>
		</html>
	`, totalBalance, averageDebit, averageCredit, monthlyTransactionsContent)

	err = storiChallenge.SendEmail(subject, bodyHTML)
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}
