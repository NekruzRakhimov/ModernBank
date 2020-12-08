package core

import (
	"OnlineBanking/db"
	"OnlineBanking/models"
	"database/sql"
	"fmt"
)

const UsersServicesText = `Выберите услугу:
1.Мои счета
2.Перевод денег
3.Оплата комунальных услуг
4.История транзакций
5.Список банкоматов
6.Выход`

const PublicServicesText = `Выберите услугу которую хотите оплатить
1.Электричество
2.Водаканал
3.Кабельное ТВ
4.Выход`

func UsersServices(database *sql.DB, User models.User) {
	for {
		fmt.Println("		Пользователь >", User.Name, User.Surname)
		fmt.Println(UsersServicesText)
		var cmd int64
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			PrintingListOfAccounts(database, User)
		case 2:
			Transfer(database, User)
		case 3:
			PublicServices(database, User)
		case 4:
			ShowTransactionsHistory(database, User)
		case 5:
			models.PrintingListOfATMs(database)
		case 6:
			return
		default:
			fmt.Println("Вы ввели некоректные данные. Попробуйте ещё раз!")
		}
	}
}

func ShowTransactionsHistory(database *sql.DB, user models.User) {
	PrintingListOfAccounts(database, user)
	fmt.Println("Введите номер счёта чтобы увидеть его историю:")
	var account string
	fmt.Scan(&account)
	rows, err := database.Query(db.SelectTransactionsHistory, account)
	if err != nil {
		fmt.Println("Error while printing list of transactions. Error is", err)
	}
	fmt.Println("№", "Номер счёта ", "Номер счёта получателя", "Сумма", "Дата", "Время")
	for rows.Next() {
		transaction := models.TransactionHistory{}
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

func CheckPublicServices(account models.Account, amount int64) (status bool) {
	status = true
	if account.Removed == true {
		fmt.Println("Введённый аккаунт удалён.")
		status = false
	} else if account.Number == "" {
		fmt.Println("Вы ввели неверный номер счёта!")
		status = false
	} else if amount > account.Amount {
		fmt.Println("У вас недостаточно средств")
		status = false
	}
	return status
}

func PublicServicesOperation(database *sql.DB) {
	fmt.Println("Введите номер счёта через который хотите оплатить данную услугу:")
	var number string
	fmt.Scan(&number)
	fmt.Println("Введите сумму которую хотите оплатить:")
	var amount int64
	fmt.Scan(&amount)
	myAccount := SearchingAccountByNumber(database, number)
	status := CheckPublicServices(myAccount, amount)
	if status == true {
		models.AddNewTransactionHistory(database, myAccount.Number, "Комунальные услуги", amount)
		myAccount.Amount -= amount
		database.Exec(db.UpadateAccountsAmount, myAccount.Amount, myAccount.Number)
		fmt.Println("Ваша оплата прошла успешно")
	}
}

func PublicServices(database *sql.DB, user models.User) {
	fmt.Println(PublicServicesText)
	var cmd int64
	fmt.Scan(&cmd)
	if cmd >= 1 && cmd <= 3 {
		PrintingListOfAccounts(database, user)
		PublicServicesOperation(database)
	}
}

func PrintingListOfAccounts(database *sql.DB, User models.User) {
	rows, err := database.Query(db.SelectAccountByUser_id, User.ID)
	if err != nil {
		fmt.Println("Error while printing list of accounts. Error is", err)
	}
	fmt.Println("№", "Баланс", "Номер карты          ", "Валюта", "Система")
	for rows.Next() {
		account := models.Account{}
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
		} else {
			fmt.Println(account.ID, account.Amount, " ", account.Number, " ", account.Currency, " ", account.System)
		}
	}
}

func TransferQuestions() (amount int64, from string, to string) {
	fmt.Println("Введите номер своего счёта:")
	fmt.Scan(&from)
	fmt.Println("Введите номер счёта на который нужно перевести деньги:")
	fmt.Scan(&to)
	fmt.Println("Введите сумму которую хотите перевести:")
	fmt.Scan(&amount)
	return amount, from, to
}

func SearchingAccountByNumber(database *sql.DB, number string) (Account models.Account) {
	err := database.QueryRow(db.SelectAccountByNumber, number).Scan(
		&Account.ID,
		&Account.User_id,
		&Account.Amount,
		&Account.Number,
		&Account.System,
		&Account.Currency,
		&Account.Removed,
	)
	if err != nil {
		fmt.Println("Error while selecting an account. Error is^", err)
	}
	return Account
}

func CheckTransfer(amount int64, accountFrom, accountTo models.Account) (status bool) {
	status = true
	if accountFrom.Number == "" {
		status = false
		fmt.Println("Вы неправильно ввели номер своего счёта!")
	} else if accountTo.Removed == true {
		status = false
		fmt.Println("Этот счёт", accountTo.Number, "удалённый")
	} else if accountTo.Number == "" {
		status = false
		fmt.Println("Вы неправильно ввели номер счёта на который хотите перевести деньги!")
	} else if amount > accountFrom.Amount {
		status = false
		fmt.Println("У вас на счету недостаточно средств!")
	}
	return status
}

func TransferOperation(database *sql.DB, amount int64, accountFrom, accountTo models.Account) {
	models.AddNewTransactionHistory(database, accountFrom.Number, accountTo.Number, amount)
	accountFrom.Amount -= amount
	accountTo.Amount += amount
	database.Exec(db.UpadateAccountsAmount, accountFrom.Amount, accountFrom.Number)
	database.Exec(db.UpadateAccountsAmount, accountTo.Amount, accountTo.Number)
	fmt.Println("Перевод совершён успешно!")
}

func Transfer(database *sql.DB, User models.User) {
	fmt.Println("---> Transfer")
	PrintingListOfAccounts(database, User)
	amount, from, to := TransferQuestions()
	accountFrom := SearchingAccountByNumber(database, from)
	accountTo := SearchingAccountByNumber(database, to)
	status := CheckTransfer(amount, accountFrom, accountTo)
	if status == true {
		TransferOperation(database, amount, accountFrom, accountTo)
	}
}
