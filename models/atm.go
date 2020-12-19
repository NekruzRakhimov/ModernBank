package models

import (
	"OnlineBanking/db"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type ATM struct {
	ID      int64
	Address string
	Status  bool
}

func SavingATMsTableToXML(database *sql.DB) {
	var arr []ATM
	rows, err := database.Query("select * from atms")
	if err != nil {
		fmt.Println("Error while saving list of atms to xml file. Error is:", err)
	}
	for rows.Next() {
		atm := ATM{}
		err = rows.Scan(&atm.ID, &atm.Address, &atm.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		arr = append(arr, atm)
	}
	bytes, err := xml.Marshal(arr)
	err = ioutil.WriteFile("DBcopy/xml/atms.xml", bytes, 6604)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Данные успешно сохранены по адресу DataBaseCopy/xml/atms.xml")
	}
}

func SavingATMsTableToJSON(database *sql.DB) {
	var arr []ATM
	rows, err := database.Query("select * from atms")
	if err != nil {
		fmt.Println("Error while printing list of atms. Error is:", err)
	}
	for rows.Next() {
		atm := ATM{}
		err = rows.Scan(&atm.ID, &atm.Address, &atm.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		arr = append(arr, atm)
	}
	bytes, err := json.Marshal(arr)
	err = ioutil.WriteFile("DBcopy/json/atms.json", bytes, 6604)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Данные успешно сохранены по адресу DataBaseCopy/json/atms.json")
	}
}

func AddNewATM(database *sql.DB, address string) {
	_, err := database.Exec("INSERT INTO atms(address) values (($1))", address)
	if err != nil {
		fmt.Println("Can't add new ATM. Error is:", err)
	} else {
		fmt.Println("New ATM was successfully added!")
	}
}

func RemoveATMById(database *sql.DB, id int64) {
	_, err := database.Exec(db.RemoveATM, id)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdatingATM(database *sql.DB, atm ATM) {
	_, err := database.Exec(db.UpadateATM, atm.Address, atm.ID)
	if err != nil {
		fmt.Println("Error is ", err)
	}
}

func PrintingListOfATMs(database *sql.DB) {
	rows, err := database.Query("select * from atms")
	if err != nil {
		fmt.Println("Error while printing list of ATMs. Error is:", err)
	}
	fmt.Println("№", "Адрес:			", "Работает:")
	for rows.Next() {
		atm := ATM{}
		err = rows.Scan(&atm.ID, &atm.Address, &atm.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(atm.ID, " ", atm.Address)
		if atm.Status == true {
			fmt.Print("	да\n")
		} else {
			fmt.Print("	нет\n")
		}
	}
}
