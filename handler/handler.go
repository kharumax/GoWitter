package handler

import (
	"GoWitter/model"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func SetUpTemplate(t *template.Template)  {
	tpl = t
}

func TopHandler(w http.ResponseWriter,r *http.Request)  {
	user,isLogin,err := model.GetCurrentUser(r)
	if err != nil {
		tpl.ExecuteTemplate(w,"index.html",user)
		return
	}
	if isLogin {
		tpl.ExecuteTemplate(w,"index.html",user)
		return
	}
	tpl.ExecuteTemplate(w,"index.html",user)
}

func BaseHandler(w http.ResponseWriter,r *http.Request)  {
	// ここでURLのパラメータ解析をして、それぞれのハンドラに振る
	// ここでは、URLの末尾に「/」が含まれている場合は、それを削除する
	if strings.HasSuffix(r.URL.Path,"/") {
		r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
	}
	url := r.URL.Path
	switch {
	case url == "" && r.Method == http.MethodGet:
		TopHandler(w,r)
		//トップページへの繊維を行う
	case strings.Contains(url,"users"):
		//ユーザー機能
	case strings.Contains(url,"posts"):
		//投稿機能＋いいね＋コメント
	case url == "/signup":
		//signupの処理に呼ぶ
		SignUpHandler(w,r)
	case url == "/login":
		LoginHandler(w,r)
		//loginの処理を呼ぶ
	case url == "/logout" && r.Method == http.MethodPost:
		//ログアウトの処理
	default:
		//404 NotFound
	}
}
