package main

import (
	"OnlineBanking/db"
	"OnlineBanking/models"
	"OnlineBanking/pkg/core"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const WelcomeText = `Выберите команду:
1.Авторизация
2.Список банкоматов`

func Start(database *sql.DB) {
	fmt.Println("Добро пожаловать!")
	for {
		fmt.Println(WelcomeText)
		var cmd int64
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			core.Authorization(database)
		case 2:
			models.PrintingListOfATMs(database)
		default:
			fmt.Println("Вы ввели некоректные данные. Попробуйте ещё раз!")
		}
	}
}

func main() {
	database, err := sql.Open("sqlite3", "OnlineBankingDB")
	if err != nil {
		log.Fatal("Can't open DB. Error is:", err)
	}
	db.DBinit(database)
	Start(database)
}
