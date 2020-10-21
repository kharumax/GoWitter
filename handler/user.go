package handler

import (
	"GoWitter/model"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignUpError struct {
	ErrorMessage string
}

func SignUpHandler(w http.ResponseWriter,r *http.Request)  {
	errorMsg := SignUpError{}
	if r.Method == http.MethodGet {
		//サインページを出力する
		errorMsg.ErrorMessage = ""
		tpl.ExecuteTemplate(w,"signup.html",errorMsg)
		return
	}
	if r.Method == http.MethodPost {
		//ここで会員登録情報を送られた場合の処理
		// email,password,password_confirmationを受け取る
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		passwordConfirmation := r.PostFormValue("password_confirmation")
		if email == "" || password == "" || passwordConfirmation == "" {
			//ここでどれかが空の場合にエラーを返す
			errorMsg.ErrorMessage = "未記入の欄があります。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		if password != passwordConfirmation {
		 	errorMsg.ErrorMessage = "パスワードが違います。"
		 	tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		if len(password) < 6 {
			errorMsg.ErrorMessage = "パスワードは最低６文字以上です。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		//ここで送られてきたemailが既に登録されているかを確認する
		checkUserExist := model.CheckUserExist(email)
		if checkUserExist {
			errorMsg.ErrorMessage = "そのメールアドレスは既に利用されています。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		//ここからパスワードのハッシュ化を行い、データベースに登録する
		hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil {
			errorMsg.ErrorMessage = "ネットワークの障害が発生しました。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
		}
		password_digest := string(hash)
		//hashFromStr := []byte(string(hash))
		//err := bcrypt.CompareHashAndPassword(hashFromStr,[]byte(password))
		//fmt.Println("This is compare : ",err)
	}
}
