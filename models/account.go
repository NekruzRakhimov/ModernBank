package models

import (
	"OnlineBanking/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Account struct {
	ID       int64
	User_id  int64
	Amount   int64
	Number   string
	System   string
	Currency string
	Removed  bool
}

func SavingAccountsTableToJSON(database *sql.DB) {
	var arr []Account
	rows, err := database.Query("select * from accounts")
	if err != nil {
		fmt.Println("Error while printing list of Accounts. Error is:", err)
	}
	for rows.Next() {
		account := Account{}
		err = rows.Scan(
			&account.ID,
			&account.User_id,
			&account.Amount,
			&account.Number,
			&account.System,
			&account.Currency,
			&account.Removed,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		arr = append(arr, account)
	}
	bytes, err := json.Marshal(arr)
	err = ioutil.WriteFile("DBcopy/accounts.json", bytes, 6604)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Данные успешно сохранены по адресу DataBaseCopy/accounts.json")
	}
}

func RemoveAccountById(database *sql.DB, id int64) {
	_, err := database.Exec(db.RemoveAccount, id)
	if err != nil {
		fmt.Println(err)
	}
}

func AddNewAccount(database *sql.DB, account Account) {
	_, err := database.Exec(db.AddAccount, account.User_id, account.Amount, account.Number, account.System, account.Currency)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdatingAccount(database *sql.DB, account Account, id int64) {
	_, err := database.Exec(db.UpadateAccount, id, account.User_id, account.System, account.Currency, id)
	if err != nil {
		fmt.Println("Error is ", err)
	}
}

func PrintingListOfAccounts(database *sql.DB) {
	rows, err := database.Query("select * from accounts")
	if err != nil {
		fmt.Println("Error while printing list of Accounts. Error is:", err)
	}
	fmt.Println("№", "User_id", "amount", "number", "system", "currency")
	for rows.Next() {
		account := Account{}
		err = rows.Scan(
			&account.ID,
			&account.User_id,
			&account.Amount,
			&account.Number,
			&account.System,
			&account.Currency,
			&account.Removed,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if account.Removed == true {
			continue
		}
		fmt.Println(account.ID, account.User_id, account.Amount, account.Number, account.System, account.Currency)
	}
}
