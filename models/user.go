package models

import (
	"OnlineBanking/db"
	"database/sql"
	"fmt"
)

type User struct {
	ID int64
	Name string
	Surname string
	Age int64
	Gender string
	Role string
	Login string
	Password string
	Removed bool
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
	for rows.Next(){
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

func RemoveUserById(database *sql.DB, id int64) () {
	_, err := database.Exec(db.RemoveUser, id)
	if err != nil {
		fmt.Println(err)
	}
}