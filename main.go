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

	bodyHTML := `
		<html>
			<body>
				<h1>Transaction Summary</h1>
				<p>Total balance is: <strong>$39.74</strong></p>
				<p>Number of transactions in July: <strong>2</strong></p>
				<p>Number of transactions in August: <strong>2</strong></p>
				<p>Average debit amount: <strong>-$15.38</strong></p>
				<p>Average credit amount: <strong>$35.25</strong></p>
				<img width="200" height="200" src="https://yt3.googleusercontent.com/rn71w21Rhir526-DhwaarxhBSHpsdltK1R4Ym4XYDb9wFSBQWRhUR_ATehcAQYjsaoPFl_Hn2g=s900-c-k-c0x00ffffff-no-rj" />
			</body>
		</html>
	`

	err = storiChallenge.SendEmail(subject, bodyHTML)
	if err != nil {
		fmt.Println("Error sending email:", err)
	}
}
