package models

import (
	"OnlineBanking/db"
	"database/sql"
	"fmt"
	time2 "time"
)

type TransactionHistory struct {
	ID int64
	MyAccount string
	ToAccount string
	Amount int64
	Data string
	Time string
}

func AddNewTransactionHistory(database *sql.DB, myAccount, toAccount string, Amount int64,) {
	time := time2.Now()
	data := time.Format("02.January.2006")
	Time := time.Format("15:40")
	_, err := database.Exec(db.AddTransactionHistory, myAccount, toAccount, Amount, data, Time)
	if err != nil {
		fmt.Println("Error")
	}
}

func PrintingListOfTransactions(database *sql.DB) {
	rows, err := database.Query("select * from transactionshistory")
	if err != nil {
		fmt.Println("Error while printing list of transactions. Error is", err)
	}
	fmt.Println("№",  "Номер счёта ", "Номер счёта получателя", "Сумма", "Дата", "Время")
	for rows.Next() {
		transaction := TransactionHistory{}
		err = rows.Scan(
			&transaction.ID,
			&transaction.MyAccount,
			&transaction.ToAccount,
			&transaction.Amount,
			&transaction.Data,
			&transaction.Time,
		)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(transaction.ID, transaction.MyAccount, transaction.ToAccount, transaction.Amount, transaction.Data, transaction.Time)
		}
	}

}

