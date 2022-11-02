package main

import (
	"html/template"
	"net/http"
	"unicode/utf8"
)

type LoginErrorMsg struct {
	UserNameMessage string
	PasswordMessage string
}

var loginErrorMsg LoginErrorMsg = LoginErrorMsg{}

/*
ログインボタン押下時の処理
*/
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginErrorMsg = LoginErrorMsg{}
		returnLogin(w, loginErrorMsg)
	}
	if r.Method == "POST" {
		loginErrorMsg.UserNameMessage = ""
		loginErrorMsg.PasswordMessage = ""
		r.ParseForm()

		// ユーザーネーム入力チェック
		stringLen := utf8.RuneCountInString(r.Form.Get("userName"))
		if stringLen == 0 {
			loginErrorMsg.UserNameMessage = "ユーザー名が入力されていません"
		} else if stringLen > 10 {
			loginErrorMsg.UserNameMessage = "10文字以内で入力してください"
		}

		// パスワード入力チェック

		// 入力内容によってページ遷移
		if loginErrorMsg.UserNameMessage == "" && loginErrorMsg.PasswordMessage == "" {
			userInformation = UserInformation{
				UserName: r.Form.Get("userName"),
			}
			http.Redirect(w, r, "/search/", 301)
		} else {
			returnLogin(w, loginErrorMsg)
		}
	}
}

/*
ログイン画面に跳ね返す
*/
func returnLogin(w http.ResponseWriter, loginErrorMsg LoginErrorMsg) {
	t, err := template.ParseFiles("./assets/html/login.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, loginErrorMsg); err != nil {
		panic(err.Error())
	}
}
