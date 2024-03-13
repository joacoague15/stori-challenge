package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import for side effects
	"log"
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

	// Creating the local DB
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the transaction table
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY,
		date TEXT,
		amount REAL
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// Add the new transactions into the db
	tr, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tr.Prepare("INSERT INTO transactions (id, date, amount) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, txn := range transactions {
		_, err = stmt.Exec(txn.ID, txn.Date, txn.Amount)
		if err != nil {
			// If an error occurs, rollback the transaction
			tr.Rollback()
			log.Fatal(err)
		}
	}

	// If everything went well, commit the transaction
	err = tr.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
