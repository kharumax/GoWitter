package model

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
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
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil
}

func GetUserById(id int) (User,error)  {
	user := User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = ?",id).
		Scan(&user.Id,&user.Email,&user.Name,&user.ProfileImage,&user.Description,&user.Password)
	if err != nil {
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil
}

func GetCurrentUser(r *http.Request) (User,bool,error)  {
	//ここでセッションからEmailを取得する
	user := User{}
	sessionId,cookieError := r.Cookie("sessionId")
	if cookieError != nil {
		return user,false,cookieError
	}
	// []byte()
	sessionDecode,err :=  base64.RawStdEncoding.DecodeString(sessionId.Value)
	if err != nil {
		return user,false,err
	}
	// email = string(sessionDecode)
	user,getUserError := GetUser(string(sessionDecode))
	if getUserError != nil {
		return user,false,getUserError
	}
	return user,true,nil

}
