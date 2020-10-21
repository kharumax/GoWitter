package handler

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func SetUpTemplate(t *template.Template)  {
	tpl = t
}

func BaseHandler(w http.ResponseWriter,r *http.Request)  {
	// ここでURLのパラメータ解析をして、それぞれのハンドラに振る
	url := r.URL.Path
	switch {
	case url == "/" && r.Method == http.MethodGet:
		//トップページへの繊維を行う
	case url == "/users/":
		//ユーザー機能
	case url == "/posts/":
		//投稿機能＋いいね＋コメント
	case url == "/signup/" || url == "/signup":
		//signupの処理に呼ぶ
		SignUpHandler(w,r)
	case url == "/login":
		//loginの処理を呼ぶ
	case url == "/logout" && r.Method == http.MethodPost:
		//ログアウトの処理
	default:
		//404 NotFound
	}
}