package main

import (
	"OnlineBanking/db"
	"OnlineBanking/models"
	"OnlineBanking/pkg/core"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const WelcomeText = `Выберите команду:
1.Авторизация
2.Список банкоматов
0.Выход`

func WelcomeWindow(database *sql.DB) {
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
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Вы ввели некоректные данные. Попробуйте ещё раз!")
		}
	}
}

func Start(database *sql.DB) {
	WelcomeWindow(database)
}

func main() {
	database, err := sql.Open("sqlite3", "OnlineBankingDB")
	if err != nil {
		log.Fatal("Can't open DB. Error is:", err)
	}
	db.DBinit(database)
	Start(database)
}
