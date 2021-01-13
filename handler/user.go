package handler

import (
	"GoWitter/model"
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignUpError struct {
	ErrorMessage string
}
type LoginError struct {
	Email string
	Password string
	ErrorMessage string
}

func SignUpHandler(w http.ResponseWriter,r *http.Request)  {
	errorMsg := SignUpError{}
	// 既にログインしているかどうかを判断する関数を作る
	_,err := getCurrentUser(r)
	if err == nil {
		// 既にログインしている場合は、一旦indexに飛ばす
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

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
		isUserExist := model.IsUserExist(email)
		if isUserExist {
			// 既にそのメアドがDBに保存されていた場合
			errorMsg.ErrorMessage = "そのメールアドレスは既に利用されています。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		//ここからパスワードのハッシュ化を行い、データベースに登録する
		hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil {
			errorMsg.ErrorMessage = "ネットワークの障害が発生しました。"
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		passwordDigest := string(hash)
		user,insertError := model.CreateUser(email,passwordDigest)
		if insertError != nil {
			errorMsg.ErrorMessage = insertError.Error()
			tpl.ExecuteTemplate(w,"signup.html",errorMsg)
			return
		}
		setSessionToCookie(w,user.Id)
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}
	errorMsg.ErrorMessage = "無効なHTTPリクエストです。"
	tpl.ExecuteTemplate(w,"signup.html",errorMsg)
	return
}


func LoginHandler(w http.ResponseWriter,r *http.Request)  {
	errorMsg := LoginError{}
	// 既にログインしているかどうかを判断する関数を作る
	_,err := getCurrentUser(r)
	if err == nil {
		// 既にログインしている場合は、一旦indexに飛ばす
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w,"login.html",errorMsg)
		return
	}
	if r.Method == http.MethodPost {
		//ここにPOSTリクエストの場合の処理を記述する
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		user,getUserError := model.GetUser(email)
		if getUserError != nil && getUserError == sql.ErrNoRows {
			errorMsg.ErrorMessage = "指定されたユーザーは存在しません。"
			tpl.ExecuteTemplate(w,"login.html",errorMsg)
			return
		}
		if getUserError != nil && getUserError != sql.ErrNoRows {
			errorMsg.ErrorMessage = getUserError.Error()
			tpl.ExecuteTemplate(w,"login.html",errorMsg)
			return
		}
		//ここでパスワードの確認を行う
		passwordFromHash := []byte(user.Password)
		err := bcrypt.CompareHashAndPassword(passwordFromHash,[]byte(password))
		if err != nil {
			errorMsg.ErrorMessage = "emailかパスワードが間違っています。"
			tpl.ExecuteTemplate(w,"login.html",errorMsg)
			return
		}
		// ユーザーIDをEncodeしてCookieに保存する処理
		setSessionToCookie(w,user.Id)
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}
	errorMsg.ErrorMessage = "無効なHTTPリクエストです。"
	tpl.ExecuteTemplate(w,"login.html",errorMsg)
	return
}

func LogoutHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == http.MethodPost {
		_,err := getCurrentUser(r)
		if err != nil {
			return
		}
		err = deleteSessionFromCookie(w,r)
		if err != nil {
			return
		}
		http.Redirect(w,r,"/",http.StatusFound)
	}
}

func UserShowHandler(w http.ResponseWriter,r *http.Request)  {
	//　ここでユーザーページを取得する
	id,err := getUserIdFromURL(r)
	if err != nil {
		//ここで404 NotFound呼ぶ
	}
	user,getUserError := model.GetUserById(id)
	if getUserError != nil {
		//　ここで404 NotFound
	}
	tpl.ExecuteTemplate(w,"users_show.html",user)
}


