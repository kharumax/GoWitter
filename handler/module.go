package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"GoWitter/model"
)


// ユーザーがログインしている場合はその情報を、していない場合はErrorを返す
func getCurrentUser(r *http.Request) (model.User,error) {
	user := model.User{}
	sessionId,err := readSessionFromCookie(r)
	if err != nil {
		return user,err
	}
	user,err = model.GetUserById(sessionId)
	if err != nil {
		return user,err
	}
	return user,nil
}

// URLから「id」を取得する
func getUserIdFromURL(r *http.Request) (int,error)  {
	url := r.URL.Path
	urlSlice := strings.Split(url,"/")
	if len(urlSlice) < 3 {
		err := errors.New("URL Length Invalid")
		return 0,err
	}
	id,err := strconv.Atoi(urlSlice[2])
	if err != nil {
		return 0,err
	}
	return id,nil
}

func setSessionToCookie(w http.ResponseWriter,userId int) {
	sessionByte := []byte(strconv.Itoa(userId))
	sessionEncode := base64.RawStdEncoding.EncodeToString(sessionByte)
	http.SetCookie(w,&http.Cookie{
		Name: "sessionId",
		Value: sessionEncode,
	})
}

func readSessionFromCookie(r *http.Request) (int,error) {
	cookie,err := r.Cookie("sessionId")
	sessionEncode := cookie.Value
	if err != nil {
		return 0,err
	}
	fmt.Printf("sessionEncode is %v",sessionEncode)
	// ここでEncodeをデコードしてUserIDのbyteを取得する
	sessionByte,err := base64.RawStdEncoding.DecodeString(sessionEncode)
	if err != nil {
		return 0,err
	}
	// Byte -> String -> Int
	sessionId,err := strconv.Atoi(string(sessionByte))
	if err != nil {
		return 0,err
	}
	fmt.Printf("sessionId is %v",sessionId)
	return sessionId,nil
}

func deleteSessionFromCookie(w http.ResponseWriter,r *http.Request) error {
	cookie,err := r.Cookie("sessionId")
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	http.SetCookie(w,cookie)
	return nil
}