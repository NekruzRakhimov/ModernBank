package core

import (
	"OnlineBanking/models"
	"database/sql"
	"fmt"
)

const AuthorizationFailedText = `Неправильный логин или пароль. Выберите команду:
1.Попробовать ещё
2.Выйти`

func AskingLogin () (login, password string) {
	fmt.Println("Введите логин и пароль")
	fmt.Println("Введите логин")
	fmt.Scan(&login)
	fmt.Println("Введите пароль")
	fmt.Scan(&password)
	return
}

func SearchingUserInDB(database *sql.DB, login, password string)(User models.User) {
	database.QueryRow(`SELECT * FROM users WHERE login = ($1) and password = ($2)`, login, password).Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Role,
		&User.Login,
		&User.Password,
		&User.Removed,
	)
	return User
}

func Authorization(database *sql.DB) {
	login, password := AskingLogin()
	User := SearchingUserInDB(database, login, password)
	if User.Login == "" && User.Password == ""{
		fmt.Println(AuthorizationFailedText)
		var cmd int64
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			Authorization(database)
		case 2:
			return
		default:
			fmt.Println("Вы ввели некоректные данные")
		}
	} else {
		fmt.Println("____________________С возвращением!____________________")
		if User.Role == "user" {
			UsersServices(database, User)
		}else if User.Role == "admin" {
			AdminsTools(database, User)
		}
	}
}
