package core

import (
	"OnlineBanking/models"
	"database/sql"
	"fmt"
)

const UserToolsText = `Выберите категорию:
1.Пользователи
2.Счета
3.Банкоматы
4.История транзвкций
5.Выход`

const UsersControlText = `Выберите команду:
1.Добавить пользователя
2.Обновить данные о пользователе
3.Удалить пользователя
4.Список всех пользователей
5.Назад`

const AccountsControlText = `Выберите команду:
1.Добавить счёт
2.Обновить данные о счёте
3.Удалить счёт
4.Список всех счетов
5.Назад`

const ATMsControlText = `Выберите команду:
1.Добавить банкомат
2.Обновить банкомат
3.Удалить банкомат
4.Список всех банкоматов
5.Назад`

func AdminsTools(database *sql.DB, user models.User) {
	for {
		fmt.Println("		Администратор >", user.Name, user.Surname)
		fmt.Println(UserToolsText)
		var cmd int64
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			UsersControl(database)
		case 2:
			AccountsControl(database)
		case 3:
			ATMsControl(database)
		case 4:
			models.PrintingListOfTransactions(database)
		case 5:
			return
		default:
			fmt.Println("Вы ввели некоректные данные. Попробуйте ещё!")
		}
	}
}

func AddUserQuestions(database *sql.DB) (user models.User){
	fmt.Println("Заполните все поля!")
	fmt.Println("Введите имя")
	fmt.Scan(&user.Name)
	fmt.Println("Введите фамилию")
	fmt.Scan(&user.Surname)
	fmt.Println("Введите возраст")
	fmt.Scan(&user.Age)
	fmt.Println("Введите пол")
	fmt.Scan(&user.Gender)
	fmt.Println("Введите логин")
	fmt.Scan(&user.Login)
	fmt.Println("Введите пароль")
	fmt.Scan(&user.Password)
	return user
}

func UpdatingUserQuestions(database *sql.DB) (user models.User, id int64) {
	fmt.Println("Заполните все поля!")
	fmt.Println("Введите ID счёта")
	fmt.Scan(&id)
	fmt.Println("Введите имя")
	fmt.Scan(&user.Name)
	fmt.Println("Введите фамилию")
	fmt.Scan(&user.Surname)
	fmt.Println("Введите возраст")
	fmt.Scan(&user.Age)
	fmt.Println("Введите пол")
	fmt.Scan(&user.Gender)
	fmt.Println("Введите роль: пользователь/администратор")
	fmt.Scan(&user.Role)
	fmt.Println("Введите логин")
	fmt.Scan(&user.Login)
	fmt.Println("Введите пароль")
	fmt.Scan(&user.Password)
	return user, id
}

func UsersControl(database *sql.DB)  {
	fmt.Println(UsersControlText)
	var cmd int64
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		user := AddUserQuestions(database)
		models.AddNewUser(database, user)
	case 2:
		user, id := UpdatingUserQuestions(database)
		models.UpdatingUser(database, user, id)
	case 3:
		fmt.Println("Введите ID пользователя:")
		var id int64
		fmt.Scan(&id)
		models.RemoveUserById(database, id)
	case 4:
		models.PrintingListOfUsers(database)
	}
}

func AddAccountQuestions(database *sql.DB) (account models.Account){
	fmt.Println("Заполните все поля!")
	fmt.Println("Введите ID владельца данного аккаунта")
	fmt.Scan(&account.User_id)
	account.Amount =  0
	fmt.Println("Введите номер сёта")
	fmt.Scan(&account.Number)
	fmt.Println("Введите платёжную систему счёта")
	fmt.Scan(&account.System)
	fmt.Println("Введите валюту")
	fmt.Scan(&account.Currency)
	fmt.Println("Введите пароль")
	return account
}

func UpdatingAccountQuestions(database *sql.DB) (account models.Account, id int64) {
		fmt.Println("Заполните все поля!")
		fmt.Println("Введите ID счёта")
		fmt.Scan(&id)
		fmt.Println("Введите ID владельца счёта")
		fmt.Scan(&account.User_id)
		fmt.Println("Введите название платёжной системы счёта")
		fmt.Scan(&account.System)
		fmt.Println("Введите валюту счёта")
		fmt.Scan(&account.Currency)
		return account, id
}

func UpdatingATMQuestions(database *sql.DB) (atm models.ATM){
	fmt.Println("Введите ID банкомата")
	fmt.Scan(&atm.ID)
	fmt.Println("Введите адрес банкомата")
	fmt.Scan(&atm.Address)
	return atm
}

func AccountsControl(database *sql.DB) {
	fmt.Println(AccountsControlText)
	var cmd int64
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		account := AddAccountQuestions(database)
		models.AddNewAccount(database, account)
	case 2:
		account, id := UpdatingAccountQuestions(database)
		models.UpdatingAccount(database, account, id)
	case 3:
		fmt.Println("Введите ID счёта:")
		var id int64
		fmt.Scan(&id)
		models.RemoveAccountById(database, id)
	case 4:
		models.PrintingListOfAccounts(database)
	}
}

func AddATMQuestions(database *sql.DB) (address string){
	fmt.Println("Введите адрес банкомата")
	fmt.Scan(&address)
	return address
}

func ATMsControl(database *sql.DB) {
	fmt.Println(ATMsControlText)
	var cmd int64
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		address := AddATMQuestions(database)
		models.AddNewATM(database, address)
	case 2:
		atm := UpdatingATMQuestions(database)
		models.UpdatingATM(database, atm)
	case 3:
		fmt.Println("Введите ID банкомата:")
		var id int64
		fmt.Scan(&id)
		models.RemoveATMById(database, id)
	case 4:
		models.PrintingListOfATMs(database)
	}
}