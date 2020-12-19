package models

import (
	"OnlineBanking/db"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type User struct {
	ID       int64  `xml:"id" json:"id"`
	Name     string `xml:"name" json:"name"`
	Surname  string `xml:"surname" json:"surname"`
	Age      int64  `xml:"age" json:"age"`
	Gender   string `xml:"gender" json:"gender"`
	Role     string `xml:"role" json:"role"`
	Login    string `xml:"login" json:"login"`
	Password string `xml:"password" json:"password"`
	Removed  bool   `xml:"removed" json:"removed"`
}

func SavingUsersTableToXML(database *sql.DB) {
	var arr []User
	rows, err := database.Query("select * from users")
	if err != nil {
		fmt.Println("Error while saving list of Users to xml file. Error is:", err)
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Age,
			&user.Gender,
			&user.Role,
			&user.Login,
			&user.Password,
			&user.Removed,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		arr = append(arr, user)
	}
	bytes, err := xml.Marshal(arr)
	err = ioutil.WriteFile("DBcopy/xml/users.xml", bytes, 6604)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Данные успешно сохранены по адресу DataBaseCopy/xml/users.xml")
	}
}

func SavingUsersTableToJSON(database *sql.DB) {
	var arr []User
	rows, err := database.Query("select * from users")
	if err != nil {
		fmt.Println("Error while printing list of Users. Error is:", err)
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Age,
			&user.Gender,
			&user.Role,
			&user.Login,
			&user.Password,
			&user.Removed,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		arr = append(arr, user)
	}
	bytes, err := json.Marshal(arr)
	err = ioutil.WriteFile("DBcopy/json/users.json", bytes, 6604)
	if err != nil {
		fmt.Println("Error is", err)
	} else {
		fmt.Println("Данные успешно сохранены по адресу DataBaseCopy/json/users.json")
	}
}

func AddNewUser(database *sql.DB, user User) {
	_, err := database.Exec(db.AddUser, user.Name, user.Surname, user.Age, user.Gender, user.Role, user.Login, user.Password)
	if err != nil {
		fmt.Println("Error is", err)
	}
}

func UpdatingUser(database *sql.DB, user User, id int64) {
	_, err := database.Exec(db.UpadateUser, user.Name, user.Surname, user.Age, user.Gender, user.Role, user.Login, user.Password, id)
	if err != nil {
		fmt.Println("Error is ", err)
	}
}

func PrintingListOfUsers(database *sql.DB) {
	rows, err := database.Query("select * from users")
	if err != nil {
		fmt.Println("Error while printing list of Users. Error is:", err)
	}
	fmt.Println("№", "Name:	", "Surname	", "Age", "Gender	", "Role	", "Login  ", "Password ")
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Age,
			&user.Gender,
			&user.Role,
			&user.Login,
			&user.Password,
			&user.Removed,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if user.Removed == true {
			continue
		}
		fmt.Println(user.ID, user.Name, user.Surname, user.Age, user.Gender, user.Role, user.Login, user.Password)
	}
}

func RemoveUserById(database *sql.DB, id int64) {
	_, err := database.Exec(db.RemoveUser, id)
	if err != nil {
		fmt.Println(err)
	}
}
