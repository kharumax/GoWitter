package model

import (
	"database/sql"
)

func CheckUserExist(email string) bool  {
	user := User{}
	err := db.QueryRow("SELECT id FROM users WHERE email = ?",email).Scan(&user.Id)
	if err != nil && err == sql.ErrNoRows{
		return false
	}
	return true
}

func CreateUser(email,password string) (User,error)  {

}
