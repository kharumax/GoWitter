package model

import (
	"database/sql"
)

func IsUserExist(email string) bool  {
	user := User{}
	err := db.QueryRow("SELECT id FROM users WHERE email = ?",email).Scan(&user.Id)
	if err != nil && err == sql.ErrNoRows{
		return false
	}
	return true
}

func CreateUser(email,password string) (User,error)  {
	user := User{}
	user.Email = email
	stmt,err := db.Prepare(`INSERT INTO users (email,password) VALUES (?,?)`)
	if err != nil {
		return user,err
	}
	defer stmt.Close()
	res,insertError := stmt.Exec(email,password)
	if insertError != nil {
		return user,err
	}
	id,getIdError := res.LastInsertId()
	if getIdError != nil {
		 return user,err
	}
	user.Id = int(id)
	return user,nil
}

func GetUser(email string) (User,error)  {
	user := User{}
	err := db.QueryRow("SELECT * FROM users WHERE email = ?",email).
		Scan(&user.Id,&user.Email,&user.Name,&user.ProfileImage,&user.Description,&user.Password)
	if err != nil {
		return user,err
	}
	return user,nil
}

